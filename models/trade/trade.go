package trade

import (
	"github.com/shopspring/decimal"
	"github.com/wzhsh90/go1688/models"
	"github.com/wzhsh90/go1688/utils"
)

type Address struct {
	AddressId uint64 `json:"addressId,omitempty"`
	//名称至少到2位长度
	FullName     string `json:"fullName,omitempty"`
	Mobile       string `json:"mobile,omitempty"`
	Phone        string `json:"phone,omitempty"`
	PostCode     string `json:"postCode,omitempty"`
	CityText     string `json:"cityText,omitempty"`
	ProvinceText string `json:"provinceText,omitempty"`
	AreaText     string `json:"areaText,omitempty"`
	TownText     string `json:"townText,omitempty"`
	Address      string `json:"address,omitempty"`
	DistrictCode string `json:"districtCode,omitempty"`
}

type Cargo struct {
	OfferId     uint64  `json:"offerId,omitempty"`
	SpecId      string  `json:"specId,omitempty"`
	Quantify    float64 `json:"quantity,omitempty"`
	Channel     string  `json:"channel,omitempty"`
	OutItemCode string  `json:"outItemCode,omitempty"`
	OutShopCode string  `json:"outShopCode,omitempty"`
}

type OuterOrderInfo struct {
	MediaOrderId string            `json:"mediaOrderId,omitempty"`
	Phone        string            `json:"phone,omitempty"`
	Offers       []OuterOrderOffer `json:"offers,omitempty"`
}

type OuterOrderOffer struct {
	Id     uint64 `json:"id,omitempty"`
	SpecId string `json:"specId,omitempty"`
	Price  int64  `json:"price,omitempty"`
	Num    int64  `json:"num,omitempty"`
}

type CreateOrderResult struct {
	TotalSuccessAmount int64  `json:"totalSuccessAmount,omitempty"`
	OrderId            string `json:"orderId,omitempty"`
	PostFee            int64  `json:"postFee,omitempty"`
}

type OrderPreview struct {
	DiscountFee              uint        `json:"discountFee,omitempty"`
	TradeModeNameList        []string    `json:"tradeModeNameList,omitempty"`
	Status                   bool        `json:"status,omitempty"`
	TaoSampleSinglePromotion bool        `json:"taoSampleSinglePromotion,omitempty"`
	SumPayment               int64       `json:"sumPayment,omitempty"`
	Message                  string      `json:"message,omitempty"`
	SumCarriage              int64       `json:"sumCarriage,omitempty"`
	ResultCode               string      `json:"resultCode,omitempty"`
	SumPaymentNoCarriage     int64       `json:"sumPaymentNoCarriage,omitempty"`
	AdditionalFee            int64       `json:"additionalFee,omitempty"`
	FlowFlag                 string      `json:"flowFlag,omitempty"`
	CargoList                []CargoInfo `json:"cargoList,omitempty"`
	ShopPromotionList        []Promotion `json:"shopPromotionList,omitempty"`
	TradeModeList            []TradeMode `json:"tradeModeList,omitempty"`
}

type CargoInfo struct {
	OfferId        uint64          `json:"offerId,omitempty"`
	Amount         decimal.Decimal `json:"amount,omitempty"`
	Message        string          `json:"message,omitempty"`
	FinalUnitPrice decimal.Decimal `json:"finalUnitPrice,omitempty"`
	SpecId         string          `json:"specId,omitempty"`
	SkuId          uint64          `json:"skuId,omitempty"`
	ResultCode     string          `json:"resultCode,omitempty"`
	PromotionList  []Promotion     `json:"cargoPromotionList,omitempty"`
}

type Promotion struct {
	Id          string `json:"promotionId,omitempty"`
	Selected    bool   `json:"selected,omitempty"`
	Text        string `json:"text,omitempty"`
	Desc        string `json:"desc,omitempty"`
	FreePostage bool   `json:"freePostage,omitempty"`
	DiscountFee int64  `json:"discountFee,omitempty"`
}

type TradeMode struct {
	TradeWay    string `json:"tradeWay,omitempty"`
	Name        string `json:"name,omitempty"`
	TradeType   string `json:"tradeType,omitempty"`
	Description string `json:"description,omitempty"`
	OpSupport   bool   `json:"opSupport,omitempty"`
}

type GetBuyerViewReq struct {
	Website       string   `json:"webSite,omitempty"`
	OrderId       uint64   `json:"orderId"`
	IncludeFields string   `json:"includeFields,omitempty"`
	AttributeKeys []string `json:"attributeKeys,omitempty"`
}

func (r GetBuyerViewReq) Name() string {
	return "alibaba.trade.get.buyerView"
}

func (r GetBuyerViewReq) Params() map[string]string {
	ret := make(map[string]string, 4)
	ret["webSite"] = r.Website
	ret["orderId"] = utils.FormatUint(r.OrderId)
	if r.IncludeFields != "" {
		ret["includeFields"] = r.IncludeFields
	}
	if len(r.AttributeKeys) > 0 {
		ret["attributeKeys"] = utils.AliMarshal(r.AttributeKeys)
	}
	return ret
}

type GetBuyerViewResp struct {
	models.BaseStrResponse
	Result *Info `json:"result"`
}

type Info struct {
	BaseInfo             *OrderBaseInfo        `json:"baseInfo,omitempty"`
	BizInfo              *OrderBizInfo         `json:"orderBizInfo,omitempty"`
	TradeTerms           []TradeTermInfo       `json:"tradeTerms,omitempty"`
	ProductItems         []ProductItemInfo     `json:"productItems,omitempty"`
	NativeLogistics      *NativeLogisticsInfo  `json:"nativeLogistics,omitempty"`
	InvoiceInfo          *OrderInvoiceInfo     `json:"orderInvoiceInfo,omitempty"`
	GuaranteesTerms      *GuaranteesTermInfo   `json:"guaranteesTerms,omitempty"`
	OrderRateInfo        *OrderRateInfo        `json:"orderRateInfo,omitempty"`
	OverseasExtraAddress *OverseasExtraAddress `json:"overseasExtraAddress,omitempty"`
	Customs              *OrderCustoms         `json:"customs,omitempty"`
	QuoteList            []CaigouQuoteInfo     `json:"quoteList,omitempty"`
	ExtAttributes        []KeyValuePair        `json:"extAttributes,omitempty"`
}

type OrderBaseInfo struct {
	AllDeliveredTime  string          `json:"allDeliveredTime,omitempty"`
	SellerCreditLevel string          `json:"sellerCreditLevel,omitempty"`
	PayTime           string          `json:"payTime,omitempty"`
	Discount          int64           `json:"discount,omitempty"`
	AlipayTradeId     string          `json:"alipayTradeId,omitempty"`
	SumProductPayment decimal.Decimal `json:"sumProductPayment,omitempty"`
	BuyerFeedback     string          `json:"buyerFeedback,omitempty"`
	FlowTemplateCode  string          `json:"flowTemplateCode,omitempty"`
	SellerOrder       bool            `json:"sellerOrder,omitempty"`
	BuyerLoginId      string          `json:"buyerLoginId,omitempty"`
	ModifyTime        string          `json:"modifyTime,omitempty"`
	SubBuyerLoginId   string          `json:"subBuyerLoginId,omitempty"`
	Id                uint64          `json:"id,omitempty"`
	CloseReason       string          `json:"closeReason,omitempty"`
	BuyerContact      *Contact        `json:"buyerContact,omitempty"`
	SellerAlipayId    string          `json:"sellerAlipayId,omitempty"`
	CompleteTime      string          `json:"completeTime,omitempty"`
	SellerLoginId     string          `json:"sellerLoginId,omitempty"`
	BuyerId           string          `json:"buyerID,omitempty"`
	CloseOperateType  string          `json:"closeOperateType,omitempty"`
	TotalAmount       decimal.Decimal `json:"totalAmount,omitempty"`
	SellerId          string          `json:"sellerID,omitempty"`
	ShippingFee       decimal.Decimal `json:"shippingFee,omitempty"`
	BuyerUserId       uint64          `json:"buyerUserId,omitempty"`
	BuyerMemo         string          `json:"buyerMemo,omitempty"`
	Refund            decimal.Decimal `json:"refund,omitempty"`
	Status            string          `json:"status,omitempty"`
	RefundPayment     int64           `json:"refundPayment,omitempty"`
	SellerContact     *Contact        `json:"sellerContact,omitempty"`
	RefundStatus      string          `json:"refundStatus,omitempty"`
	Remark            string          `json:"remark,omitempty"`
	PreOrderId        uint64          `json:"preOrderId,omitempty"`
	ConfirmedTime     string          `json:"confirmedTime,omitempty"`
	CloseRemark       string          `json:"closeRemark,omitempty"`
	TradeType         string          `json:"tradeType,omitempty"`
	ReceivingTime     string          `json:"receivingTime,omitempty"`
	StepAgreementPath string          `json:"stepAgreementPath,omitempty"`
	IdOfStr           string          `json:"idOfStr,omitempty"`
	RefundStatusForAs string          `json:"refundStatusForAs,omitempty"`
	StepPayAll        bool            `json:"stepPayAll,omitempty"`
	SellerUserId      uint64          `json:"sellerUserId,omitempty"`
	StepOrderList     []StepOrder     `json:"stepOrderList,omitempty"`
	BuyerAlipayId     string          `json:"buyerAlipayId,omitempty"`
	CreateTime        string          `json:"createTime,omitempty"`
	BusinessType      string          `json:"businessType,omitempty"`
	OverSeaOrder      bool            `json:"overSeaOrder,omitempty"`
	RefundId          string          `json:"refundId,omitempty"`
	TradeTypeDesc     string          `json:"tradeTypeDesc,omitempty"`
	PayChannelList    []string        `json:"payChannelList,omitempty"`
	TradeTypeCode     string          `json:"tradeTypeCode,omitempty"`
	PayTimeout        int64           `json:"payTimeout,omitempty"`
	PayTimeoutType    uint            `json:"payTimeoutType,omitempty"`
}

type Contact struct {
	Phone         string `json:"phone,omitempty"`
	Fax           string `json:"fax,omitempty"`
	Email         string `json:"email,omitempty"`
	ImInPlatform  string `json:"imInPlatform,omitempty"`
	Name          string `json:"name,omitempty"`
	Mobile        string `json:"mobile,omitempty"`
	CompanyName   string `json:"companyName,omitempty"`
	WgSenderName  string `json:"wgSenderName,omitempty"`
	WgSenderPhone string `json:"wgSenderPhone,omitempty"`
}

type StepOrder struct {
	StepOrderId          uint64          `json:"stepOrderId,omitempty"`
	Status               string          `json:"stepOrderStatus,omitempty"`
	PayStatus            int             `json:"stepPayStatus,omitempty"`
	No                   int             `json:"stepNo,omitepty"`
	LastStep             bool            `json:"lastStep,omitempty"`
	HasDisbursed         bool            `json:"hasDisbursed,omitempty"`
	PayFee               decimal.Decimal `json:"payFee,omitempty"`
	ActualPayFee         decimal.Decimal `json:"actualPayFee,omitempty"`
	DiscountFee          decimal.Decimal `json:"discountFee,omitempty"`
	ItemDiscountFee      decimal.Decimal `json:"itemDiscountFee,omitempty"`
	Price                decimal.Decimal `json:"price,omitempty"`
	Amount               int             `json:"amount,omitempty"`
	PostFee              decimal.Decimal `json:"postFee,omitempty"`
	AdjustPostFee        decimal.Decimal `json:"adjustPostFee,omitempty"`
	GmtCreate            string          `json:"gmtCreate,omitempty"`
	GmtModified          string          `json:"gmtModified,omitempty"`
	EnterTime            string          `json:"enterTime,omitempty"`
	PayTime              string          `json:"payTime,omitempty"`
	SellerActionTime     string          `json:"sellerActionTime,omitempty"`
	EndTime              string          `json:"endTime,omitempty"`
	MessagePath          string          `json:"messagePath,omitempty"`
	PicturePath          string          `json:"picturePath,omitempty"`
	Message              string          `json:"message,omitempty"`
	TemplateId           uint64          `json:"templateId,omitempty"`
	Name                 string          `json:"stepName,omitempty"`
	SellerActionName     string          `json:"sellerActionName,omitempty"`
	BuyerPayTimeout      int64           `json:"buyerPayTimeout,omitempty"`
	BuyerConfirmTimeout  int64           `json:"buyerConfirmTimeout,omitempty"`
	NeedLogistics        bool            `json:"needLogistics,omitempty"`
	NeedSellerAction     bool            `json:"needSellerAction,omitempty"`
	TransferAfterConfirm bool            `json:"transferAfterConfirm,omitempty"`
	NeedSellerCallNext   bool            `json:"needSellerCallNext,omitempty"`
	InstantPay           bool            `json:"instantPay,omitempty"`
}

type OrderBizInfo struct {
	OdsCyd            bool               `json:"odsCyd,omitempty"`
	AccountPeriodTime string             `json:"accountPeriodTime,omitempty"`
	CreditOrder       bool               `json:"creditOrder,omitempty"`
	CreditOrderDetail *CreditOrderDetail `json:"creditOrderDetail,omitempty"`
	PreOrderInfo      *PreOrderInfo      `json:"preOrderInfo,omitempty"`
	LstOrderInfo      *LstOrderInfo      `json:"lstOrderInfo,omitempty"`
}

type CreditOrderDetail struct {
	PayAmount          int64  `json:"payAmount,omitempty"`
	CreateTime         string `json:"createTime,omitempty"`
	Status             string `json:"status,omitempty"`
	GracePeriodEndTime string `json:"gracePeriodEndTime,omitempty"`
	StatusStr          string `json:"statusStr,omitempty"`
	RestRepayAmount    int64  `json:"restRepayAmount,omitempty"`
}

// PreOrderInfo 预订单信息
type PreOrderInfo struct {
	MarketName        string `json:"marketName,omitempty"`
	CreatePreOrderApp bool   `json:"createPreOrderApp,omitempty"`
}

type LstOrderInfo struct {
	LstWarehouseType string `json:"lstWarehouseType,omitempty"`
}

// TradeTermInfo 交易条款
type TradeTermInfo struct {
	PayStatus      string          `json:"payStatus,omitempty"`
	PayTime        string          `json:"payTime,omitempty"`
	PayWay         string          `json:"payWay,omitempty"`
	PhasAmount     decimal.Decimal `json:"phasAmount,omitempty"`
	Phase          uint64          `json:"phase,omitempty"`
	PhaseCondition string          `json:"phaseCondition,omitempty"`
	PhaseDate      string          `json:"phaseDate,omitempty"`
	CardPay        bool            `json:"cardPay,omitempty"`
	ExpressPay     bool            `json:"expressPay,omitempty"`
	PayWayDesc     string          `json:"payWayDesc,omitempty"`
}

// ProductItemInfo 商品条目信息
type ProductItemInfo struct {
	CargoNumber        string               `json:"cargoNumber,omitempty"`
	Description        string               `json:"description,omitempty"`
	ItemAmount         decimal.Decimal      `json:"itemAmount,omitempty"`
	Name               string               `json:"name,omitempty"`
	Price              decimal.Decimal      `json:"price,omitempty"`
	ProductId          uint64               `json:"productID,omitempty"`
	ImgURL             []string             `json:"productImgUrl,omitempty"`
	SnapshotURL        string               `json:"productSnapshotUrl,omitempty"`
	Quantity           decimal.Decimal      `json:"quantity,omitempty"`
	Refund             decimal.Decimal      `json:"refund,omitempty"`
	SkuId              uint64               `json:"skuID,omitempty"`
	Sort               int                  `json:"sort,omitempty"`
	Status             string               `json:"status,omitempty"`
	SubItemId          uint64               `json:"subItemId,omitempty"`
	Type               string               `json:"type,omitempty"`
	Unit               string               `json:"unit,omitempty"`
	Weight             string               `json:"weight,omitempty"`
	WeightUnit         string               `json:"weightUnit,omitempty"`
	GuaranteesTerms    []GuaranteesTermInfo `json:"guaranteesTerms,omitempty"`
	ProductCargoNumber string               `json:"productCargoNumber,omitempty"`
	SkuInfos           []SkuItemDesc        `json:"skuInfos,omitempty"`
	EntryDiscount      int64                `json:"entryDiscount,omitempty"`
	SpecId             string               `json:"specId,omitempty"`
	QuantityFactor     decimal.Decimal      `json:"quantityFactor,omitempty"`
	StatusStr          string               `json:"statusStr,omitempty"`
	RefundStatus       string               `json:"refundStatus,omitempty"`
	CloseReason        string               `json:"closeReason,omitempty"`
	LogisticsStatus    int                  `json:"logisticsStatus,omitempty"`
	GmtCreate          string               `json:"gmtCreate,omitempty"`
	GmtModified        string               `json:"gmtModified,omitempty"`
	GmtCompleted       string               `json:"gmtCompleted,omitempty"`
	GmtPayExpireTime   string               `json:"gmtPayExpireTime,omitempty"`
	RefundId           string               `json:"refundId,omitempty"`
	SubItemIdString    string               `json:"subItemIDString,omitempty"`
	RefundIdForAs      string               `json:"refundIdForAs,omitempty"`
}

type SkuItemDesc struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type GuaranteesTermInfo struct {
	AssuranceInfo        string `json:"assuranceInfo,omitempty"`
	AssuranceType        string `json:"assuranceType,omitempty"`
	QualityAssuranceType string `json:"qualityAssuranceType,omitempty"`
}

type NativeLogisticsInfo struct {
	Address        string                    `json:"address,omitempty"`
	Area           string                    `json:"area,omitempty"`
	AreaCode       string                    `json:"areaCode,omitempty"`
	City           string                    `json:"city,omitempty"`
	ContactPerson  string                    `json:"contactPerson,omitempty"`
	Fax            string                    `json:"fax,omitempty"`
	Mobile         string                    `json:"mobile,omitempty"`
	Province       string                    `json:"province,omitempty"`
	Telephone      string                    `json:"telephone,omitempty"`
	Zip            string                    `json:"zip,omitempty"`
	LogisticsItems []NativeLogisticsItemInfo `json:"logisticsItems,omitempty"`
	TownCode       string                    `json:"townCode,omitempty"`
	Town           string                    `json:"town,omitempty"`
}

type NativeLogisticsItemInfo struct {
	DeliveredTime        string          `json:"deliveredTime,omitempty"`
	LogisticsCode        string          `json:"logisticsCode,omitempty"`
	Type                 string          `json:"type,omitempty"`
	Id                   uint64          `json:"id,omitempty"`
	Status               string          `json:"status,omitempty"`
	GmtModified          string          `json:"gmtModified,omitempty"`
	GmtCreate            string          `json:"gmtCreate,omitempty"`
	Carriage             decimal.Decimal `json:"carriage,omitempty"`
	FromProvince         string          `json:"fromProvince,omitempty"`
	FromCity             string          `json:"fromCity,omitempty"`
	FromArea             string          `json:"fromArea,omitempty"`
	FromAddress          string          `json:"fromAddress,omitempty"`
	FromPhone            string          `json:"fromPhone,omitempty"`
	FromMobile           string          `json:"fromMobile,omitempty"`
	FromPost             string          `json:"fromPost,omitempty"`
	LogisticsCompanyId   uint64          `json:"logisticsCompanyId,omitempty"`
	LogisticesCompanyNo  string          `json:"logisticsCompanyNo,omitempty"`
	LogisticsCompanyName string          `json:"logisticsCompanyName,omitempty"`
	LogisticsBillNo      string          `json:"logisticsBillNo,omitempty"`
	SubItemIds           string          `json:"subItemIds,omitempty"`
	ToProvince           string          `json:"toProvince,omitempty"`
	ToCity               string          `json:"toCity,omitempty"`
	ToArea               string          `json:"toArea,omitempty"`
	ToAddress            string          `json:"toAddress,omitempty"`
	ToPhone              string          `json:"toPhone,omitempty"`
	ToMobile             string          `json:"toMobile,omitempty"`
	ToPost               string          `json:"toPost,omitempty"`
	NoLogisticsName      string          `json:"noLogisticsName,omitempty"`
	NoLogisticsTel       string          `json:"noLogisticsTel,omitempty"`
	NoLogisticsBillNo    string          `json:"noLogisticsBillNo,omitempty"`
	NoLogisticsCondition string          `json:"noLogisticsCondition,omitempty"`
	IsTimePromise        bool            `json:"isTimePromise,omitempty"`
	ArriveTime           string          `json:"arriveTime,omitempty"`
}

// OrderInvoiceInfo 发票信息
type OrderInvoiceInfo struct {
	InvoiceCompanyName string `json:"invoiceCompanyName,omitempty"`
	InvoiceType        int    `json:"invoiceType,omitempty"`
	LocalInvoiceId     uint64 `json:"localInvoiceId,omitempty"`
	OrderId            uint64 `json:"orderId,omitempty"`
	ReceiveCode        string `json:"receiveCode,omitempty"`
	ReceiveCodeText    string `json:"receiveCodeText,omitempty"`
	ReceiveMobile      string `json:"receiveMobile,omitempty"`
	ReceiveName        string `json:"receiveName,omitempty"`
	ReceivePhone       string `json:"receivePhone,omitempty"`
	ReceivePost        string `json:"receivePost,omitempty"`
	ReceiveStreet      string `json:"receiveStreet,omitempty"`
	RegisterAccountId  string `json:"registerAccountId,omitempty"`
	RegisterBank       string `json:"registerBank,omitempty"`
	RegisterCode       string `json:"registerCode,omitempty"`
	RegisterCodeText   string `json:"registerCodeText,omitempty"`
	RegisterPhone      string `json:"registerPhone,omitempty"`
	RegisterStreet     string `json:"registerStreet,omitempty"`
	TaxPayerIdentify   string `json:"taxpayerIdentify,omitempty"`
}

type OrderRateInfo struct {
	BuyerRateStatus  int               `json:"buyerRateStatus,omitempty"`
	SellerRateStatus int               `json:"sellerRateStatus,omitempty"`
	BuyerRateList    []OrderRateDetail `json:"buyerRateList,omitempty"`
	SellerRateList   []OrderRateDetail `json:"sellerRateList,omitempty"`
}

type OrderRateDetail struct {
	StarLevel    int    `json:"starLevel,omitempty"`
	Content      string `json:"content,omitempty"`
	ReceiverNick string `json:"receiverNick,omitempty"`
	PosterNick   string `json:"posterNick,omitempty"`
	PublishTime  string `json:"publishTime,omitempty"`
}

type OverseasExtraAddress struct {
	ChannelName         string `json:"channelName,omitempty"`
	ChannelId           string `json:"channelId,omitempty"`
	ShippingCompanyId   string `json:"shippingCompanyId,omitempty"`
	ShippingCompanyName string `json:"shoppingCompanyName,omitempty"`
	CountryCode         string `json:"countryCode,omitempty"`
	Country             string `json:"country,omitempty"`
	Email               string `json:"email,omitempty"`
}

type OrderCustoms struct {
	Id          uint64                  `json:"id,omitempty"`
	GmtCreate   string                  `json:"gmtCreate,omitempty"`
	GmtModified string                  `json:"gmtModified,omitempty"`
	BuyerId     uint64                  `json:"buyerId,omitempty"`
	OrderId     uint64                  `json:"orderId,omitempty"`
	Type        int                     `json:"type,omitempty"`
	Attributes  []CustomerAttributeInfo `json:"attributes,omitempty"`
}

type CustomerAttributeInfo struct {
	Sku      string  `json:"sku,omitempty"`
	CnName   string  `json:"cnName,omitempty"`
	EnName   string  `json:"enName,omitempty"`
	Amount   float64 `json:"amount,omitempty"`
	Quantity float64 `json:"quantify,omitempty"`
	Weight   float64 `json:"weight,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

type CaigouQuoteInfo struct {
	ProductQuoteName string          `json:"productQuoteName,omitempty"`
	Price            decimal.Decimal `json:"price,omitempty"`
	Count            float64         `json:"count,omitempty"`
}

type KeyValuePair struct {
	Key         string `json:"key,omitempty"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
}
type CreateOrderResp struct {
	models.BaseBoolResponse
	Result *CreateOrderResult `json:"result,omitempty"`
}

type CreateOrderPreviewResp struct {
	models.BaseBoolResponse
	PostFeeByDescOfferList []uint64       `json:"postFeeByDescOfferList,omitempty"`
	ConsignOfferList       []uint64       `json:"consignOfferList,omitempty"`
	OrderPreview           []OrderPreview `json:"orderPreviewResuslt,omitempty"`
}
