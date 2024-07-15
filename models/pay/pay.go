package pay

import (
	"github.com/wzhsh90/go1688/models"
	"github.com/wzhsh90/go1688/utils"
)

type PreparePayReq struct {
	OrderId uint64
}

func (r PreparePayReq) Name() string {
	return "alibaba.trade.pay.protocolPay.preparePay"
}

func (r PreparePayReq) Params() map[string]string {
	return map[string]string{
		"tradeWithholdPreparePayParam": utils.AliMarshal(map[string]uint64{
			"orderId": r.OrderId,
		}),
	}
}

type PreparePayResp struct {
	models.BaseBoolResponse
	Result PreparePayRespResult `json:"result"`
}
type PreparePayRespResult struct {
	PayChannel string `json:"payChannel,omitempty"`
	PaySuccess bool   `json:"paySuccess,omitempty"`
}

type CancelPayReq struct {
	Website string `json:"webSite,omitempty"`
	TradeId uint64 `json:"tradeID,omitempty"`
	Reason  string `json:"cancelReason,omitempty"`
	Remark  string `json:"remark,omitempty"`
}

func (r CancelPayReq) Name() string {
	return "alibaba.trade.cancel"
}

func (r CancelPayReq) Params() map[string]string {
	ret := make(map[string]string, 4)
	ret["webSite"] = r.Website
	ret["tradeID"] = utils.FormatUint(r.TradeId)
	ret["cancelReason"] = r.Reason
	if r.Remark != "" {
		ret["remark"] = r.Remark
	}
	return ret
}

type CancelPayResp struct {
	models.BaseResponse
	Success bool `json:"success,omitempty"`
}
