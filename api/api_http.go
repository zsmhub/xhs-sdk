package api

import (
	"errors"
	"github.com/valyala/fasthttp"
	"net"
	"net/http"
	"time"
)

const HttpTTL = 1 * time.Minute

var FastClient = CreateFastHttpClient()

func CreateFastHttpClient() fasthttp.Client {
	var defaultDialer = &fasthttp.TCPDialer{Concurrency: 300}

	return fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			idx := 3 // 重试三次
			for {
				idx--
				conn, err := defaultDialer.DialTimeout(addr, 10*time.Second) // tcp连接超时时间10s
				if err != fasthttp.ErrDialTimeout || idx == 0 {
					return conn, err
				}
			}
		},
	}
}

func Post(req []byte) ([]byte, error) {
	var resp []byte
	httpReq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(httpReq)

	httpResp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(httpResp)

	httpReq.SetRequestURI(ApiUrl)
	httpReq.Header.SetContentType("application/json;charset=utf-8")
	httpReq.SetBody(req)
	httpReq.Header.SetMethod(http.MethodPost)

	if err := FastClient.DoTimeout(httpReq, httpResp, HttpTTL); err != nil {
		return resp, err
	}

	resp = httpResp.Body()
	if len(resp) == 0 { // 避免 json.Unmarshal 出现 unexpected end of JSON input 错误
		return resp, errors.New("http resp body is nil")
	}

	return resp, nil
}
