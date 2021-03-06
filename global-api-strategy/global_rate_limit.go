// 全局限流器（令牌桶）
package global_api_strategy

import (
	"errors"
	go_application "github.com/pefish/go-application"
	_type "github.com/pefish/go-core-type/api-session"
	"github.com/pefish/go-error"
	"time"
)

type GlobalRateLimitStrategyClass struct {
	tokenBucket chan struct{}
	errorCode uint64
}

var GlobalRateLimitStrategy = GlobalRateLimitStrategyClass{
	tokenBucket: make(chan struct{}, 200),
}

func (globalRateLimit *GlobalRateLimitStrategyClass) GetName() string {
	return `GlobalRateLimit`
}

func (globalRateLimit *GlobalRateLimitStrategyClass) GetDescription() string {
	return `global rate limit for all api`
}

func (globalRateLimit *GlobalRateLimitStrategyClass) SetErrorCode(code uint64) {
	globalRateLimit.errorCode = code
}

func (globalRateLimit *GlobalRateLimitStrategyClass) GetErrorCode() uint64 {
	if globalRateLimit.errorCode == 0 {
		return go_error.INTERNAL_ERROR_CODE
	}
	return globalRateLimit.errorCode
}


func (globalRateLimit *GlobalRateLimitStrategyClass) Init(param interface{}) {
	go func() {
		params := param.(GlobalRateLimitStrategyParam)
		ticker := time.NewTicker(params.FillInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				select {
				case globalRateLimit.tokenBucket <- struct{}{}:
				default:
				}
			case <- go_application.Application.OnFinished():
				return
			}
		}
	}()
}

type GlobalRateLimitStrategyParam struct {
	FillInterval time.Duration  // 每这么长时间往令牌桶塞一个令牌
}

func (globalRateLimit *GlobalRateLimitStrategyClass) Execute(out _type.IApiSession, param interface{}) *go_error.ErrorInfo {
	out.Logger().DebugF(`api-strategy %s trigger`, globalRateLimit.GetName())

	succ := globalRateLimit.takeAvailable(out, false)
	if !succ {
		return go_error.WrapWithAll(errors.New(`global rate limit`), globalRateLimit.errorCode, nil)
	}

	return nil
}

func (globalRateLimit *GlobalRateLimitStrategyClass) takeAvailable(out _type.IApiSession, block bool) bool{
	var takenResult bool
	if block {
		select {
		case <-globalRateLimit.tokenBucket:
			takenResult = true
		}
	} else {
		select {
		case <-globalRateLimit.tokenBucket:
			takenResult = true
		default:
			takenResult = false
		}
	}
	out.Logger().DebugF("current global rate limit token count: %d", len(globalRateLimit.tokenBucket))
	return takenResult
}
