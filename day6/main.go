package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	planets := make(map[string]*planet)
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		words := strings.Split(scanner.Text(), ")")
		if _, ok := planets[words[0]]; !ok {
			planets[words[0]] = new(planet)
		}

		if _, ok := planets[words[1]]; !ok {
			planets[words[1]] = new(planet)
		}

		parent := planets[words[0]]
		child := planets[words[1]]
		parent.orbits = append(parent.orbits, child)
		child.parent = parent
	}

	var total int
	for _, p := range planets {
		total += p.totalOrbits()
	}

	fmt.Println("Part 1:", total)
	fmt.Println("Part 2;", planets["YOU"].parent.distanceTo(planets["SAN"].parent))
}
