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
}

var IpFilterStrategyInstance = NewIpFilterStrategy()

func NewIpFilterStrategy() *IpFilterStrategy {
	return &IpFilterStrategy{}
}

func (ifs *IpFilterStrategy) Init(param interface{}) {

}

type IpFilterParam struct {
	ValidIp func(apiSession api_session.IApiSession) []string
}

func (ifs *IpFilterStrategy) Name() string {
	return `IpFilterStrategy`
}

func (ifs *IpFilterStrategy) Description() string {
	return `filter ip`
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

func (ifs *IpFilterStrategy) Execute(out api_session.IApiSession, param interface{}) *go_error.ErrorInfo {
	out.Logger().DebugF(`api-strategy %s trigger`, ifs.Name())
	if param == nil {
		out.Logger().ErrorF(`strategy need param`)
		return go_error.WrapWithAll(fmt.Errorf(ifs.ErrorMsg()), ifs.ErrorCode(), nil)
	}
	newParam := param.(IpFilterParam)
	if newParam.ValidIp == nil {
		return nil
	}
	clientIp := out.RemoteAddress()
	allowedIps := newParam.ValidIp(out)
	for _, ip := range allowedIps {
		if ip == clientIp {
			return nil
		}
	}
	return go_error.WrapWithAll(fmt.Errorf(ifs.ErrorMsg()), ifs.ErrorCode(), nil)
}
