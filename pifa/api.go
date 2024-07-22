package pifa

import (
	"errors"
	"github.com/wzhsh90/go1688/client"
	"github.com/wzhsh90/go1688/consts"
	"github.com/wzhsh90/go1688/models"
	"github.com/wzhsh90/go1688/models/trade"
	"github.com/wzhsh90/go1688/pifa/model"
)

type Api struct {
	client *client.Client
}

func NewPiFaApi(client *client.Client) *Api {
	return &Api{
		client: client,
	}
}
func (c *Api) GetProduct(offerId uint64) (*model.ProductInfo, error) {
	req := &model.ProductReq{
		OfferId: offerId,
	}
	resp := &model.ProductResp{}
	err := c.client.Do(models.NewRequest(consts.FENXIAO_NAMESPACE, req), resp)
	if err != nil {
		return nil, err
	}
	list := resp.Result.Result
	if len(list) == 0 {
		return nil, errors.New("empty")
	}
	return &list[0].ProductInfo, err
}

func (c *Api) CreateOrder(address *trade.Address, cargo []trade.Cargo, offers []trade.OuterOrderOffer, tradeNo, remark string) (*trade.CreateOrderResult, error) {
	outerOrderInfo := trade.OuterOrderInfo{
		MediaOrderId: tradeNo,
		Phone:        "",
		Offers:       offers,
	}
	req := &model.FastCreateOrderReq{
		Address:        address,
		CargoList:      cargo,
		Message:        remark,
		OuterOrderInfo: &outerOrderInfo,
		IsvBizTypeStr:  "fenxiaoMedia",
		Flow:           "ttpft",
	}
	resp := &model.FastCreateOrderResp{}
	err := c.client.Do(models.NewRequest(consts.TRADE_SPACE, req), resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil

}
func (c *Api) OrderPreview(address *trade.Address, cargo []trade.Cargo) (*model.CreateOrderPreviewResp, error) {
	req := &model.CreateOrderPreviewReq{
		Address:   address,
		CargoList: cargo,
		Flow:      "ttpft",
	}
	resp := &model.CreateOrderPreviewResp{}
	err := c.client.Do(models.NewRequest(consts.TRADE_SPACE, req), resp)
	if err != nil {
		return nil, err
	}
	return resp, err

}
