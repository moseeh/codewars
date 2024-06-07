package kata

func Snail(snaipMap [][]int) []int {
	// Check if the matrix is empty
	if len(snaipMap) == 0 {
		return []int{}
	}

	// Initialize the result slice
	result := []int{}

	// Initialize the boundaries
	top, bottom := 0, len(snaipMap)-1
	left, right := 0, len(snaipMap[0])-1

	// Traverse the matrix in a spiral order
	for top <= bottom && left <= right {
		// Traverse from left to right along the top boundary
		for i := left; i <= right; i++ {
			result = append(result, snaipMap[top][i])
		}
		top++

		// Traverse from top to bottom along the right boundary
		for i := top; i <= bottom; i++ {
			result = append(result, snaipMap[i][right])
		}
		right--

		// Traverse from right to left along the bottom boundary, if still within bounds
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, snaipMap[bottom][i])
			}
			bottom--
		}

		// Traverse from bottom to top along the left boundary, if still within bounds
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, snaipMap[i][left])
			}
			left++
		}
	}

	return result
}
