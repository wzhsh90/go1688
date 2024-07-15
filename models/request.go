package models

import (
	"github.com/wzhsh90/go1688/consts"
	"github.com/wzhsh90/go1688/utils"
)

type Request interface {
	Namespace() string
	Name() string
	Version() string
	Path() string
	Params() map[string]string
}

type RequestData interface {
	Name() string
	Params() map[string]string
}
type BaseRequest struct {
	namespace string
	name      string
	version   string
	params    map[string]string
}

func (r *BaseRequest) SetVersion(version string) {
	r.version = version
}

func (r *BaseRequest) Name() string {
	return r.name
}

func (r *BaseRequest) Namespace() string {
	return r.namespace
}

func (r *BaseRequest) Version() string {
	if r.version == "" {
		return consts.VERSION
	}
	return r.version
}
func (r *BaseRequest) Params() map[string]string {
	return r.params
}

func (r *BaseRequest) Path() string {
	builder := utils.GetStringsBuilder()
	defer utils.PutStringsBuilder(builder)
	builder.WriteString(r.Version())
	builder.WriteString("/")
	builder.WriteString(r.Namespace())
	builder.WriteString("/")
	builder.WriteString(r.Name())
	return builder.String()
}

func NewRequest(namespace string, req RequestData) *BaseRequest {
	return &BaseRequest{
		namespace: namespace,
		name:      req.Name(),
		params:    req.Params(),
	}
}
