package api_strategy

import (
	"fmt"

	i_core "github.com/pefish/go-interface/i-core"
	t_error "github.com/pefish/go-interface/t-error"
)

type IpFilterStrategy struct {
	errorCode uint64
	errorMsg  string
	params    *IpFilterParams
}

func NewIpFilterStrategy() *IpFilterStrategy {
	return &IpFilterStrategy{}
}

type IpFilterParams struct {
	ValidIp func(apiSession i_core.IApiSession) []string
}

func (ifs *IpFilterStrategy) Name() string {
	return `IpFilterStrategy`
}

func (ifs *IpFilterStrategy) Description() string {
	return `filter ip`
}

func (b *IpFilterStrategy) SetParams(params *IpFilterParams) *IpFilterStrategy {
	b.params = params
	return b
}

func (ifs *IpFilterStrategy) SetErrorCode(code uint64) i_core.IApiStrategy {
	ifs.errorCode = code
	return ifs
}

func (ifs *IpFilterStrategy) ErrorCode() uint64 {
	if ifs.errorCode == 0 {
		return t_error.INTERNAL_ERROR_CODE
	}
	return ifs.errorCode
}

func (ifs *IpFilterStrategy) SetErrorMsg(msg string) i_core.IApiStrategy {
	ifs.errorMsg = msg
	return ifs
}

func (ifs *IpFilterStrategy) ErrorMsg() string {
	if ifs.errorMsg == "" {
		return "Ip is baned."
	}
	return ifs.errorMsg
}

func (ifs *IpFilterStrategy) Execute(out i_core.IApiSession) *t_error.ErrorInfo {
	out.Logger().DebugF(`Api strategy %s trigger`, ifs.Name())
	if ifs.params.ValidIp == nil {
		return nil
	}
	clientIp := out.RemoteAddress()
	allowedIps := ifs.params.ValidIp(out)
	for _, ip := range allowedIps {
		if ip == clientIp {
			return nil
		}
	}
	return t_error.WrapWithAll(fmt.Errorf(ifs.ErrorMsg()), ifs.ErrorCode(), nil)
}
