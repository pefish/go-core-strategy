package global_api_strategy

import (
	_type "github.com/pefish/go-core-type/api-session"
	go_error "github.com/pefish/go-error"
)

type IGlobalStrategy interface {
	Init(param interface{})
	Execute(out _type.IApiSession, param interface{}) *go_error.ErrorInfo
	GetName() string
	GetDescription() string
	GetErrorCode() uint64
	SetErrorCode(code uint64) IGlobalStrategy
	SetErrorMsg(msg string) IGlobalStrategy
	GetErrorMsg() string
}
