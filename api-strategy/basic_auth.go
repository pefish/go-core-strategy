package api_strategy

import (
	"fmt"
	"strings"

	i_core "github.com/pefish/go-interface/i-core"
	t_error "github.com/pefish/go-interface/t-error"
)

type BasicAuthParams struct {
	Username string
	Password string
}

type BasicAuthStrategy struct {
	errorCode uint64
	errorMsg  string
	params    *BasicAuthParams
}

func NewBasicAuthStrategy() *BasicAuthStrategy {
	return &BasicAuthStrategy{}
}

func (b *BasicAuthStrategy) Name() string {
	return `BasicAuthStrategy`
}

func (b *BasicAuthStrategy) Description() string {
	return `basic auth`
}

func (b *BasicAuthStrategy) SetErrorCode(code uint64) i_core.IApiStrategy {
	b.errorCode = code
	return b
}

func (b *BasicAuthStrategy) SetParams(params *BasicAuthParams) *BasicAuthStrategy {
	b.params = params
	return b
}

func (b *BasicAuthStrategy) SetErrorMsg(msg string) i_core.IApiStrategy {
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
		return t_error.INTERNAL_ERROR_CODE
	}
	return b.errorCode
}

func (b *BasicAuthStrategy) Execute(out i_core.IApiSession) *t_error.ErrorInfo {
	out.Logger().DebugF(`Api strategy %s trigger`, b.Name())

	u, p, ok := out.Request().BasicAuth()
	if !ok || !strings.EqualFold(b.params.Username, u) || !strings.EqualFold(b.params.Password, p) {
		return t_error.WrapWithAll(fmt.Errorf(b.ErrorMsg()), b.ErrorCode(), nil)
	}

	return nil
}
