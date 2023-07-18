package api_strategy

import (
	"fmt"
	_type "github.com/pefish/go-core-type/api-session"

	"github.com/pefish/go-error"
)

type IpFilterStrategy struct {
	errorCode uint64
	errorMsg  string
}

var IpFilterStrategyInstance = NewIpFilterStrategy()

func NewIpFilterStrategy() *IpFilterStrategy {
	return &IpFilterStrategy{}
}

type IpFilterParam struct {
	GetValidIp func(apiSession _type.IApiSession) []string
}

func (ifs *IpFilterStrategy) GetName() string {
	return `IpFilterStrategy`
}

func (ifs *IpFilterStrategy) GetDescription() string {
	return `filter ip`
}

func (ifs *IpFilterStrategy) SetErrorCode(code uint64) IStrategy {
	ifs.errorCode = code
	return ifs
}

func (ifs *IpFilterStrategy) GetErrorCode() uint64 {
	if ifs.errorCode == 0 {
		return go_error.INTERNAL_ERROR_CODE
	}
	return ifs.errorCode
}

func (ifs *IpFilterStrategy) SetErrorMsg(msg string) IStrategy {
	ifs.errorMsg = msg
	return ifs
}

func (ifs *IpFilterStrategy) GetErrorMsg() string {
	if ifs.errorMsg == "" {
		return "Ip is baned."
	}
	return ifs.errorMsg
}

func (ifs *IpFilterStrategy) Execute(out _type.IApiSession, param interface{}) *go_error.ErrorInfo {
	out.Logger().DebugF(`api-strategy %s trigger`, ifs.GetName())
	if param == nil {
		out.Logger().ErrorF(`strategy need param`)
		return go_error.WrapWithAll(fmt.Errorf(ifs.GetErrorMsg()), ifs.GetErrorCode(), nil)
	}
	newParam := param.(IpFilterParam)
	if newParam.GetValidIp == nil {
		return nil
	}
	clientIp := out.RemoteAddress()
	allowedIps := newParam.GetValidIp(out)
	for _, ip := range allowedIps {
		if ip == clientIp {
			return nil
		}
	}
	return go_error.WrapWithAll(fmt.Errorf(ifs.GetErrorMsg()), ifs.GetErrorCode(), nil)
}
