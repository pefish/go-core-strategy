// 全局限流器（令牌桶）
package global_api_strategy

import (
	"context"
	"errors"
	"time"

	_type "github.com/pefish/go-core-type/api-session"
	api_strategy "github.com/pefish/go-core-type/api-strategy"
	go_error "github.com/pefish/go-error"
)

type GlobalRateLimitStrategy struct {
	ctx         context.Context
	tokenBucket chan struct{}
	errorCode   uint64
	errorMsg    string
}

var GlobalRateLimitStrategyInstance = NewGlobalRateLimitStrategy(context.Background())

func NewGlobalRateLimitStrategy(ctx context.Context) *GlobalRateLimitStrategy {
	return &GlobalRateLimitStrategy{
		ctx:         ctx,
		tokenBucket: make(chan struct{}, 200),
	}
}

func (grls *GlobalRateLimitStrategy) Name() string {
	return `GlobalRateLimitStrategy`
}

func (grls *GlobalRateLimitStrategy) Description() string {
	return `global rate limit for all api`
}

func (grls *GlobalRateLimitStrategy) SetErrorCode(code uint64) api_strategy.IApiStrategy {
	grls.errorCode = code
	return grls
}

func (grls *GlobalRateLimitStrategy) SetErrorMsg(msg string) api_strategy.IApiStrategy {
	grls.errorMsg = msg
	return grls
}

func (grls *GlobalRateLimitStrategy) ErrorMsg() string {
	if grls.errorMsg == "" {
		return "Global rate limit."
	}
	return grls.errorMsg
}

func (grls *GlobalRateLimitStrategy) ErrorCode() uint64 {
	if grls.errorCode == 0 {
		return go_error.INTERNAL_ERROR_CODE
	}
	return grls.errorCode
}

func (grls *GlobalRateLimitStrategy) Init(param interface{}) api_strategy.IApiStrategy {
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
			case <-grls.ctx.Done():
				return
			}
		}
	}()
	return grls
}

type GlobalRateLimitStrategyParam struct {
	FillInterval time.Duration // 每这么长时间往令牌桶塞一个令牌
}

func (grls *GlobalRateLimitStrategy) Execute(out _type.IApiSession, param interface{}) *go_error.ErrorInfo {
	out.Logger().DebugF(`api-strategy %s trigger`, grls.Name())

	succ := grls.takeAvailable(out, false)
	if !succ {
		return go_error.WrapWithAll(errors.New(grls.ErrorMsg()), grls.ErrorCode(), nil)
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
