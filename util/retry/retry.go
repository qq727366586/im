package retry

import (
	"context"
	"errors"
	"go_util/util_retry/retry_error"
	"sync"
	"time"
)

type BackOff interface {
	Next() (next time.Duration, stop bool)
}

type BackOffFunc func() (time.Duration, bool)

func (b BackOffFunc) Next() (time.Duration, bool) {
	return b()
}

func MaxRetry(maxTime uint64, next BackOff) BackOff {
	var mx sync.Mutex
	var attempt uint64

	return BackOffFunc(func() (time.Duration, bool) {
		mx.Lock()
		defer mx.Unlock()
		if attempt >= maxTime {
			return 0, true
		}
		attempt++

		val, stop := next.Next()
		if stop {
			return 0, true
		}
		return val, false
	})
}

func Do(ctx context.Context, b BackOff, f retry_error.RetryFunc) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		err := f(ctx)
		if err == nil {
			return nil
		}
		// Not retryable
		var rerr *retry_error.RetryError
		if !errors.As(err, &rerr) {
			return err
		}

		next, stop := b.Next()
		if stop {
			return rerr.Unwrap()
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		t := time.NewTimer(next)
		select {
		case <-ctx.Done():
			t.Stop()
			return ctx.Err()
		case <-t.C:
			continue
		}
	}
}
