package message

import (
	"github.com/wzhsh90/go1688/utils"
)

type Template struct {
	ID        uint64 `json:"msgId"`
	GmtBorn   int64  `json:"gmtBorn"`
	Data      string `json:"data"`
	UserInfo  string `json:"userInfo"`
	Type      string `json:"type"`
	ExtraInfo string `json:"extraInfo"`
}

type NotifyMsg struct {
	// Message 对接方获取这个参数的值，然后通过json的反序列化，就得到了消息模型
	Message string `form:"message" json:"message" binding:"required"`
	// Sign 针对消息的一个签名，可防篡改
	Sign string `form:"_aop_signature" json:"_aop_signature" binding:"required"`
}

func (c *NotifyMsg) Valid(appSecret []byte) bool {
	return c.Sign == utils.Sign("", appSecret, map[string]string{"message": c.Message})
}
