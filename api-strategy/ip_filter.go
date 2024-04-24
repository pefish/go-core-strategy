package api_strategy

import (
	"fmt"

	api_session "github.com/pefish/go-core-type/api-session"
	api_strategy "github.com/pefish/go-core-type/api-strategy"
	go_error "github.com/pefish/go-error"
)

type IpFilterStrategy struct {
	errorCode uint64
	errorMsg  string
	params    IpFilterParams
}

func NewIpFilterStrategy() *IpFilterStrategy {
	return &IpFilterStrategy{}
}

type IpFilterParams struct {
	ValidIp func(apiSession api_session.IApiSession) []string
}

func (ifs *IpFilterStrategy) Name() string {
	return `IpFilterStrategy`
}

func (ifs *IpFilterStrategy) Description() string {
	return `filter ip`
}

func (b *IpFilterStrategy) SetParams(params IpFilterParams) api_strategy.IApiStrategy {
	b.params = params
	return b
}

func (ifs *IpFilterStrategy) SetErrorCode(code uint64) api_strategy.IApiStrategy {
	ifs.errorCode = code
	return ifs
}

func (ifs *IpFilterStrategy) ErrorCode() uint64 {
	if ifs.errorCode == 0 {
		return go_error.INTERNAL_ERROR_CODE
	}
	return ifs.errorCode
}

func (ifs *IpFilterStrategy) SetErrorMsg(msg string) api_strategy.IApiStrategy {
	ifs.errorMsg = msg
	return ifs
}

func (ifs *IpFilterStrategy) ErrorMsg() string {
	if ifs.errorMsg == "" {
		return "Ip is baned."
	}
	return ifs.errorMsg
}

func (ifs *IpFilterStrategy) Execute(out api_session.IApiSession) *go_error.ErrorInfo {
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
	return go_error.WrapWithAll(fmt.Errorf(ifs.ErrorMsg()), ifs.ErrorCode(), nil)
}
