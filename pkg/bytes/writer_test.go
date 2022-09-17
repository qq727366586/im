package bytes

import (
	"fmt"
	"im/pkg/bufio"
	"testing"
)

func TestWriter(t *testing.T) {
	w := bufio.NewWriterSize(10)
	w.Write(make([]byte, 10000, 10000))
	fmt.Println(w.Len(), w.Size(), len(w.Buffer()))
}
