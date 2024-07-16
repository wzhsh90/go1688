package client

import (
	"github.com/wzhsh90/go1688/consts"
	"github.com/wzhsh90/go1688/models"
	"github.com/wzhsh90/go1688/models/refund"
)

func (c *Client) GetRefundReasonList(orderId uint64, goodsStatus string) (*refund.GetRefundReasonListResult, error) {

	req := &refund.GetRefundReasonListReq{
		OrderId:       orderId,
		OrderEntryIds: []uint64{orderId},
		GoodsStatus:   goodsStatus,
	}
	resp := &refund.GetRefundReasonListResp{}
	err := c.Do(models.NewRequest(consts.TRADE_SPACE, req), resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}
func (c *Client) CreateRefund(req refund.CreateRefundReq) (*refund.CreateRefundResult, error) {

	resp := &refund.CreateRefundResp{}
	err := c.Do(models.NewRequest(consts.TRADE_SPACE, req), resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}
