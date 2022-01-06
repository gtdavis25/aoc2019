package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)
	words := strings.Split(input, ",")
	program := make([]int, len(words))
	for i := range words {
		fmt.Sscan(words[i], &program[i])
	}

	beam := tractorBeam{program: program}
	var count int
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if beam.contains(point{x, y}) {
				count++
			}
		}
	}

	fmt.Println("Part 1:", count)
	top, left, size := 0, 0, 100
	for !beam.contains(point{left + size - 1, top}) || !beam.contains(point{left, top + size - 1}) {
		for !beam.contains(point{left, top + size}) {
			left++
		}

		for !beam.contains(point{left + size, top}) {
			top++
		}
	}

	for moved := true; moved; {
		moved = false
		for dy := 5; dy >= 0; dy-- {
			for dx := 5; dx >= 0; dx-- {
				if dx == 0 && dy == 0 {
					break
				}

				x, y := left-dx, top-dy
				if beam.contains(point{x, y + size - 1}) && beam.contains(point{x + size - 1, y}) {
					moved = true
					left, top = x, y
				}
			}
		}
	}

	fmt.Println("Part 2:", 10000*left+top)
}
