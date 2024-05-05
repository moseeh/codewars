package kata

import "math"

func PowersOfTwo(n int) []uint64 {
	result := make([]uint64, n+1)
	for i := 0; i <= n; i++ {
		result[i] = uint64(math.Pow(2, float64(i)))
	}
	return result

}
