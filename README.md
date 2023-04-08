## 小红书开放平台 GO SDK

### 平台链接

- [小红书开放平台](https://open.xiaohongshu.com/home?from=/application/subscribeMsg)
- [小红书商家后台](https://customer.xiaohongshu.com/login?service=https://ark.xiaohongshu.com/ark/login)

### 安装命令

```sh
go get github.com/zsmhub/xhs-sdk
```

### sdk调用示例

**去 ./demo 文件夹查看完整示例！**

[点击查看完整demo](https://github.com/zsmhub/xhs-sdk/-/tree/master/demo)

### 注意点

1. 消息推送机制不健全：订单待支付状态时没有创建订单的消息推送，售后状态变更没有消息推送
2. 由于授权码有效期为10分钟，为了避免店铺重新授权的情况，我们需要用定时任务定式刷新 access_token，保证不过期