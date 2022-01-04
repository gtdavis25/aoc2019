package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)
	numbers := make([]int, len(input))
	for i := range numbers {
		fmt.Sscan(input[i:i+1], &numbers[i])
	}

	result := numbers
	for i := 0; i < 100; i++ {
		result = applyFFT(result, 0)
	}

	w := new(strings.Builder)
	for i := 0; i < 8; i++ {
		fmt.Fprint(w, result[i])
	}

	fmt.Println("Part 1:", w.String())
	var offset int
	for i := 0; i < 7; i++ {
		offset = offset*10 + numbers[i]
	}

	result = repeat(numbers, 10000)
	for i := 0; i < 100; i++ {
		result = applyFFT(result, offset)
	}

	w.Reset()
	for i := 0; i < 8; i++ {
		fmt.Fprint(w, result[offset+i])
	}

	fmt.Println("Part 2:", w.String())
}

func applyFFT(numbers []int, offset int) []int {
	result := make([]int, len(numbers))
	prefix := prefix(numbers)
	sign := 1
	for i := offset; i < len(result); i++ {
		for j := i; j < len(numbers); j += 2 * (i + 1) {
			result[i] += sign * (prefix[min(len(prefix)-1, j+i+1)] - prefix[j])
			sign = -sign
		}

		result[i] %= 10
		if result[i] < 0 {
			result[i] = -result[i]
		}
	}

	return result
}

func prefix(numbers []int) []int {
	var sum int
	prefix := make([]int, len(numbers)+1)
	for i := range numbers {
		sum += numbers[i]
		prefix[i+1] = sum
	}

	return prefix
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func repeat(sequence []int, times int) []int {
	result := make([]int, len(sequence)*times)
	for i := range result {
		result[i] = sequence[i%len(sequence)]
	}

	return result
}
