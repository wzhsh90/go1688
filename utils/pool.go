package utils

import (
	"bytes"
	"encoding/json"
	"strings"
	"sync"
)

var stringsBuilderPool = sync.Pool{
	New: func() any {
		return new(strings.Builder)
	},
}

func GetStringsBuilder() *strings.Builder {
	builder := stringsBuilderPool.Get().(*strings.Builder)
	builder.Reset()
	return builder
}

func PutStringsBuilder(builder *strings.Builder) {
	builder.Reset()
	stringsBuilderPool.Put(builder)
}

var bytesBufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func GetBytesBuffer() *bytes.Buffer {
	buf := bytesBufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

func PutBytesBuffer(buf *bytes.Buffer) {
	bytesBufferPool.Put(buf)
}

// JSONMarshal encode json without html escape
func AliMarshal(req interface{}) string {
	buf := GetBytesBuffer()
	defer PutBytesBuffer(buf)
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	encoder.Encode(req)
	return buf.String()
}
