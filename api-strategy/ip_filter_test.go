package api_strategy

import (
	"github.com/golang/mock/gomock"
	mock_api_session "github.com/pefish/go-core-strategy/api-strategy/mock/mock-api-session"
	_type "github.com/pefish/go-core-type/api-session"
	go_error "github.com/pefish/go-error"
	go_logger "github.com/pefish/go-logger"
	"github.com/pefish/go-test-assert"
	"testing"
)

func TestIpFilterStrategyClass_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	apiSessionInstance := mock_api_session.NewMockIApiSession(ctrl)
	apiSessionInstance.EXPECT().RemoteAddress().Return("127.0.0.34").AnyTimes()
	apiSessionInstance.EXPECT().Logger().Return(go_logger.Logger).AnyTimes()
	err := IpFilterStrategy.Execute(apiSessionInstance, IpFilterParam{GetValidIp: func(apiSession _type.IApiSession) []string {
		return []string{"127.0.0.34"}
	}})
	test.Equal(t, (*go_error.ErrorInfo)(nil), err)

	err = IpFilterStrategy.Execute(apiSessionInstance, IpFilterParam{GetValidIp: func(apiSession _type.IApiSession) []string {
		return []string{"127.0.0.1"}
	}})
	test.Equal(t, IpFilterStrategy.GetErrorCode(), err.Code)
	test.Equal(t, "ip is baned", err.Err.Error())
}
