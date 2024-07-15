package models

type Response interface {
	GetErrorMsg() string
	IsError() bool
}

// 错误码在不同情况下返回不一样
// 1、签名错误时返回error_code 2、业务单号不存在是返回errorCode
// error_message/errorMsg/errorMessage
// {"errorMessage":"该订单没有物流跟踪信息。","success":true} 物流追踪返回没有errorCode
// {"success":false,"code":"500","message":"不是自己的订单"} PreparePay
type BaseResponse struct {
	ErrCode      string `json:"error_code,omitempty"`
	ErrMessage   string `json:"error_message,omitempty"`
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMsg     string `json:"errorMsg,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (c *BaseResponse) IsError() bool {
	return c.ErrMessage != "" || c.ErrorMessage != "" || c.ErrorMsg != ""
}
func (c *BaseResponse) GetErrorMsg() string {
	if c.ErrMessage != "" {
		return c.ErrMessage
	} else if c.ErrorMessage != "" {
		return c.ErrorMessage
	} else if c.ErrorMsg != "" {
		return c.ErrorMsg
	}
	return ""
}

type BaseBoolResponse struct {
	BaseResponse
	Success bool `json:"success,omitempty"`
}
type BaseStrResponse struct {
	BaseResponse
	Success string `json:"success,omitempty"`
}
