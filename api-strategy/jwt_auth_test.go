package api_strategy

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mock_api_session "github.com/pefish/go-core-strategy/api-strategy/mock/mock-api-session"
	i_logger "github.com/pefish/go-interface/i-logger"
	t_error "github.com/pefish/go-interface/t-error"
	go_jwt "github.com/pefish/go-jwt"
	go_test_ "github.com/pefish/go-test"
)

func TestJwtAuthStrategyClass_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	apiSessionInstance := mock_api_session.NewMockIApiSession(ctrl)
	apiSessionInstance.EXPECT().Logger().Return(i_logger.DefaultLogger).AnyTimes()
	apiSessionInstance.EXPECT().Header("jwt").Return("gsfdg")
	var userIdResult uint64
	apiSessionInstance.EXPECT().SetJwtBody(gomock.Any()).AnyTimes()
	apiSessionInstance.EXPECT().SetUserId(gomock.Any()).DoAndReturn(func(userId uint64) {
		userIdResult = userId
	}).AnyTimes()
	apiSessionInstance.EXPECT().SetJwtHeaderName(gomock.Any()).AnyTimes()
	jwtAuthApiStrategyInstance := NewJwtAuthStrategy()
	jwtAuthApiStrategyInstance.Init(nil)
	jwtAuthApiStrategyInstance.SetHeaderName("jwt")
	err := jwtAuthApiStrategyInstance.Execute(apiSessionInstance, nil)
	go_test_.Equal(t, jwtAuthApiStrategyInstance.ErrorCode(), err.Code)
	go_test_.Equal(t, "Unauthorized.", err.Err.Error())

	pkey, pubkey, err1 := go_jwt.GeneRsaKeyPair()
	go_test_.Equal(t, nil, err1)
	jwt, err2 := go_jwt.JwtInstance.GetJwt(pkey, 60*time.Second, map[string]interface{}{
		"user_id": 6356,
	})
	go_test_.Equal(t, nil, err2)
	jwtAuthApiStrategyInstance.SetPubKey(pubkey)
	apiSessionInstance.EXPECT().Header("jwt").Return(jwt).AnyTimes()
	apiSessionInstance.EXPECT().Data(gomock.Any()).AnyTimes()
	apiSessionInstance.EXPECT().SetData(gomock.Any(), gomock.Any()).AnyTimes()
	err = jwtAuthApiStrategyInstance.Execute(apiSessionInstance, nil)
	go_test_.Equal(t, (*t_error.ErrorInfo)(nil), err)
	go_test_.Equal(t, uint64(6356), userIdResult)
}
