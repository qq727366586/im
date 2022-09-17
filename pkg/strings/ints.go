package strings

import (
	"bytes"
	"im/constant/str"
	"strconv"
	"strings"
	"sync"
)

var (
	bufferPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer([]byte{})
		},
	}
)

// 插入字符 例如: 1,2,3,4
func JoinInt32s(is []int32, s string) string {
	if len(is) == 0 {
		return str.StringIsEmpty
	}
	if len(is) == 1 {
		return strconv.FormatInt(int64(is[0]), 10)
	}
	buffer := bufferPool.Get().(*bytes.Buffer)
	for _, v := range is {
		buffer.WriteString(strconv.FormatInt(int64(v), 10))
		buffer.WriteString(s)
	}
	if buffer.Len() > 0 {
		buffer.Truncate(buffer.Len() - 1)
	}
	out := buffer.String()
	buffer.Reset()
	bufferPool.Put(buffer)
	return out
}

func SplitInt32s(s, p string) ([]int32, error) {
	if s == "" {
		return nil, nil
	}
	strSlice := strings.Split(s, p)
	res := make([]int32, 0, len(strSlice))
	for _, v := range strSlice {
		int32s, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return nil, err
		}
		res = append(res, int32(int32s))
	}
	return res, nil
}

// 64位
func JoinInt64s(is []int64, s string) string {
	if len(is) == 0 {
		return str.StringIsEmpty
	}
	if len(is) == 1 {
		return strconv.FormatInt(is[0], 10)
	}
	buffer := bufferPool.Get().(*bytes.Buffer)
	for _, v := range is {
		buffer.WriteString(strconv.FormatInt(v, 10))
		buffer.WriteString(s)
	}
	if buffer.Len() > 0 {
		buffer.Truncate(buffer.Len() - 1)
	}
	out := buffer.String()
	buffer.Reset()
	bufferPool.Put(buffer)
	return out
}

func SplitInt64s(s, p string) ([]int64, error) {
	if s == str.StringIsEmpty {
		return nil, nil
	}
	strSlice := strings.Split(s, p)
	res := make([]int64, 0, len(strSlice))
	for _, v := range strSlice {
		int64s, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, int64s)
	}
	return res, nil
}
