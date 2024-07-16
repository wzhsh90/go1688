package client

import (
	"github.com/wzhsh90/go1688/consts"
	"github.com/wzhsh90/go1688/fenxiao/model"
	"github.com/wzhsh90/go1688/models"
)

func (c *Client) FollowProduct(productId uint64) (bool, error) {
	req := &model.FollowReq{
		ProductId: productId,
	}
	resp := &model.FollowResp{}
	err := c.Do(models.NewRequest(consts.PRODUCT_SPACE, req), resp)
	if err != nil {
		return false, err
	}
	return resp.Code == 0, nil
}
func (c *Client) UnFollowProduct(productId uint64) (bool, error) {
	req := &model.UnFollowReq{
		ProductId: productId,
	}
	resp := &model.UnFollowResp{}
	err := c.Do(models.NewRequest(consts.PRODUCT_SPACE, req), resp)
	if err != nil {
		return false, err
	}
	return resp.Code == 0, nil
}
