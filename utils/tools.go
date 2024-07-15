package utils

import (
	"sort"
	"strconv"
)

func FormatUint(val uint64) string {

	return strconv.FormatUint(val, 10)
}

func Combine(params map[string]string) string {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	builder := GetStringsBuilder()
	defer PutStringsBuilder(builder)
	for _, k := range keys {
		builder.WriteString(k)
		builder.WriteString(params[k])
	}
	return builder.String()
}
func Sign(path string, appSecret []byte, params map[string]string) string {
	builder := GetBytesBuffer()
	defer PutBytesBuffer(builder)
	builder.WriteString(path)
	builder.WriteString(Combine(params))
	return HmacSha1Upper(appSecret, builder.Bytes())
}
