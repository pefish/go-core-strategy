package api_strategy

import (
	"fmt"
	"strings"

	api_session "github.com/pefish/go-core-type/api-session"
	api_strategy "github.com/pefish/go-core-type/api-strategy"
	go_error "github.com/pefish/go-error"
)

type BasicAuthStrategy struct {
	errorCode uint64
	errorMsg  string
}

func NewBasicAuthStrategy() *BasicAuthStrategy {
	return &BasicAuthStrategy{}
}

func (b *BasicAuthStrategy) Init(param interface{}) api_strategy.IApiStrategy {
	return b
}

func (b *BasicAuthStrategy) Name() string {
	return `BasicAuthStrategy`
}

func (b *BasicAuthStrategy) Description() string {
	return `basic auth`
}

func (b *BasicAuthStrategy) SetErrorCode(code uint64) api_strategy.IApiStrategy {
	b.errorCode = code
	return b
}

func (b *BasicAuthStrategy) SetErrorMsg(msg string) api_strategy.IApiStrategy {
	b.errorMsg = msg
	return b
}

func (b *BasicAuthStrategy) ErrorMsg() string {
	if b.errorMsg == "" {
		return "Unauthorized."
	}
	return b.errorMsg
}

func (b *BasicAuthStrategy) ErrorCode() uint64 {
	if b.errorCode == 0 {
		return go_error.INTERNAL_ERROR_CODE
	}
	return b.errorCode
}

type BasicAuthParam struct {
	Username string
	Password string
}

func (b *BasicAuthStrategy) Execute(out api_session.IApiSession, param interface{}) *go_error.ErrorInfo {
	out.Logger().DebugF(`Api strategy %s trigger`, b.Name())
	if param == nil {
		out.Logger().ErrorF(`Strategy need param.`)
		return go_error.WrapWithAll(fmt.Errorf(b.ErrorMsg()), b.ErrorCode(), nil)
	}
	params := param.(BasicAuthParam)

	u, p, ok := out.Request().BasicAuth()
	if !ok || !strings.EqualFold(params.Username, u) || !strings.EqualFold(params.Password, p) {
		return go_error.WrapWithAll(fmt.Errorf(b.ErrorMsg()), b.ErrorCode(), nil)
	}

	return nil
}
