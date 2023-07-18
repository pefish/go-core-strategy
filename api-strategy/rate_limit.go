package api_strategy

import (
	"fmt"
	go_application "github.com/pefish/go-application"
	_type "github.com/pefish/go-core-type/api-session"
	"time"

	"github.com/pefish/go-error"
)

type RateLimitStrategy struct {
	errorCode      uint64
	errorMsg       string
	secondPerToken time.Duration
	tokenBucket    chan struct{}
}

var RateLimitStrategyInstance = NewRateLimitStrategy(time.Second, 5)

func NewRateLimitStrategy(secondPerToken time.Duration, maxTokenCount int) *RateLimitStrategy {
	rateLimitStrategyInstance := &RateLimitStrategy{
		secondPerToken: secondPerToken,
		tokenBucket:    make(chan struct{}, maxTokenCount),
	}
	go func(rateLimitStrategyInstance *RateLimitStrategy) {
		timer := time.NewTimer(0)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				select {
				case rateLimitStrategyInstance.tokenBucket <- struct{}{}:
				default:
				}
				timer.Reset(rateLimitStrategyInstance.secondPerToken)
			case <-go_application.Application.OnFinished():
				return
			}
		}
	}(rateLimitStrategyInstance)
	return rateLimitStrategyInstance
}

func (rls *RateLimitStrategy) GetName() string {
	return `RateLimitStrategy`
}

func (rls *RateLimitStrategy) GetDescription() string {
	return `rate limit`
}

func (rls *RateLimitStrategy) SetErrorCode(code uint64) IStrategy {
	rls.errorCode = code
	return rls
}

func (rls *RateLimitStrategy) GetErrorCode() uint64 {
	if rls.errorCode == 0 {
		return go_error.INTERNAL_ERROR_CODE
	}
	return rls.errorCode
}

func (rls *RateLimitStrategy) SetErrorMsg(msg string) IStrategy {
	rls.errorMsg = msg
	return rls
}

func (rls *RateLimitStrategy) GetErrorMsg() string {
	if rls.errorMsg == "" {
		return "Rate limit reached."
	}
	return rls.errorMsg
}

func (rls *RateLimitStrategy) Execute(out _type.IApiSession, param interface{}) *go_error.ErrorInfo {
	out.Logger().DebugF(`api-strategy %s trigger`, rls.GetName())
	out.Logger().Error(rls)
	succ := rls.takeAvailable(out, false)
	if !succ {
		return go_error.WrapWithAll(fmt.Errorf(rls.GetErrorMsg()), rls.GetErrorCode(), nil)
	}

	return nil
}

func (rls *RateLimitStrategy) takeAvailable(out _type.IApiSession, block bool) bool {
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
	out.Logger().DebugF("current global rate limit token count: %d", len(rls.tokenBucket))
	return takenResult
}
