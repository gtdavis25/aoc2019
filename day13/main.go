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

	var buffer []int
	tiles := make(map[point]tileID)
	computer := intcode.NewComputer(program)
	computer.OnOutput(func(value int) {
		buffer = append(buffer, value)
		if len(buffer) == 3 {
			p := point{x: buffer[0], y: buffer[1]}
			tiles[p] = tileID(buffer[2])
			buffer = buffer[:0]
		}
	})

	computer.Execute()
	var count int
	for _, tile := range tiles {
		if tile == block {
			count++
		}
	}

	fmt.Println("Part 1:", count)
	computer = intcode.NewComputer(program)
	var score, ballX, paddleX int
	computer.SetValue(0, 2)
	tiles = make(map[point]tileID)
	computer.OnOutput(func(value int) {
		buffer = append(buffer, value)
		if len(buffer) == 3 {
			if x := buffer[0]; x == -1 {
				score = buffer[2]
			} else if tileID(buffer[2]) == ball {
				ballX = x
			} else if tileID(buffer[2]) == paddle {
				paddleX = x
			}

			buffer = buffer[:0]
		}
	})

	for !computer.Halted() {
		computer.Execute()
		computer.Input().Write(sign(ballX - paddleX))
	}

	fmt.Println("Part 2:", score)
}

func sign(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	}

	return 0
}
