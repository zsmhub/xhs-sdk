package api

const (
	ApiUrl  = "https://ark.xiaohongshu.com/ark/open_api/v3/common_controller" // 接口调用链接
	Version = "2.0"                                                           // 版本号，采用oauth2授权后均填入2.0
)

type BaseRequest struct {
	Method      string `json:"method"`
	AppId       string `json:"appId"`
	Sign        string `json:"sign"`
	Timestamp   int64  `json:"timestamp"`
	Version     string `json:"version"`
	AccessToken string `json:"accessToken,omitempty"`
}

type Request interface {
	Method() string
	Params(BaseRequest) []byte
}

type BaseResponse struct {
	ErrorCode int64  `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	Success   bool   `json:"success"`
}

type Response interface {
	ErrorCode() int64
	ErrorMsg() string
	Success() bool
}

const (
	SuccessCode = 0
)
