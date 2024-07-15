package client

import (
	"github.com/wzhsh90/go1688/consts"
	"github.com/wzhsh90/go1688/models"
	"github.com/wzhsh90/go1688/models/logistics"
)

func (c *Client) GetLogisticsInfos(orderId uint64) ([]*logistics.Order, error) {
	req := logistics.GetLogisticsInfosBuyerViewReq{
		OrderId: orderId,
		Fields:  "",
		Website: "1688",
	}
	resp := &logistics.GetLogisticsInfosBuyerViewResp{}
	err := c.Do(models.NewRequest(consts.LOGISTIC_SPACE, req), resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

func (c *Client) GetLogisticsTraceInfosBuyerViewReq(orderId uint64) ([]logistics.Trace, error) {
	req := &logistics.GetLogisticsTraceInfosBuyerViewReq{
		OrderId:     orderId,
		LogisticsId: "",
		Website:     "1688",
	}
	resp := &logistics.GetLogisticsTraceInfosBuyerViewResp{}
	err := c.Do(models.NewRequest(consts.LOGISTIC_SPACE, req), resp)
	if err != nil {
		return nil, err
	}
	return resp.Traces, nil

}
