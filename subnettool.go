/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : subnettool.go

* Purpose :

* Creation Date : 02-20-2014

* Last Modified : Thu Feb 20 18:16:43 2014

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package subnettool

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func Ntoa(ipnr int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

func Aton(ipnr net.IP) int64 {
	bits := strings.Split(ipnr.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

func Mtos(str string) int {
	mask := net.IPMask(net.ParseIP(str).To4())
	size, _ := mask.Size()
	return size
}

func Stom(size int) string {
	divisor := size / 8
	remainder := size % 8
	var s string
	var f bool
	for i := 1; i <= 4; i++ {
		if f {
			s = s + "0."
			continue
		}
		if i <= divisor {
			s = s + "255."
		} else {
			s = s + fmt.Sprint(2<<7-2<<(7-uint(remainder))) + "."
			f = true
		}
	}
	return s[:len(s)-1]
}
