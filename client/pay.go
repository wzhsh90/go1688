package client

import (
	"github.com/wzhsh90/go1688/consts"
	"github.com/wzhsh90/go1688/models"
	"github.com/wzhsh90/go1688/models/pay"
)

func (c *Client) PreparePay(orderId uint64) (bool, error) {
	req := &pay.PreparePayReq{
		OrderId: orderId,
	}
	resp := &pay.PreparePayResp{}
	err := c.Do(models.NewRequest(consts.TRADE_SPACE, req), resp)
	if err != nil {
		return false, err
	}
	return resp.Result.PaySuccess, nil
}
func (c *Client) ProtocolPay(orderId uint64) (bool, error) {
	req := &pay.ProtocolPayReq{
		OrderId: orderId,
	}
	resp := &pay.ProtocolPayResp{}
	err := c.Do(models.NewRequest(consts.TRADE_SPACE, req), resp)
	if err != nil {
		return false, err
	}
	return resp.Code == "0", nil
}
func (c *Client) CancelPay(orderId uint64) (bool, error) {
	req := &pay.CancelPayReq{
		Website: "1688",
		TradeId: orderId,
		Reason:  "other",
		Remark:  "用户取消",
	}
	resp := &pay.CancelPayResp{}
	err := c.Do(models.NewRequest(consts.TRADE_SPACE, req), resp)
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}
