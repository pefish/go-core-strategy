// 全局限流器（令牌桶）
package global_api_strategy

import (
	"errors"
	go_application "github.com/pefish/go-application"
	_type "github.com/pefish/go-core-type/api-session"
	global_api_strategy "github.com/pefish/go-core-type/global-api-strategy"
	"github.com/pefish/go-error"
	"time"
)

type GlobalRateLimitStrategy struct {
	tokenBucket chan struct{}
	errorCode   uint64
	errorMsg    string
}

var GlobalRateLimitStrategyInstance = NewGlobalRateLimitStrategy()

func NewGlobalRateLimitStrategy() *GlobalRateLimitStrategy {
	return &GlobalRateLimitStrategy{
		tokenBucket: make(chan struct{}, 200),
	}
}

func (grls *GlobalRateLimitStrategy) GetName() string {
	return `GlobalRateLimitStrategy`
}

func (grls *GlobalRateLimitStrategy) GetDescription() string {
	return `global rate limit for all api`
}

func (grls *GlobalRateLimitStrategy) SetErrorCode(code uint64) global_api_strategy.IGlobalApiStrategy {
	grls.errorCode = code
	return grls
}

func (grls *GlobalRateLimitStrategy) SetErrorMsg(msg string) global_api_strategy.IGlobalApiStrategy {
	grls.errorMsg = msg
	return grls
}

func (grls *GlobalRateLimitStrategy) GetErrorMsg() string {
	if grls.errorMsg == "" {
		return "Global rate limit."
	}
	return grls.errorMsg
}

func (grls *GlobalRateLimitStrategy) GetErrorCode() uint64 {
	if grls.errorCode == 0 {
		return go_error.INTERNAL_ERROR_CODE
	}
	return grls.errorCode
}

func (grls *GlobalRateLimitStrategy) Init(param interface{}) {
	go func() {
		params := param.(GlobalRateLimitStrategyParam)
		ticker := time.NewTicker(params.FillInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				select {
				case grls.tokenBucket <- struct{}{}:
				default:
				}
			case <-go_application.Application.OnFinished():
				return
			}
		}
	}()
}

type GlobalRateLimitStrategyParam struct {
	FillInterval time.Duration // 每这么长时间往令牌桶塞一个令牌
}

func (grls *GlobalRateLimitStrategy) Execute(out _type.IApiSession, param interface{}) *go_error.ErrorInfo {
	out.Logger().DebugF(`api-strategy %s trigger`, grls.GetName())

	succ := grls.takeAvailable(out, false)
	if !succ {
		return go_error.WrapWithAll(errors.New(grls.GetErrorMsg()), grls.GetErrorCode(), nil)
	}

	return nil
}

func (grls *GlobalRateLimitStrategy) takeAvailable(out _type.IApiSession, block bool) bool {
	var takenResult bool
	if block {
		select {
		case <-grls.tokenBucket:
			takenResult = true
		}
	} else {
		select {
		case <-grls.tokenBucket:
			takenResult = true
		default:
			takenResult = false
		}
	}
	out.Logger().DebugF("current global rate limit token count: %d", len(grls.tokenBucket))
	return takenResult
}
