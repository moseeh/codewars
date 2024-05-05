package kata

import "strings"

func ToCamelCase(s string) string {
	st := ""
	stst := ""
	ss := strings.Split(s, "_")
	st = ss[0]
	for i := 1; i < len(ss); i++ {
		ss[i] = strings.Title(ss[i])
		st += ss[i]
	}
	sts := strings.Split(st, "-")
	stst = sts[0]
	for i := 1; i < len(sts); i++ {
		sts[i] = strings.Title(sts[i])
		stst += sts[i]
	}
	return stst
}
