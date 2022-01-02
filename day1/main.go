package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var modules []int
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		var n int
		fmt.Sscanf(scanner.Text(), "%d", &n)
		modules = append(modules, n)
	}

	var total int
	for i := range modules {
		total += getFuel(modules[i])
	}

	fmt.Println("Part 1:", total)
	total = 0
	for i := range modules {
		for mass := getFuel(modules[i]); mass > 0; mass = getFuel(mass) {
			total += mass
		}
	}

	fmt.Println("Part 2:", total)
}

func getFuel(mass int) int {
	if fuel := mass/3 - 2; fuel > 0 {
		return fuel
	}

	return 0
}
