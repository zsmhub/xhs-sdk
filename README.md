## 小红书开放平台 GO SDK

Go语言实现小红书开放平台sdk，使用简单，扩展灵活

### 平台链接

- [小红书开放平台](https://open.xiaohongshu.com/home?from=/application/subscribeMsg)
- [小红书商家后台](https://customer.xiaohongshu.com/login?service=https://ark.xiaohongshu.com/ark/login)

### 安装命令

```sh
go get github.com/zsmhub/xhs-sdk
```

### sdk调用示例

**强烈建议去 ./demo 文件夹查看完整示例！**

[点击查看完整demo](https://github.com/zsmhub/xhs-sdk/tree/main/demo)

## 推送消息调用示例

```go
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
```

### api sdk 调用示例

```go
func GetOAuthUrl(redirectUri, state string) {
	authUrl := xhs_sdk.NewXhsOAuthClient(AppKey, AppSecret).GetOAuthUrl(redirectUri, state)
	fmt.Printf("authUrl=%s\n", authUrl)
}

func GetAccessToken() {
	accessToken, err := xhs_sdk.NewXhsOAuthClient(AppKey, AppSecret).GetAccessToken(Code)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", accessToken)
}

func GetExpressCompanyList() {
	var resp api.RespGetExpressCompanyList
	err := xhs_sdk.NewXhsClient(AppKey, AppSecret).SetAccessToken(AccessToken).Post(api.ReqGetExpressCompanyList{}, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	ret, _ := json.Marshal(resp)
	fmt.Printf("%s\n", string(ret))
}
```

### 目录结构

```sh
.
├── api          调用API
├── message      推送消息
├── demo         sdk调用示例
├── constant.go  全局枚举值定义
└── sdk.go       入口文件
```

### 注意点

1. 如果你发现了sdk中，没有某个推送消息或某个api，可自行加上，然后提交下pr
2. 消息推送机制不健全：订单待支付状态时没有创建订单的消息推送，售后状态变更没有消息推送，需要自行用定时任务轮询小红书接口获取
3. 由于授权码有效期为10分钟，为了避免店铺重新授权的情况，我们需要用定时任务定时刷新 access_token，保证不过期

## 推荐开源项目

- [企业微信 GO SDK](https://github.com/zsmhub/workweixin)
- [抖店开放平台GO SDK](https://github.com/zsmhub/doudian-sdk)
- [微信视频号GO SDK](https://github.com/zsmhub/wx-channels-sdk)
- [小红书开放平台 GO SDK](https://github.com/zsmhub/xhs-sdk)