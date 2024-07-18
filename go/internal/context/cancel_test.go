package context

import (
	"context"
	"testing"
)

func Test_cancelFunc(t *testing.T) {
	ctx := context.Background()
	useCancelCtx(ctx)
}

func Test_timeoutFunc(t *testing.T) {
	ctx := context.Background()
	useTimeoutCtx(ctx)

}
