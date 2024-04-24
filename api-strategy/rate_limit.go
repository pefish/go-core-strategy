package api_strategy

import (
	"context"
	"fmt"
	"time"

	api_session "github.com/pefish/go-core-type/api-session"
	api_strategy "github.com/pefish/go-core-type/api-strategy"
	go_logger "github.com/pefish/go-logger"

	go_error "github.com/pefish/go-error"
)

type RateLimitStrategy struct {
	ctx            context.Context
	logger         go_logger.InterfaceLogger
	errorCode      uint64
	errorMsg       string
	secondPerToken time.Duration
	tokenBucket    chan struct{}
}

func NewRateLimitStrategy(
	ctx context.Context,
	logger go_logger.InterfaceLogger,
	secondPerToken time.Duration,
	maxTokenCount int,
) *RateLimitStrategy {
	return &RateLimitStrategy{
		ctx:            ctx,
		logger:         logger,
		secondPerToken: secondPerToken,
		tokenBucket:    make(chan struct{}, maxTokenCount),
	}
}

func (rls *RateLimitStrategy) Init(param interface{}) api_strategy.IApiStrategy {
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
				timer.Reset(rls.secondPerToken)
			case <-rls.ctx.Done():
				return
			}
		}
	}()
	return rls
}

func (rls *RateLimitStrategy) Name() string {
	return `RateLimitStrategy`
}

func (rls *RateLimitStrategy) Description() string {
	return `rate limit`
}

func (rls *RateLimitStrategy) SetErrorCode(code uint64) api_strategy.IApiStrategy {
	rls.errorCode = code
	return rls
}

func (rls *RateLimitStrategy) ErrorCode() uint64 {
	if rls.errorCode == 0 {
		return go_error.INTERNAL_ERROR_CODE
	}
	return rls.errorCode
}

func (rls *RateLimitStrategy) SetErrorMsg(msg string) api_strategy.IApiStrategy {
	rls.errorMsg = msg
	return rls
}

func (rls *RateLimitStrategy) ErrorMsg() string {
	if rls.errorMsg == "" {
		return "Rate limit reached."
	}
	return rls.errorMsg
}

func (rls *RateLimitStrategy) Execute(out api_session.IApiSession, param interface{}) *go_error.ErrorInfo {
	rls.logger.DebugF(`Api strategy %s trigger.`, rls.Name())
	succ := rls.takeAvailable(out, false)
	if !succ {
		return go_error.WrapWithAll(fmt.Errorf(rls.ErrorMsg()), rls.ErrorCode(), nil)
	}

	return nil
}

func (rls *RateLimitStrategy) takeAvailable(out api_session.IApiSession, block bool) bool {
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
