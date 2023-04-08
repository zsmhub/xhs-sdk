package xhs_sdk

// 订单状态
const (
	OrderStatusWaitPay    int64 = 1  // 已下单待付款
	OrderStatusPaid       int64 = 2  // 已支付处理中
	OrderStatusClean      int64 = 3  // 清关中
	OrderStatusStocking   int64 = 4  // 待发货
	OrderStatusPartSend   int64 = 5  // 部分发货
	OrderStatusSend       int64 = 6  // 待收货
	OrderStatusFinish     int64 = 7  // 已完成
	OrderStatusClose      int64 = 8  // 已关闭
	OrderStatusCancel     int64 = 9  // 已取消
	OrderStatusExchanging int64 = 10 // 换货申请中
)

// 订单售后状态
const (
	OrderAfterSalesStatusNo       int64 = 1 //  无售后
	OrderAfterSalesStatusIng      int64 = 2 //  售后处理中
	OrderAfterSalesStatusFinish   int64 = 3 //  售后完成(含取消)
	OrderAfterSalesStatusRefuse   int64 = 4 //  售后拒绝
	OrderAfterSalesStatusClose    int64 = 5 //  售后关闭
	OrderAfterSalesStatusPlatform int64 = 6 //  平台介入中
)

// 订单类型
const (
	OrderTypeNormal   int64 = 1 // 普通
	OrderTypePresell  int64 = 2 // 定金预售
	OrderTypeFullPay  int64 = 3 // 全款预售
	OrderTypeDelay    int64 = 4 // 延迟发货
	OrderTypeExchange int64 = 5 // 换货补发
)

// 订单申请取消状态
const (
	OrderCancelStatusNo  int64 = 0 // 未申请取消
	OrderCancelStatusIng int64 = 1 // 取消处理中
)

// 物流模式
const (
	LogisticsRedExpress  = "red_express"  // 三方备货直邮(备货海外仓)
	LogisticsRedStandard = "red_standard" // 三方备货保税仓
	LogisticsRedAuto     = "red_auto"     // 三方自主发货
	LogisticsRedBox      = "red_box"      // 三方小包
	LogisticsRedBonded   = "red_bonded"   // 三方保税
)

// 订单列表筛选时间类型
const (
	OrderTimeType1 int64 = 1 // 创建时间 限制 end-start<=24h
	OrderTimeType2 int64 = 2 // 更新时间 限制 end-start<=30min 倒序拉取 最后一页到第一页
)

// 售后退货类型
const (
	AfsReturnTypeReturn       int64 = 1 // 退货退款
	AfsReturnTypeExchange     int64 = 2 // 换货
	AfsReturnTypeRefundOld    int64 = 3 // 仅退款(old)，理论上不会有3出现
	AfsReturnTypeRefundNew    int64 = 4 // 仅退款(new)
	AfsReturnTypeNoSendRefund int64 = 5 // 未发货仅退款
)

// 售后单售后状态
const (
	AfsStatusWaitAudit      int64 = 1    // 待审核
	AfsStatusWaitBack       int64 = 2    // 待用户寄回
	AfsStatusWaitDelivery   int64 = 3    // 待收货
	AfsStatusFinish         int64 = 4    // 完成
	AfsStatusCancel         int64 = 5    // 取消
	AfsStatusClose          int64 = 6    // 关闭
	AfsStatusRefuse         int64 = 9    // 拒绝
	AfsStatusMerchantRefuse int64 = 9001 // 商家收货拒绝
	AfsStatusChangeRefund   int64 = 11   // 换货转退款
	AfsStatusWaitSend       int64 = 12   // 包裹商家已确认收货，等待商家发货
	AfsStatusSend           int64 = 13   // 包裹商家已发货，等待用户确认收货。不传默认全部
	AfsStatusPlatform       int64 = 14   // 平台介入中
)

// 售后单售后子状态
const (
	AfsSubStatus301 int64 = 301 // 待审核
	AfsSubStatus302 int64 = 302 // 快递已签收
	AfsSubStatus304 int64 = 304 // 收货异常
)

// 收货异常类型
const (
	ReceiveAbnormalType1  = 1  // 商家已开工单
	ReceiveAbnormalType2  = 2  // 仓库质检异常
	ReceiveAbnormalType3  = 3  // 收货地异常
	ReceiveAbnormalType4  = 4  // 寄回包裹物流超时
	ReceiveAbnormalType5  = 5  // 仓库反向创建退货
	ReceiveAbnormalType6  = 6  // 快递轨迹异常
	ReceiveAbnormalType7  = 7  // 拒收退仓超时
	ReceiveAbnormalType8  = 8  // 退款金额超限
	ReceiveAbnormalType9  = 9  // 收货地不一致
	ReceiveAbnormalType10 = 10 // 仓库质检假货
	ReceiveAbnormalType11 = 11 // 已退款，未收货
	ReceiveAbnormalType21 = 21 // 未收到货
	ReceiveAbnormalType22 = 22 // 退货数量不符
	ReceiveAbnormalType23 = 23 // 退货商品不符
	ReceiveAbnormalType24 = 24 // 退货质检异常
	ReceiveAbnormalType25 = 25 // 其他
)

// 解密操作类型
const (
	DecryptActionType1 = "1" // 单个查看订单明文
	DecryptActionType2 = "2" // 批量解密打单
	DecryptActionType3 = "3" // 批量解密后面向三方的数据下发
	DecryptActionType4 = "4" // 其他场景
)
