package kata

import "sort"

func SortByLength(arr []string) []string {
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})
	return arr
}
