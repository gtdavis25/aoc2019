package main

import (
	"aoc2019/intcode"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	fmt.Fscan(os.Stdin, &input)
	words := strings.Split(input, ",")
	program := make([]int, len(words))
	for i := range program {
		fmt.Sscan(words[i], &program[i])
	}

	panels := make(map[point]bool)
	run(program, panels)
	var count int
	for range panels {
		count++
	}

	fmt.Println("Part 1:", count)
	panels = map[point]bool{{0, 0}: true}
	run(program, panels)
	var minX, minY, maxX, maxY int
	for p, white := range panels {
		if white {
			minX = min(p.x, minX)
			minY = min(p.y, minY)
			maxX = max(p.x, maxX)
			maxY = max(p.y, maxY)
		}
	}

	w := new(strings.Builder)
	for y := minY; y <= maxY; y++ {
		if y > minY {
			w.WriteByte('\n')
		}

		for x := minX; x <= maxX; x++ {
			if panels[point{x, y}] {
				w.WriteByte('#')
			} else {
				w.WriteByte(' ')
			}
		}
	}

	fmt.Println("Part 2:\n", w.String())
}

func run(program []int, panels map[point]bool) {
	robot := intcode.NewComputer(program)
	var output []int
	robot.OnOutput(func(value int) { output = append(output, value) })
	var position point
	direction := point{x: 0, y: -1}
	for !robot.Halted() {
		if panels[position] {
			robot.Input().Write(1)
		} else {
			robot.Input().Write(0)
		}

		output = output[:0]
		robot.Execute()
		if len(output) != 2 {
			panic("unexpected output")
		}

		switch output[0] {
		case 0:
			panels[position] = false

		case 1:
			panels[position] = true

		default:
			panic(fmt.Errorf("unexpected output: %d", output[0]))
		}

		switch output[1] {
		case 0:
			direction = direction.turnLeft()

		case 1:
			direction = direction.turnRight()

		default:
			panic(fmt.Errorf("unexpected output: %d", output[0]))
		}

		position = position.add(direction)
	}
}

func min(a, b int) int {
	if b < a {
		return b
	}

	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}

	return a
}
