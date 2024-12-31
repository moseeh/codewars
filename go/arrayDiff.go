package kata

func ArrayDiff(a, b []int) []int {
	result := []int{}
	bMap := make(map[int]bool)

	// Create a map of values in b for faster lookup
	for _, val := range b {
		bMap[val] = true
	}

	// Iterate through a, if value is not in b, add it to result
	for _, val := range a {
		if !bMap[val] {
			result = append(result, val)
		}
	}

	return result
}
