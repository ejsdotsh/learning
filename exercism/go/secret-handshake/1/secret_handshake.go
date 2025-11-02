package secret

import (
	"slices"
	"strconv"
)

func checkBit(n uint, k int) bool {
	mask := uint(1 << k)
	return (n & mask) != 0
}

func reverseResult(res []string) []string {
	var revRes []string
	for _, v := range slices.Backward(res) {
		revRes = append(revRes, v)
	}
	return revRes
}

func Handshake(code uint) []string {
	var res []string
	binStrLen := len(strconv.FormatUint(uint64(code), 2))

	for i := range binStrLen {
		switch true {
		case i == 0 && checkBit(code, i):
			res = append(res, "wink")
		case i == 1 && checkBit(code, i):
			res = append(res, "double blink")
		case i == 2 && checkBit(code, i):
			res = append(res, "close your eyes")
		case i == 3 && checkBit(code, i):
			res = append(res, "jump")
		case i == 4 && checkBit(code, i):
			res = reverseResult(res)
		}
	}

	return res
}
