package main

import (
	"fmt"
	"math"
)

func countDivisors(n int) int {
	count := 0
	sqrtN := int(math.Sqrt(float64(n)))
	
	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			if i*i == n {
				count++
			} else {
				count += 2
			}
		}
	}
	return count
}

func smallestNumberWithNDivisors(n int) int {
	num := 1
	for {
		if countDivisors(num) == n {
			return num
		}
		num++
	}
}

func main() {
	testCases := []int{6, 10, 4, 8, 12, 100}
	
	for _, n := range testCases {
		result := smallestNumberWithNDivisors(n)
		fmt.Printf("Smallest number with %d divisors: %d\n", n, result)
	}
}