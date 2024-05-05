package kata

func Parse(data string) []int {
	n := 0
	arr := []int{}
	for _, v := range data {
		if v == 'i' {
			n++
		}
		if v == 'd' {
			n--
		}
		if v == 's' {
			n = n * n
		}
		if v == 'o' {
			arr = append(arr, n)
		}
	}
	return arr
}
