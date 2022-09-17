package ip

import (
	"fmt"
	"im/constant/str"
	"testing"
)

func TestIp(t *testing.T) {
	if ip := InternalIP(); ip != str.StringIsEmpty {
		fmt.Println(ip)
	} else {
		t.FailNow()
	}
}
