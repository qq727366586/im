package conf

import (
	"fmt"
	"testing"
)

func TestBaseConfig(t *testing.T) {
	Init()
	fmt.Println(Conf.TCP)
}
