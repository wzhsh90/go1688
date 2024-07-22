package fenxiao

import (
	"github.com/wzhsh90/go1688/client"
	"github.com/wzhsh90/go1688/consts"
	"github.com/wzhsh90/go1688/fenxiao/model"
	"github.com/wzhsh90/go1688/models"
	"github.com/wzhsh90/go1688/models/trade"
)

type Api struct {
	client *client.Client
}

func NewFenXiaoApi(client *client.Client) *Api {
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
	return resp.ProductInfo, err
}

func (c *Api) CreateOrder(address *trade.Address, cargo []trade.Cargo, offers []trade.OuterOrderOffer, tradeNo, remark string) (*trade.CreateOrderResult, error) {
	outerOrderInfo := trade.OuterOrderInfo{
		MediaOrderId: tradeNo,
		Phone:        "",
		Offers:       offers,
	}
	req := &model.CreateOrderReq{
		Address:        address,
		CargoList:      cargo,
		Message:        remark,
		OuterOrderInfo: &outerOrderInfo,
		IsvBizTypeStr:  "fenxiaoMedia",
		Flow:           "fenxiao",
	}
	resp := &trade.CreateOrderResp{}
	err := c.client.Do(models.NewRequest(consts.TRADE_SPACE, req), resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil

}
func (c *Api) OrderPreview(address *trade.Address, cargo []trade.Cargo) (*trade.CreateOrderPreviewResp, error) {
	req := &model.CreateOrderPreviewReq{
		Address:   address,
		CargoList: cargo,
		Flow:      "saleproxy",
	}
	resp := &trade.CreateOrderPreviewResp{}
	err := c.client.Do(models.NewRequest(consts.TRADE_SPACE, req), resp)
	if err != nil {
		return nil, err
	}
	return resp, err

}
func (c *Api) OutShop(shopId string) (bool, error) {
	req := model.OutShopReq{
		OutShopCode: shopId,
		Channel:     "other",
	}
	resp := &model.OutShopResp{}
	err := c.client.Do(models.NewRequest(consts.FENXIAO_NAMESPACE, req), resp)
	if err != nil {
		return false, err
	}
	return resp.Data == "true", nil

}
func (c *Api) OutProduct(shopId, goodsId string, aliId uint64) (bool, error) {
	req := model.OutProductReq{
		OfferId:     aliId,
		OutItemCode: goodsId,
		OutShopCode: shopId,
		Channel:     "other",
	}
	resp := &model.OutProductResp{}
	err := c.client.Do(models.NewRequest(consts.FENXIAO_NAMESPACE, req), resp)
	if err != nil {
		return false, err
	}
	return resp.Data == "true", nil
}

func (c *Api) OutProductRel(shopId, goodsId string, aliId uint64) (bool, error) {
	req := model.OutProductRelReq{
		OfferId:     aliId,
		OutItemCode: goodsId,
		OutShopCode: shopId,
		Channel:     "other",
	}
	resp := &model.OutProductRelResp{}
	err := c.client.Do(models.NewRequest(consts.FENXIAO_NAMESPACE, req), resp)
	if err != nil {
		return false, err
	}
	return resp.Data == "true", nil

}
