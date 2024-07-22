package model

import (
	"github.com/wzhsh90/go1688/models/trade"
	"github.com/wzhsh90/go1688/utils"
)

type CreateOrderReq struct {
	Address         *trade.Address        `json:"addressParam,omitempty"`
	CargoList       []trade.Cargo         `json:"cargoParamList,omitempty"`
	Message         string                `json:"message,omitempty"`
	OuterOrderInfo  *trade.OuterOrderInfo `json:"outerOrderInfo,omitempty"`
	TradeType       string                `json:"tradeType,omitempty"`
	Flow            string                `json:"flow,omitempty"`
	ShopPromotionId string                `json:"shopPromotionId,omitempty"`
	IsvBizTypeStr   string                `json:"isvBizTypeStr,omitempty"`
	IsvBiztypeErp   bool                  `json:"isvbiztypeErp,omitempty"`
	IsvBizTypePD    bool                  `json:"isvBizTypePD,omitempty"`
}

func (r CreateOrderReq) Name() string {
	return "alibaba.trade.fenxiaoOrder.create"
}

func (r CreateOrderReq) Params() map[string]string {
	ret := make(map[string]string, 5)
	if r.Address != nil {
		ret["addressParam"] = utils.AliMarshal(r.Address)
	}
	if len(r.CargoList) > 0 {
		ret["cargoParamList"] = utils.AliMarshal(r.CargoList)
	}
	if r.Message != "" {
		ret["message"] = r.Message
	}
	if r.OuterOrderInfo != nil {
		ret["outerOrderInfo"] = utils.AliMarshal(r.OuterOrderInfo)
	}
	if r.TradeType != "" {
		ret["tradeType"] = r.TradeType
	}
	if r.Flow != "" {
		ret["flow"] = r.Flow
	}
	if r.ShopPromotionId != "" {
		ret["shopPromotionId"] = r.ShopPromotionId
	}
	if r.IsvBiztypeErp {
		ret["isvbiztypeErp"] = "true"
	}
	ret["isvBizTypePD"] = "false"
	if r.IsvBizTypeStr != "" {
		ret["isvBizTypeStr"] = r.IsvBizTypeStr
	}
	return ret
}

type CreateOrderPreviewReq struct {
	Address         *trade.Address `json:"addressParam,omitempty"`
	CargoList       []trade.Cargo  `json:"cargoParamList,omitempty"`
	UseChannelPrice bool           `json:"useChannelPrice,omitempty"`
	Flow            string         `json:"flow,omitempty"`
}

func (r CreateOrderPreviewReq) Name() string {
	return "alibaba.createOrder.preview"
}

func (r CreateOrderPreviewReq) Params() map[string]string {
	ret := make(map[string]string, 2)
	if r.Address != nil {
		ret["addressParam"] = utils.AliMarshal(r.Address)
	}
	if len(r.CargoList) > 0 {
		ret["cargoParamList"] = utils.AliMarshal(r.CargoList)
	}
	if r.UseChannelPrice {
		ret["useChannelPrice"] = "true"
	}
	if r.Flow != "" {
		ret["flow"] = r.Flow
	}
	return ret
}
