package bytes

import (
	"fmt"
	"testing"
)

func TestBytePool(t *testing.T) {
	bp := NewBytePool(10, 200)
	for i := 0; i < 10; i++ {
		b := bp.Get()
		fmt.Println(b.buf, b.next, b.Bytes())
		bp.Put(b)
	}
	for i := 0; i < 10; i++ {
		b := bp.Get()
		fmt.Println(&b.next)
	}
}
