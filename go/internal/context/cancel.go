package context

import (
	"context"
	"fmt"
	"time"
)

func useCancelCtx(parent context.Context) {
	ctx, cancelFunc := context.WithCancel(parent)

	go func() {
		fmt.Println("another goroutine start")
		time.Sleep(time.Second * 2)
		cancelFunc()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("ctx has been canceled")
	}

	return
}
