package api_strategy

import (
	"context"
	"fmt"
	"time"

	i_core "github.com/pefish/go-interface/i-core"
	i_logger "github.com/pefish/go-interface/i-logger"
	t_error "github.com/pefish/go-interface/t-error"
)

type RateLimitStrategy struct {
	ctx         context.Context
	logger      i_logger.ILogger
	errorCode   uint64
	errorMsg    string
	tokenBucket chan struct{}
	params      *RateLimitStrategyParams
}

type RateLimitStrategyParams struct {
	SecondPerToken time.Duration // 每这么长时间往令牌桶塞一个令牌
}

func NewRateLimitStrategy(
	ctx context.Context,
	logger i_logger.ILogger,
	maxTokenCount int,
) *RateLimitStrategy {
	rls := &RateLimitStrategy{
		ctx:         ctx,
		logger:      logger,
		tokenBucket: make(chan struct{}, maxTokenCount),
	}
	return rls
}

func (rls *RateLimitStrategy) Name() string {
	return `RateLimitStrategy`
}

func (rls *RateLimitStrategy) Description() string {
	return `rate limit`
}

func (rls *RateLimitStrategy) SetParamsAndRun(params *RateLimitStrategyParams) *RateLimitStrategy {
	rls.params = params
	go func() {
		timer := time.NewTimer(0)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				select {
				case rls.tokenBucket <- struct{}{}:
					rls.logger.DebugF("[%s] New token to bocket.", rls.Name())
				default:
				}
				timer.Reset(params.SecondPerToken)
			case <-rls.ctx.Done():
				return
			}
		}
	}()
	return rls
}

func (rls *RateLimitStrategy) SetErrorCode(code uint64) i_core.IApiStrategy {
	rls.errorCode = code
	return rls
}

func (rls *RateLimitStrategy) ErrorCode() uint64 {
	if rls.errorCode == 0 {
		return t_error.INTERNAL_ERROR_CODE
	}
	return rls.errorCode
}

func (rls *RateLimitStrategy) SetErrorMsg(msg string) i_core.IApiStrategy {
	rls.errorMsg = msg
	return rls
}

func (rls *RateLimitStrategy) ErrorMsg() string {
	if rls.errorMsg == "" {
		return "Rate limit reached."
	}
	return rls.errorMsg
}

func (rls *RateLimitStrategy) Execute(out i_core.IApiSession) *t_error.ErrorInfo {
	rls.logger.DebugF(`Api strategy %s trigger.`, rls.Name())
	succ := rls.takeAvailable(out, false)
	if !succ {
		return t_error.WrapWithAll(fmt.Errorf(rls.ErrorMsg()), rls.ErrorCode(), nil)
	}

	return nil
}

func (rls *RateLimitStrategy) takeAvailable(out i_core.IApiSession, block bool) bool {
	var takenResult bool
	if block {
		select {
		case <-rls.tokenBucket:
			takenResult = true
		}
	} else {
		select {
		case <-rls.tokenBucket:
			takenResult = true
		default:
			takenResult = false
		}
	}
	rls.logger.DebugF("Current rate limit token count: %d, takenResult: %t", len(rls.tokenBucket), takenResult)
	return takenResult
}
