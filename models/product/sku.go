package product

import "github.com/shopspring/decimal"

type SkuAttribute struct {
	AttributeId uint64 `json:"attributeID,omitempty"`
	Value       string `json:"attributeValue,omitempty"`
	ImageURL    string `json:"skuImageUrl,omitempty"`
	DisplayName string `json:"attributeDisplayName,omitempty"`
	Name        string `json:"attributeName,omitempty"`
}

type SkuInfo struct {
	SkuId           uint64          `json:"skuId,omitempty"`
	Attributes      []*SkuAttribute `json:"attributes,omitempty"`
	CargoNumber     string          `json:"cargoNumber,omitempty"`
	AmountOnSale    int64           `json:"amountOnSale,omitempty"`
	RetailPrice     decimal.Decimal `json:"retailPrice,omitempty"`
	Price           decimal.Decimal `json:"price,omitempty"`
	SpecId          string          `json:"specId,omitempty"`
	ConsignPrice    decimal.Decimal `json:"consignPrice,omitempty"`
	CpsSuggestPrice decimal.Decimal `json:"cpsSuggestPrice,omitempty"`
	ChannelPrice    decimal.Decimal `json:"channelPrice,omitempty"`
}
