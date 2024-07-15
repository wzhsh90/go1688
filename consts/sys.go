package consts

const (
	GATEWAY  = "http://gw.open.1688.com/openapi"
	PROTOCOL = "param2"
	VERSION  = "1"
	IMG_BASE = "https://cbu01.alicdn.com/"

	PRODUCT_SPACE     = "com.alibaba.product"
	TRADE_SPACE       = "com.alibaba.trade"
	LOGISTIC_SPACE    = "com.alibaba.logistics"
	FENXIAO_NAMESPACE = "com.alibaba.fenxiao"
)

var PayStatusMap = map[string]string{
	"1":  "未付款",
	"2":  "已付款",
	"4":  "全额退款",
	"6":  "卖家有收到钱，回款完成",
	"7":  "未创建外部支付单",
	"8":  "付款前取消",
	"9":  "正在支付中",
	"12": "账期支付,待到账",
}
