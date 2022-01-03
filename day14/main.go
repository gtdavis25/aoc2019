package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reactions []reaction
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		line := scanner.Text()
		parts := strings.Split(line, " => ")
		var r reaction
		for _, s := range strings.Split(parts[0], ", ") {
			var input resource
			fmt.Sscan(s, &input.quantity, &input.chemical)
			r.inputs = append(r.inputs, input)
		}

		fmt.Sscan(parts[1], &r.output.quantity, &r.output.chemical)
		reactions = append(reactions, r)
	}

	fmt.Println("Part 1:", getCost("FUEL", 1, reactions, make(map[string]int)))
	fmt.Println("Part 2:", findMax(func(n int) bool {
		return getCost("FUEL", n, reactions, make(map[string]int)) <= 1e12
	}))
}

func getCost(chemical string, quantity int, reactions []reaction, resources map[string]int) int {
	if chemical == "ORE" {
		return quantity
	}

	if resources[chemical] >= quantity {
		resources[chemical] -= quantity
		return 0
	}

	quantity -= resources[chemical]
	resources[chemical] = 0
	reaction := find(reactions, chemical)
	numReactions := quantity / reaction.output.quantity
	if quantity%reaction.output.quantity > 0 {
		numReactions++
	}

	var total int
	for _, r := range reaction.inputs {
		total += getCost(r.chemical, r.quantity*numReactions, reactions, resources)
	}

	resources[reaction.output.chemical] = numReactions*reaction.output.quantity - quantity
	return total
}

func find(reactions []reaction, chemical string) reaction {
	for i := range reactions {
		if reactions[i].output.chemical == chemical {
			return reactions[i]
		}
	}

	panic("no such reaction")
}

func findMax(condition func(int) bool) int {
	min, max := 0, 1
	for condition(max) {
		max *= 2
	}

	for min+1 < max {
		if mid := (min + max) / 2; condition(mid) {
			min = mid
		} else {
			max = mid
		}
	}

	return min
}
