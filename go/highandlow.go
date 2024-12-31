package kata

import (
	"strconv"
	"strings"
)

func HighAndLow(in string) string {

	sArr := strings.Fields(in)
	n, _ := strconv.Atoi(sArr[0])
	n2, _ := strconv.Atoi(sArr[0])
	for _, v := range sArr {
		num, _ := strconv.Atoi(v)
		if num > n {
			n = num
		}
		if num < n2 {
			n2 = num
		}
	}
	return strconv.Itoa(n) + " " + strconv.Itoa(n2)
}
