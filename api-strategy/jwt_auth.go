package api_strategy

import (
	"fmt"

	i_core "github.com/pefish/go-interface/i-core"
	t_error "github.com/pefish/go-interface/t-error"
	go_jwt "github.com/pefish/go-jwt"
	go_reflect "github.com/pefish/go-reflect"
)

type JwtAuthStrategy struct {
	errorCode uint64
	errorMsg  string
	params    *JwtAuthStrategyParams
}

type JwtAuthStrategyParams struct {
	PubKey        string
	HeaderName    string
	NoCheckExpire bool
	DisableUserId bool
}

func NewJwtAuthStrategy() *JwtAuthStrategy {
	return &JwtAuthStrategy{}
}

func (jas *JwtAuthStrategy) Name() string {
	return `JwtAuthStrategy`
}

func (jas *JwtAuthStrategy) Description() string {
	return `jwt auth`
}

func (jas *JwtAuthStrategy) SetErrorCode(code uint64) i_core.IApiStrategy {
	jas.errorCode = code
	return jas
}

func (jas *JwtAuthStrategy) SetErrorMsg(msg string) i_core.IApiStrategy {
	jas.errorMsg = msg
	return jas
}

func (jas *JwtAuthStrategy) ErrorMsg() string {
	if jas.errorMsg == "" {
		return "Unauthorized."
	}
	return jas.errorMsg
}

func (jas *JwtAuthStrategy) ErrorCode() uint64 {
	if jas.errorCode == 0 {
		return t_error.INTERNAL_ERROR_CODE
	}
	return jas.errorCode
}

func (jas *JwtAuthStrategy) SetParams(params *JwtAuthStrategyParams) *JwtAuthStrategy {
	jas.params = params
	return jas
}

func (jas *JwtAuthStrategy) Execute(out i_core.IApiSession) *t_error.ErrorInfo {
	out.Logger().DebugF(`Api strategy %s trigger`, jas.Name())

	headerName := jas.params.HeaderName
	if headerName == "" {
		headerName = "Json-Web-Token"
	}
	out.SetJwtHeaderName(headerName)
	jwt := out.Header(headerName)

	verifyResult, _, body, err := go_jwt.JwtInstance.VerifyJwt(jas.params.PubKey, jwt, jas.params.NoCheckExpire)
	if err != nil {
		out.Logger().ErrorF(`VerifyJwt error - %+v.`, err)
		return t_error.WrapWithAll(fmt.Errorf(jas.ErrorMsg()), jas.ErrorCode(), nil)
	}
	if !verifyResult {
		out.Logger().ErrorF(`VerifyJwt error - verify result is false.`)
		return t_error.WrapWithAll(fmt.Errorf(jas.ErrorMsg()), jas.ErrorCode(), nil)
	}
	out.SetJwtBody(body)
	if !jas.params.DisableUserId {
		jwtPayload := body[`payload`].(map[string]interface{})
		if jwtPayload[`user_id`] == nil {
			out.Logger().ErrorF(`Jwt verify error, user_id not exist.`)
			return t_error.WrapWithAll(fmt.Errorf(jas.ErrorMsg()), jas.ErrorCode(), nil)
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
