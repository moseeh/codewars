package kata

func FindUniq(arr []float32) float32 {
	count := make(map[float32]int)
	for _, v := range arr {
		count[v]++
	}
	for num, times := range count {
		if times == 1 {
			return num
		}
	}

	return -1
}
