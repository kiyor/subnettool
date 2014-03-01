/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : subnettool_test.go

* Purpose :

* Creation Date : 02-27-2014

* Last Modified : Fri Feb 28 17:45:03 2014

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package subnettool

import (
	"net"
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
