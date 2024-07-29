package api_strategy

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_api_session "github.com/pefish/go-core-strategy/api-strategy/mock/mock-api-session"
	i_logger "github.com/pefish/go-interface/i-logger"
	go_test_ "github.com/pefish/go-test"
)

func TestIpFilterStrategyClass_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	apiSessionInstance := mock_api_session.NewMockIApiSession(ctrl)
	apiSessionInstance.EXPECT().RemoteAddress().Return("127.0.0.34").AnyTimes()
	apiSessionInstance.EXPECT().Logger().Return(i_logger.DefaultLogger).AnyTimes()
	ipFilterStrategyInstance := NewIpFilterStrategy()
	ipFilterStrategyInstance.Init(nil)
	err := ipFilterStrategyInstance.Execute(
		apiSessionInstance,
		IpFilterParams{
			ValidIp: func(apiSession _type.IApiSession) []string {
				return []string{"127.0.0.34"}
			},
		},
	)
	go_test_.Equal(t, (*go_error.ErrorInfo)(nil), err)

	err = ipFilterStrategyInstance.Execute(apiSessionInstance, IpFilterParams{ValidIp: func(apiSession _type.IApiSession) []string {
		return []string{"127.0.0.1"}
	}})
	go_test_.Equal(t, ipFilterStrategyInstance.ErrorCode(), err.Code)
	go_test_.Equal(t, "Ip is baned.", err.Err.Error())
}
