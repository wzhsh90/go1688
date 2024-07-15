package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

func HmacSha1Upper(key []byte, message []byte) string {
	h := hmac.New(sha1.New, key)
	h.Write(message)
	sha := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(sha))
}
