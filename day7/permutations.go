package main

func getPermutations(n int) [][]int {
	if n == 1 {
		return [][]int{{0}}
	}

	subPermutations := getPermutations(n - 1)
	var permutations [][]int
	for _, sub := range subPermutations {
		for i := n - 1; i >= 0; i-- {
			perm := make([]int, n)
			copy(perm, sub[:i])
			perm[i] = n - 1
			copy(perm[i+1:], sub[i:])
			permutations = append(permutations, perm)
		}
	}

	return permutations
}
