package api_strategy

import (
	"fmt"
	jwt2 "github.com/dgrijalva/jwt-go"
	api_session "github.com/pefish/go-core-type/api-session"
	api_strategy "github.com/pefish/go-core-type/api-strategy"
	"github.com/pefish/go-error"
	"github.com/pefish/go-jwt"
	"github.com/pefish/go-reflect"
)

type JwtAuthStrategy struct {
	errorCode     uint64
	pubKey        string
	headerName    string
	noCheckExpire bool
	disableUserId bool
	errorMsg      string
}

var JwtAuthApiStrategyInstance = NewJwtAuthStrategy()

func NewJwtAuthStrategy() *JwtAuthStrategy {
	return &JwtAuthStrategy{}
}

func (jas *JwtAuthStrategy) GetName() string {
	return `JwtAuthStrategy`
}

func (jas *JwtAuthStrategy) GetDescription() string {
	return `jwt auth`
}

func (jas *JwtAuthStrategy) SetErrorCode(code uint64) api_strategy.IApiStrategy {
	jas.errorCode = code
	return jas
}

func (jas *JwtAuthStrategy) SetErrorMsg(msg string) api_strategy.IApiStrategy {
	jas.errorMsg = msg
	return jas
}

func (jas *JwtAuthStrategy) GetErrorMsg() string {
	if jas.errorMsg == "" {
		return "Unauthorized."
	}
	return jas.errorMsg
}

func (jas *JwtAuthStrategy) GetErrorCode() uint64 {
	if jas.errorCode == 0 {
		return go_error.INTERNAL_ERROR_CODE
	}
	return jas.errorCode
}

func (jas *JwtAuthStrategy) SetNoCheckExpire() {
	jas.noCheckExpire = true
}

func (jas *JwtAuthStrategy) DisableUserId() {
	jas.disableUserId = true
}

func (jas *JwtAuthStrategy) SetPubKey(pubKey string) {
	jas.pubKey = pubKey
}

func (jas *JwtAuthStrategy) SetHeaderName(headerName string) {
	jas.headerName = headerName
}

func (jas *JwtAuthStrategy) Execute(out api_session.IApiSession, param interface{}) *go_error.ErrorInfo {
	out.Logger().DebugF(`api-strategy %s trigger`, jas.GetName())

	headerName := jas.headerName
	if headerName == "" {
		headerName = "Json-Web-Token"
	}
	out.SetJwtHeaderName(headerName)
	jwt := out.Header(headerName)

	verifyResult, token, err := go_jwt.Jwt.VerifyJwt(jas.pubKey, jwt, jas.noCheckExpire)
	if err != nil {
		return go_error.WrapWithAll(fmt.Errorf(jas.GetErrorMsg()), jas.GetErrorCode(), nil)
	}
	if !verifyResult {
		return go_error.WrapWithAll(fmt.Errorf(jas.GetErrorMsg()), jas.GetErrorCode(), nil)
	}
	jwtBody := token.Claims.(jwt2.MapClaims)
	out.SetJwtBody(jwtBody)
	if !jas.disableUserId {
		jwtPayload := jwtBody[`payload`].(map[string]interface{})
		if jwtPayload[`user_id`] == nil {
			out.Logger().ErrorF(`jwt verify error, user_id not exist`)
			return go_error.WrapWithAll(fmt.Errorf(jas.GetErrorMsg()), jas.GetErrorCode(), nil)
		}

		userId := go_reflect.Reflect.MustToUint64(jwtPayload[`user_id`])
		out.SetUserId(userId)

		errorMsg := out.Data(`error_msg`)
		if errorMsg == nil {
			out.SetData(`error_msg`, fmt.Sprintf("%s: %v\n", `jas`, userId))
		} else {
			out.SetData(`error_msg`, fmt.Sprintf("%s%s: %v\n", errorMsg.(string), `jas`, userId))
		}
	}
	return nil
}
