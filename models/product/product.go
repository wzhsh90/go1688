package product

import "github.com/shopspring/decimal"

type SaleInfo struct {
	// SupportOnlineTrade 是否支持网上交易。true：支持 false：不支持
	SupportOnlineTrade bool            `json:"supportOnlineTrade,omitempty"`
	MixWholeSale       bool            `json:"mixWholeSale,omitempty"`
	PriceAuth          bool            `json:"priceAuth,omitempty"`
	PriceRanges        []PriceRange    `json:"priceRanges,omitempty"`
	AmountOnSale       float64         `json:"amountOnSale,omitempty"`
	SaleType           string          `json:"saleType,omitempty"`
	Unit               string          `json:"unit,omitempty"`
	MinOrderQuantity   int64           `json:"minOrderQuantity,omitempty"`
	BatchNumber        int64           `json:"batchNumber,omitempty"`
	RetailPrice        decimal.Decimal `json:"retailprice,omitempty"`
	SellUnit           string          `json:"sellunit,omitempty"`
	QuoteType          int64           `json:"quoteType,omitempty"`
	ConsignPrice       decimal.Decimal `json:"consignPrice,omitempty"`
	CpsSuggestPrice    decimal.Decimal `json:"cpsSuggestPrice,omitempty"`
	ChannelPrice       decimal.Decimal `json:"channelPrice,omitempty"`
}

type PriceRange struct {
	StartQuantity int64           `json:"startQuantity,omitempty"`
	Price         decimal.Decimal `json:"price,omitempty"`
}

// ShippingInfo 商品物流信息
type ShippingInfo struct {
	FreightTemplateId    uint64  `json:"freightTemplateID,omitempty"`
	UnitWeight           float64 `json:"unitWeight,omitempty"`
	PackageSize          string  `json:"packageSize,omitempty"`
	Volumn               int64   `json:"volumn,omitempty"`
	HandlingTime         int     `json:"handlingTime,omitempty"`
	SendGoodsAddressId   uint64  `json:"sendGoodsAddressId,omitempty"`
	SendGoodsAddressText string  `json:"sendGoodsAddressText,omitempty"`
	SuttleWeight         float64 `json:"suttleWeight,omitempty"`
	Height               float64 `json:"height,omitempty"`
	Width                float64 `json:"width,omitempty"`
	Length               float64 `json:"length,omitempty"`
}

type Attribute struct {
	AttributeId uint64 `json:"attributeID,omitempty"`
	Value       string `json:"value,omitempty"`
	ValueId     uint64 `json:"valueID,omitempty"`
	Name        string `json:"attributeName,omitempty"`
	IsCustom    bool   `json:"isCustom,omitempty"`
}

type BizGroupInfo struct {
	// Support 是否支持,isConsignMarketOffer=ture表示代销offer
	Support     bool   `json:"support,omitempty"`
	Description string `json:"description,omitempty"`
	Code        string `json:"code,omitempty"`
}
