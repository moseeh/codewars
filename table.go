package kata

func MultiplicationTable(size int) [][]int {
	grid := [][]int{}
	for i := 1; i <= size; i++ {
		arr := []int{}
		for j := 1; j <= size; j++ {
			arr = append(arr, i*j)
		}
		grid = append(grid, arr)
	}
	return grid
}
