package test

import (
	"context"
	"errors"
	"fmt"
	"go_util/util_retry"
	"go_util/util_retry/fibonacci"
	"go_util/util_retry/retry_error"
	"testing"
	"time"
)

func TestFibonacci(t *testing.T) {
	ts := time.Now()
	c, _ := context.WithTimeout(context.Background(), 3*time.Second)
	util_retry.Do(c, util_retry.MaxRetry(4, fibonacci.NewFibonacci(500*time.Millisecond)), func(ctx context.Context) error {
		fmt.Println(time.Since(ts))
		return retry_error.RetryableError(errors.New("123"))
	})

}
