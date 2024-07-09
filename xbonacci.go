package kata

func Xbonacci(signature []int64, n int) []int64 {
	length := len(signature)
	if n <= length {
		return signature[:n]
	}

	result := make([]int64, n)
	copy(result, signature)

	for i := length; i < n; i++ {
		sum := int64(0)
		for j := i - length; j < i; j++ {
			sum += result[j]
		}
		result[i] = sum
	}

	return result
}
