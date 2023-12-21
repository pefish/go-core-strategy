package api_strategy

import (
	"github.com/golang/mock/gomock"
	mock_api_session "github.com/pefish/go-core-strategy/api-strategy/mock/mock-api-session"
	go_error "github.com/pefish/go-error"
	go_jwt "github.com/pefish/go-jwt"
	go_logger "github.com/pefish/go-logger"
	go_test_ "github.com/pefish/go-test"
	"testing"
	"time"
)

func TestJwtAuthStrategyClass_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	apiSessionInstance := mock_api_session.NewMockIApiSession(ctrl)
	apiSessionInstance.EXPECT().Logger().Return(go_logger.Logger).AnyTimes()
	apiSessionInstance.EXPECT().Header("jwt").Return("gsfdg")
	var userIdResult uint64
	apiSessionInstance.EXPECT().SetJwtBody(gomock.Any()).AnyTimes()
	apiSessionInstance.EXPECT().SetUserId(gomock.Any()).DoAndReturn(func(userId uint64) {
		userIdResult = userId
	}).AnyTimes()
	apiSessionInstance.EXPECT().SetJwtHeaderName(gomock.Any()).AnyTimes()
	JwtAuthApiStrategyInstance.SetHeaderName("jwt")
	err := JwtAuthApiStrategyInstance.Execute(apiSessionInstance, nil)
	go_test_.Equal(t, JwtAuthApiStrategyInstance.GetErrorCode(), err.Code)
	go_test_.Equal(t, "Unauthorized.", err.Err.Error())

	pkey, pubkey, err1 := go_jwt.GeneRsaKeyPair()
	go_test_.Equal(t, nil, err1)
	jwt, err2 := go_jwt.JwtInstance.GetJwt(pkey, 60*time.Second, map[string]interface{}{
		"user_id": 6356,
	})
	go_test_.Equal(t, nil, err2)
	JwtAuthApiStrategyInstance.SetPubKey(pubkey)
	apiSessionInstance.EXPECT().Header("jwt").Return(jwt).AnyTimes()
	apiSessionInstance.EXPECT().Data(gomock.Any()).AnyTimes()
	apiSessionInstance.EXPECT().SetData(gomock.Any(), gomock.Any()).AnyTimes()
	err = JwtAuthApiStrategyInstance.Execute(apiSessionInstance, nil)
	go_test_.Equal(t, (*go_error.ErrorInfo)(nil), err)
	go_test_.Equal(t, uint64(6356), userIdResult)
}
