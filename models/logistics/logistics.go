package logistics

import (
	"github.com/wzhsh90/go1688/models"
	"github.com/wzhsh90/go1688/utils"
)

type GetLogisticsInfosBuyerViewReq struct {
	OrderId uint64 `json:"orderId"`
	Fields  string `json:"fields"`
	Website string `json:"webSite"`
}

func (r GetLogisticsInfosBuyerViewReq) Name() string {
	return "alibaba.trade.getLogisticsInfos.buyerView"
}

func (r GetLogisticsInfosBuyerViewReq) Params() map[string]string {
	ret := make(map[string]string)
	if r.OrderId > 0 {
		ret["orderId"] = utils.FormatUint(r.OrderId)
	}
	ret["fields"] = r.Fields
	if r.Website != "" {
		ret["webSite"] = r.Website
	}
	return ret
}

type GetLogisticsInfosBuyerViewResp struct {
	models.BaseBoolResponse
	Result []*Order `json:"result"`
}

type Order struct {
	Sender               *Sender    `json:"sender,omitempty"`
	ServiceFeature       string     `json:"serviceFeature,omitempty"`
	OrderEntryIds        string     `json:"orderEntryIds,omitempty"`
	LogisticsBillNo      string     `json:"logisticsBillNo,omitempty"`
	LogisticsId          string     `json:"logisticsId,omitempty"`
	Receiver             *Receiver  `json:"receiver,omitempty"`
	LogisticsCompanyId   string     `json:"logisticsCompanyId,omitempty"`
	LogisticsCompanyName string     `json:"logisticsCompanyName,omitempty"`
	Status               string     `json:"status,omitempty"`
	SendGoods            []SendGood `json:"sendGoods,omitempty"`
	GmtSystemSend        string     `json:"gmtSystemSend,omitempty"`
	Remarks              string     `json:"remarks,omitempty"`
}

type Sender struct {
	Name         string `json:"senderName,omitempty"`
	Phone        string `json:"senderPhone,omitempty"`
	Mobile       string `json:"senderMobile,omitempty"`
	Encrypt      string `json:"encrypt,omitempty"`
	ProvinceCode string `json:"senderProvinceCode,omitempty"`
	CityCode     string `json:"senderCityCode,omitempty"`
	CountyCode   string `json:"senderCountryCode,omitempty"`
	Address      string `json:"senderAddress,omitempty"`
	Province     string `json:"senderProvince,omitempty"`
	City         string `json:"senderCity,omitempty"`
	County       string `json:"senderCounty,omitempty"`
}

type Receiver struct {
	Name         string `json:"receiverName,omitempty"`
	Phone        string `json:"receiverPhone,omitempty"`
	Mobile       string `json:"receiverMobile,omitempty"`
	Encrypt      string `json:"encrypt,omitempty"`
	ProvinceCode string `json:"receiverProvinceCode,omitempty"`
	CityCode     string `json:"receiverCityCode,omitempty"`
	CountyCode   string `json:"receiverCountryCode,omitempty"`
	Address      string `json:"receiverAddress,omitempty"`
	Province     string `json:"receiverProvince,omitempty"`
	City         string `json:"receiverCity,omitempty"`
	County       string `json:"receiverCounty,omitempty"`
}
type SendGood struct {
	Name     string `json:"goodName,omitempty"`
	Quantity string `json:"quantity,omitempty"`
	Unit     string `json:"unit,omitempty"`
}

type GetLogisticsTraceInfosBuyerViewReq struct {
	LogisticsId string `json:"logisticsId,omitempty"`
	OrderId     uint64 `json:"orderId,omitempty"`
	Website     string `json:"webSite,omitempty"`
}

func (r GetLogisticsTraceInfosBuyerViewReq) Name() string {
	return "alibaba.trade.getLogisticsTraceInfo.buyerView"
}

func (r GetLogisticsTraceInfosBuyerViewReq) Params() map[string]string {
	ret := make(map[string]string, 3)
	if r.LogisticsId != "" {
		ret["logisticsId"] = r.LogisticsId
	}
	ret["orderId"] = utils.FormatUint(r.OrderId)
	ret["webSite"] = r.Website
	return ret
}

type GetLogisticsTraceInfosBuyerViewResp struct {
	models.BaseBoolResponse
	Traces []Trace `json:"logisticsTrace,omitempty"`
}

type Trace struct {
	LogisticsId     string `json:"logisticsId,omitempty"`
	OrderId         uint64 `json:"orderId,omitempty"`
	LogisticsBillNo string `json:"logisticsBillNo,omitempty"`
	Steps           []Step `json:"logisticsSteps,omitempty"`
}

type Step struct {
	AcceptTime string `json:"acceptTime,omitempty"`
	Remark     string `json:"remark,omitempty"`
}
