package demo

import "testing"

func TestGetOAuthUrl(t *testing.T) {
	GetOAuthUrl("https://www.xxx.com/message/auth", "123")
}

func TestGetAccessToken(t *testing.T) {
	GetAccessToken()
}

func TestRefreshAccessToken(t *testing.T) {
	RefreshAccessToken()
}

func TestGetExpressCompanyList(t *testing.T) {
	GetExpressCompanyList()
}

func TestGetZones(t *testing.T) {
	GetZones()
}

func TestGetNestZone(t *testing.T) {
	GetNestZone()
}

func TestGetOrderList(t *testing.T) {
	GetOrderList()
}

func TestGetOrderDetail(t *testing.T) {
	GetOrderDetail()
}

func TestGetOrderReceiverInfo(t *testing.T) {
	GetOrderReceiverInfo()
}

func TestBatchDecrypt(t *testing.T) {
	BatchDecrypt()
}

func TestGetAfterSaleDetail(t *testing.T) {
	GetAfterSaleDetail()
}
