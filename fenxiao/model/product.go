package model

import (
	"github.com/wzhsh90/go1688/models"
	"github.com/wzhsh90/go1688/models/product"
	"github.com/wzhsh90/go1688/utils"
)

type ProductReq struct {
	OfferId uint64
}

func (c *ProductReq) Name() string {
	return "alibaba.fenxiao.productInfo.get"
}
func (c ProductReq) Params() map[string]string {
	ret := make(map[string]string)
	ret["offerId"] = utils.FormatUint(c.OfferId)
	return ret
}

type ProductResp struct {
	models.BaseBoolResponse
	ProductInfo *ProductInfo `json:"productInfo"`
}
type ProductInfo struct {
	ProductID    uint64   `json:"productID,omitempty"`
	ProductType  string   `json:"productType,omitempty"`
	CategoryID   uint64   `json:"categoryID,omitempty"`
	CategoryName string   `json:"categoryName,omitempty"`
	GroupID      []uint64 `json:"groupID,omitempty"`
	Subject      string   `json:"subject,omitempty"`
	Description  string   `json:"description,omitempty"`
	BizType      int64    `json:"bizType,omitempty"`
	PictureAuth  bool     `json:"pictureAuth,omitempty"`
	ProductImage *struct {
		Images []string `json:"images,omitempty"`
	} `json:"productImage,omitempty"`
	ProductSkuInfos     []*product.SkuInfo      `json:"productSkuInfos,omitempty"`
	ProductSaleInfo     *product.SaleInfo       `json:"productSaleInfo,omitempty"`
	ProductShippingInfo *product.ShippingInfo   `json:"productShippingInfo,omitempty"`
	QualityLevel        int                     `json:"qualityLevel,omitempty"`
	SupplierLoginID     string                  `json:"supplierLoginId,omitempty"`
	MainVideo           string                  `json:"mainVedio,omitempty"`
	ProductCargoNumber  string                  `json:"productCargoNumber,omitempty"`
	ReferencePrice      string                  `json:"referencePrice,omitempty"`
	Attributes          []product.Attribute     `json:"productAttribute,omitempty"`
	Status              string                  `json:"status,omitempty"`
	BizGroupInfos       []*product.BizGroupInfo `json:"productBizGroupInfos"`
	ProductExtendInfos  []*ProductExtend        `json:"productExtendInfos"`
}
type ProductExtend struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type OutShopReq struct {
	OutShopCode string `json:"outShopCode,omitempty"`
	Channel     string `json:"channel,omitempty"`
}

func (r OutShopReq) Name() string {
	return "alibaba.fenxiao.buyer.outshop.add"
}
func (r OutShopReq) Params() map[string]string {
	ret := make(map[string]string)
	ret["outShopCode"] = r.OutShopCode
	ret["channel"] = r.Channel
	return ret
}

type OutShopResp struct {
	models.BaseBoolResponse
	Data string `json:"data"`
}

type OutProductReq struct {
	OfferId     uint64 `json:"offerId,omitempty"`
	OutItemCode string `json:"outItemCode,omitempty"`
	OutShopCode string `json:"outShopCode,omitempty"`
	Channel     string `json:"channel,omitempty"`
}

func (r OutProductReq) Name() string {
	return "alibaba.fenxiao.buyer.outproduct.relation.add"
}
func (r OutProductReq) Params() map[string]string {
	ret := make(map[string]string)
	ret["outShopCode"] = r.OutShopCode
	ret["outItemCode"] = r.OutItemCode
	ret["offerId"] = utils.FormatUint(r.OfferId)
	ret["channel"] = r.Channel
	return ret
}

type OutProductResp struct {
	models.BaseBoolResponse
	Data string `json:"data"`
}
type OutProductRelReq struct {
	OfferId     uint64 `json:"offerId,omitempty"`
	OutItemCode string `json:"outItemCode,omitempty"`
	OutShopCode string `json:"outShopCode,omitempty"`
	Channel     string `json:"channel,omitempty"`
}

func (r OutProductRelReq) Name() string {
	return "alibaba.fenxiao.buyer.outproduct.relation.get"
}
func (r OutProductRelReq) Params() map[string]string {
	ret := make(map[string]string)
	ret["outShopCode"] = r.OutShopCode
	ret["outItemCode"] = r.OutItemCode
	ret["offerId"] = utils.FormatUint(r.OfferId)
	ret["channel"] = r.Channel
	return ret
}

type OutProductRelResp struct {
	models.BaseBoolResponse
	Data string `json:"data"`
}

type FollowReq struct {
	ProductId uint64 `json:"productId,omitempty"`
}

func (r FollowReq) Name() string {
	return "alibaba.product.follow"
}

func (r FollowReq) Params() map[string]string {
	return map[string]string{
		"productId": utils.FormatUint(r.ProductId),
	}
}

type FollowResp struct {
	models.BaseBoolResponse
	Code    int64  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type UnFollowReq struct {
	ProductId uint64 `json:"productId"`
}

func (r UnFollowReq) Name() string {
	return "alibaba.product.unfollow.crossborder"
}

func (r UnFollowReq) Params() map[string]string {
	return map[string]string{
		"productId": utils.FormatUint(r.ProductId),
	}
}

type UnFollowResp struct {
	models.BaseBoolResponse
	Code    int64  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
