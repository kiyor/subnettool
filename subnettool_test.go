/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : subnettool_test.go

* Purpose :

* Creation Date : 02-27-2014

* Last Modified : Sun 14 May 2017 05:04:59 AM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package subnettool

import (
	// 	"fmt"
	"net"
	"strconv"
	"testing"
)

func Test_Aton(t *testing.T) {
	if Aton(net.ParseIP("8.8.8.8")) != 134744072 {
		t.Error("Aton did not work as expected.")
	} else {
		t.Log("Aton passed.")
	}
}

func Test_Ntoa(t *testing.T) {
	if Ntoa(134744072).String() != "8.8.8.8" {
		t.Error("Ntoa did not work as expected.")
	} else {
		t.Log("Ntoa passed.")
	}
}

func Test_Mtos(t *testing.T) {
	if Mtos(net.ParseIP("255.255.255.0")) != 24 {
		t.Error("Mtos did not work as expected.")
		t.Error("255.255.255.0 >", net.ParseIP("255.255.255.0"))
	} else {
		t.Log("Mtos passed.")
	}
}

func Test_Stom(t *testing.T) {
	if Stom(24).String() != "255.255.255.0" {
		t.Error("Stom did not work as expected.")
		t.Error("24 >", Stom(24).String())
	} else {
		t.Log("Stom passed.")
	}
}

func Test_GetAllIP(t *testing.T) {
	var ips []net.IP
	for i := 0; i < 4; i++ {
		ipstr := "192.168.1." + strconv.Itoa(i)
		ips = append(ips, net.ParseIP(ipstr))
	}
	if match(GetAllIP("192.168.1.1/30"), ips) {
		t.Log("GetAllIP passed.")
	} else {
		t.Error(GetAllIP("192.168.1.1/30"))
	}
}

func Test_Nth(t *testing.T) {
	if Nth("192.168.1.0/30", 1).String() == net.ParseIP("192.168.1.0").String() {
		t.Log("GetAllIP passed.")
	} else {
		t.Error(Nth("192.168.1.0/30", 1).String())
	}
}

func Test_Base(t *testing.T) {
	if Base("192.168.1.100/24").String() == net.ParseIP("192.168.1.0").String() {
		t.Log("Base passed.")
	} else {
		t.Error(Base("192.168.1.100/24").String())
	}
}

func Test_Len(t *testing.T) {
	if Len("192.168.1.1/24") == 256 {
		t.Log("Len passed.")
	} else {
		t.Error(Len("192.168.1.1/24"))
	}
}

func Test_GetMask(t *testing.T) {
	if GetMask("192.168.1.1/24") == 24 {
		t.Log("GetMask passed.")
	} else {
		t.Error(GetMask("192.168.1.1/24"))
	}
}

func Test_ParseIPInt(t *testing.T) {
	res := ParseIPInt(net.ParseIP("192.168.1.100"))
	if res[0] == 192 && res[1] == 168 && res[2] == 1 && res[3] == 100 {
		t.Log("ParseIPInt passed.")
	} else {
		t.Error(ParseIPInt(net.ParseIP("192.168.1.100")))
	}
}

func Test_TestCIDRMatch(t *testing.T) {
	type item struct {
		addr    string
		cidr    string
		matches bool
	}
	var testdata = []item{
		item{"192.168.1.67", "192.168.1.0/24", true},
		item{"192.168.1.67", "192.168.1.0/28", false},
		item{"192.168.1.67", "0.0.0.0/0", true},
	}
	for _, it := range testdata {
		_, cidrnet, err := net.ParseCIDR(it.cidr)
		if err != nil {
			panic(err) // assuming I did it right above
		}
		myaddr := net.ParseIP(it.addr)
		if cidrnet.Contains(myaddr) != it.matches {
			t.Fatalf("Wrong on %+v")
		}
	}

}

func match(a, b []net.IP) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.String() != b[i].String() {
			return false
		}
	}
	return true
}
