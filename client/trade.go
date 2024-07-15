package client

import (
	"github.com/wzhsh90/go1688/consts"
	"github.com/wzhsh90/go1688/models"
	"github.com/wzhsh90/go1688/models/trade"
)

func (c *Client) GetTradeInfo(orderId uint64) (*trade.Info, error) {
	req := &trade.GetBuyerViewReq{
		Website:       "1688",
		OrderId:       orderId,
		IncludeFields: "",
		AttributeKeys: nil,
	}
	resp := &trade.GetBuyerViewResp{}
	err := c.Do(models.NewRequest(consts.TRADE_SPACE, req), resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}
