package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var lines []string
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		lines = append(lines, scanner.Text())
	}

	t := parseTile(lines)
	seen := map[uint32]bool{t.cells: true}
	for {
		t = t.nextState()
		if seen[t.cells] {
			break
		}

		seen[t.cells] = true
	}

	fmt.Println("Part 1:", t.cells)
	cells := make(map[cell]bool)
	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == '#' {
				cells[cell{x, y, 0}] = true
			}
		}
	}

	for i := 0; i < 200; i++ {
		adjacent := make(map[cell]int)
		for c := range cells {
			for _, a := range c.adjacentCells() {
				adjacent[a]++
			}
		}

		next := make(map[cell]bool)
		for c, neighbours := range adjacent {
			if neighbours == 1 || neighbours == 2 && !cells[c] {
				next[c] = true
			}
		}

		cells = next
	}

	var count int
	for range cells {
		count++
	}

	fmt.Println("Part 2:", count)
}
