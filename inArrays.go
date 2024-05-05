package kata

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	r := []string{}
	found := make(map[string]bool)
	for _, v1 := range array1 {
		for _, v2 := range array2 {
			if strings.Contains(v2, v1) && !found[v1] {
				r = append(r, v1)
				found[v1] = true
				break
			}
		}
	}
	sort.Strings(r)
	return r
}
