package string_tool

import (
	"fmt"
	"testing"
)

func TestFormatMac(t *testing.T) {
	var (
		mac1 = "00:12:6F:5F:32:B7"
		mac2 = "00126F5F32B7"
		mac3 = "00:12:6F:5F:32:B"
		mac4 = "00:12:6F:5F-32-b7"
		mac5 = "00:12:6F:5F:32:B7:00"
		mac6 = "00:12:6F:5F:32:B7:0"
	)

	if mac, err := FormatMac(mac1); err != nil || mac != mac1 {
		t.Error(err, mac)
	}

	if mac, err := FormatMac(mac2); err != nil || mac != mac1 {
		t.Error(err, mac)
	}

	if mac, err := FormatMac(mac3); err == nil {
		t.Errorf("应该无法生成mac地址，但是确得到%v", mac)
	}

	if mac, err := FormatMac(mac4); err != nil || mac != mac1 {
		t.Error(err, mac)
	}

	if mac, err := FormatMac(mac5); err == nil {
		t.Errorf("应该无法生成mac地址，但是确得到%v", mac)
	}

	if mac, err := FormatMac(mac6); err == nil {
		t.Errorf("应该无法生成mac地址，但是确得到%v", mac)
	}
}

func TestRandom(t *testing.T) {

	for i := 0; i < 1000; i++ {
		fmt.Println(GetRandomString(10))
		fmt.Println(GetRandomNumber(20))
		fmt.Println(GetRandomChar(30))
	}

}
