package ip

import (
	"testing"

	"go-demo/utils/ip/address"
)

func Test_Address(t *testing.T) {
	ip := GetInternetIP()
	address, _ := address.GetAddressByIP(ip)
	t.Logf("ip:%s, %+v", ip, address)
	localIp := GetLocalIP()
	t.Logf("ip:%s", localIp)
}
