package fibonacci

import (
	"math"
	"sync/atomic"
	"time"
	"unsafe"
)

type state [2]time.Duration

type fibonacciBackoff struct {
	state unsafe.Pointer
}

func NewFibonacci(base time.Duration) *fibonacciBackoff {
	if base <= 0 {
		panic("base must be greater than 0")
	}
	return &fibonacciBackoff{
		state: unsafe.Pointer(&state{0, base}),
	}
}

func (f *fibonacciBackoff) Next() (time.Duration, bool) {
	for {
		// get addr
		cur := atomic.LoadPointer(&f.state)
		curState := (*state)(cur)

		next := curState[0] + curState[1]

		if next <= 0 {
			return math.MaxInt64, false
		}
		if atomic.CompareAndSwapPointer(&f.state, cur, unsafe.Pointer(&state{curState[1], next})) {
			return next, false
		}
	}
}
