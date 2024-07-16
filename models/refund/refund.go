package refund

import (
	"github.com/wzhsh90/go1688/models"
	"github.com/wzhsh90/go1688/utils"
	"strconv"
)

type GetRefundReasonListReq struct {
	OrderId       uint64   `json:"orderId,omitempty"`
	OrderEntryIds []uint64 `json:"orderEntryIds,omitempty"`
	GoodsStatus   string   `json:"goodsStatus,omitempty"`
}

func (r GetRefundReasonListReq) Name() string {
	return "alibaba.trade.getRefundReasonList"
}

func (r GetRefundReasonListReq) Params() map[string]string {
	return map[string]string{
		"orderId":       strconv.FormatUint(r.OrderId, 10),
		"orderEntryIds": utils.AliMarshal(r.OrderEntryIds),
		"goodsStatus":   r.GoodsStatus,
	}
}

type GetRefundReasonListResp struct {
	models.BaseResponse
	Result *GetRefundReasonListResult `json:"result,omitempty"`
}

type GetRefundReasonListResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Result  struct {
		Reasons []OrderRefundReason `json:"reasons,omitempty"`
	} `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
type OrderRefundReason struct {
	Id               uint64 `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	NeedVoucher      bool   `json:"needVoucher,omitempty"`
	NoRefundCarriage bool   `json:"noRefundCarriage,omitempty"`
	Tip              string `json:"tip,omitempty"`
}

type CreateRefundReq struct {
	OrderID             uint64            `json:"orderId,omitempty"`
	OrderEntryIds       []uint64          `json:"orderEntryIds,omitempty"`
	DisputeRequest      string            `json:"disputeRequest,omitempty"`
	ApplyPayment        int64             `json:"applyPayment,omitempty"`
	ApplyCarriage       int64             `json:"applyCarriage,omitempty"`
	ApplyReasonID       uint64            `json:"applyReasonId,omitempty"`
	Description         string            `json:"description,omitempty"`
	GoodsStatus         string            `json:"goodsStatus,omitempty"`
	Vouchers            []string          `json:"vouchers,omitempty"`
	OrderEntryCountList []OrderEntryCount `json:"orderEntryCountList,omitempty"`
}

func (r CreateRefundReq) Name() string {
	return "alibaba.trade.createRefund"
}

func (r CreateRefundReq) Params() map[string]string {
	ret := make(map[string]string, 10)
	ret["orderId"] = strconv.FormatUint(r.OrderID, 10)
	if len(r.OrderEntryIds) > 0 {
		ret["orderEntryIds"] = utils.AliMarshal(r.OrderEntryIds)
	}
	ret["disputeRequest"] = r.DisputeRequest
	ret["applyPayment"] = strconv.FormatInt(r.ApplyPayment, 10)
	ret["applyCarriage"] = strconv.FormatInt(r.ApplyCarriage, 10)
	ret["applyReasonId"] = strconv.FormatUint(r.ApplyReasonID, 10)
	ret["description"] = r.Description
	ret["goodsStatus"] = r.GoodsStatus
	if len(r.Vouchers) > 0 {
		ret["vouchers"] = utils.AliMarshal(r.Vouchers)
	}
	if len(r.OrderEntryCountList) > 0 {
		ret["orderEntryCountList"] = utils.AliMarshal(r.OrderEntryCountList)
	}
	return ret
}

type OrderEntryCount struct {
	Id    uint64 `json:"id,omitempty"`
	Count int    `json:"count,omitempty"`
}

type CreateRefundResp struct {
	models.BaseResponse
	Result *CreateRefundResult `json:"result,omitempty"`
}

type CreateRefundResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Result  struct {
		RefundId string `json:"refundId,omitempty"`
	} `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
