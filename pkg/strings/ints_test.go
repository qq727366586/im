package strings

import (
	"fmt"
	"testing"
)

func TestInt32sAnd64s(t *testing.T) {
	s := JoinInt32s([]int32{1, 2, 4, 5}, ",")
	fmt.Println(s)
	sli, err := SplitInt32s(s, ",")
	fmt.Println(sli, err)
	fmt.Println("---------------------------------------")
	s64 := JoinInt64s([]int64{1, 2, 4, 5}, ",")
	fmt.Println(s64)
	sli64, err := SplitInt64s(s64, ",")
	fmt.Println(sli64, err)
}
