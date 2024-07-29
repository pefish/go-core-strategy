// 全局限流器（令牌桶）
package global_api_strategy

import (
	"context"
	"errors"
	"time"

	i_core "github.com/pefish/go-interface/i-core"
	t_error "github.com/pefish/go-interface/t-error"
)

type GlobalRateLimitStrategy struct {
	ctx         context.Context
	tokenBucket chan struct{}
	errorCode   uint64
	errorMsg    string
	params      *GlobalRateLimitStrategyParams
}

var GlobalRateLimitStrategyInstance = NewGlobalRateLimitStrategy(context.Background())

func NewGlobalRateLimitStrategy(
	ctx context.Context,
) *GlobalRateLimitStrategy {
	grls := &GlobalRateLimitStrategy{
		ctx:         ctx,
		tokenBucket: make(chan struct{}, 200),
	}
	return grls
}

func (grls *GlobalRateLimitStrategy) Name() string {
	return `GlobalRateLimitStrategy`
}

func (grls *GlobalRateLimitStrategy) Description() string {
	return `global rate limit for all api`
}

func (grls *GlobalRateLimitStrategy) SetParamsAndRun(params *GlobalRateLimitStrategyParams) i_core.IApiStrategy {
	grls.params = params
	go func() {
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

func (grls *GlobalRateLimitStrategy) SetErrorCode(code uint64) i_core.IApiStrategy {
	grls.errorCode = code
	return grls
}

func (grls *GlobalRateLimitStrategy) SetErrorMsg(msg string) i_core.IApiStrategy {
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
		return t_error.INTERNAL_ERROR_CODE
	}
	return grls.errorCode
}

type GlobalRateLimitStrategyParams struct {
	FillInterval time.Duration // 每这么长时间往令牌桶塞一个令牌
}

func (grls *GlobalRateLimitStrategy) Execute(out i_core.IApiSession) *t_error.ErrorInfo {
	out.Logger().DebugF(`api-strategy %s trigger`, grls.Name())

	succ := grls.takeAvailable(out, false)
	if !succ {
		return t_error.WrapWithAll(errors.New(grls.ErrorMsg()), grls.ErrorCode(), nil)
	}

	return nil
}

func (grls *GlobalRateLimitStrategy) takeAvailable(out i_core.IApiSession, block bool) bool {
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
