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
	return "alibaba.pifatuan.product.detail.list"
}
func (c ProductReq) Params() map[string]string {
	ret := make(map[string]string)
	ret["offerIds"] = utils.AliMarshal([]uint64{c.OfferId})
	return ret
}

type ProductResp struct {
	models.BaseBoolResponse
	Result ProductInfoInner `json:"result"`
}

type ProductInfoInner struct {
	Result []ProductInfoRespReal `json:"result"`
}
type ProductInfoRespReal struct {
	InstanceCode string      `json:"instanceCode"`
	ProductInfo  ProductInfo `json:"productInfo"`
}
type ProductInfo struct {
	Id           uint64   `json:"productID,omitempty"`
	CategoryId   uint64   `json:"categoryID,omitempty"`
	CategoryName string   `json:"categoryName,omitempty"`
	GroupId      []uint64 `json:"groupID,omitempty"`
	Subject      string   `json:"subject,omitempty"`
	Description  string   `json:"description,omitempty"`
	PictureAuth  bool     `json:"pictureAuth,omitempty"`
	Image        *struct {
		Images []string `json:"images,omitempty"`
	} `json:"image,omitempty"`
	SkuInfos           []*product.SkuInfo      `json:"skuInfos,omitempty"`
	SaleInfo           *product.SaleInfo       `json:"saleInfo,omitempty"`
	ShippingInfo       *product.ShippingInfo   `json:"shippingInfo,omitempty"`
	QualityLevel       int                     `json:"qualityLevel,omitempty"`
	SupplierLoginId    string                  `json:"supplierLoginId,omitempty"`
	MainVedio          string                  `json:"mainVedio,omitempty"`
	ProductCargoNumber string                  `json:"productCargoNumber,omitempty"`
	ReferencePrice     string                  `json:"referencePrice,omitempty"`
	Attributes         []product.Attribute     `json:"attributes,omitempty"`
	Status             string                  `json:"status,omitempty"`
	BizGroupInfos      []*product.BizGroupInfo `json:"bizGroupInfos"`
}
