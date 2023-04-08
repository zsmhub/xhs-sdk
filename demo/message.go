package demo

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/zsmhub/xhs-sdk/message"
	"net/http"
)

// 接收小红书开放平台消息推送示例
func CallbackMain() {
	e := echo.New()

	e.GET("/message/auth", func(c echo.Context) error {
		// 打印消息推送请求体，便于调试
		// requestBody, _ := ioutil.ReadAll(c.Request().Body)
		// fmt.Printf("xhs auth : uri=%s, req=%s\n", c.Request().RequestURI, string(requestBody))

		code := c.Param("code")
		state := c.Param("state")
		fmt.Printf("code=%s, state=%s\n", code, state)

		return c.JSON(http.StatusOK, message.MessageResp{
			Success:   true,
			ErrorCode: 0,
			ErrorMsg:  "",
		})
	}).Name = "应用授权回调地址"

	e.POST("/message/push", func(c echo.Context) error {
		// 打印消息推送请求体，便于调试
		// requestBody, _ := ioutil.ReadAll(c.Request().Body)
		// fmt.Printf("xhs push : uri=%s, req=%s\n", c.Request().RequestURI, string(requestBody))

		var req []message.Message
		_ = c.Bind(&req)
		for _, v := range req {
			vJson, _ := json.Marshal(v)
			fmt.Printf("msg=%s\n", vJson)

			// 消息data解析示例
			if v.MsgTag == (message.MsgFulfillmentStatusChange{}.MsgTag()) {
				// 方案一
				//var resp message.MsgFulfillmentStatusChange
				//if err := v.DecodeData(&resp); err != nil {
				//    fmt.Println(err)
				//    continue
				//} else {
				//    fmt.Printf("resp=%+v\n", resp)
				//}

				// 方案二
				resp, err := message.MsgFulfillmentStatusChange{}.DecodeData(v.Data)
				if err != nil {
					fmt.Println(err)
					continue
				} else {
					fmt.Printf("resp=%+v\n", resp)
				}
			}
		}

		return c.JSON(http.StatusOK, message.MessageResp{
			Success:   true,
			ErrorCode: 0,
			ErrorMsg:  "",
		})
	}).Name = "应用推送回调地址"

	e.Logger.Fatal(e.Start(":1323"))
	select {}
}
