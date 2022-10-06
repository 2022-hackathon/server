package module

import "strings"

func EditDigit(num string) string {

	numSlice := strings.Split(num, "")

	cnt := 0
	flag := false

	for _, v := range numSlice {

		if flag {
			cnt++
		}

		if v == "." {
			flag = true
		}
	}

	num = strings.ReplaceAll(num, ".", "")

	switch cnt {
	case 0:
		num += "00"
	case 1:
		num += "0"
	default:
		return num
	}

	return num
}
