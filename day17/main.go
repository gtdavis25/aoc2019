package main

import (
	"aoc2019/intcode"
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	program, err := intcode.ReadProgram(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	computer := intcode.NewComputer(program)
	var buffer []byte
	var lines []string
	computer.OnOutput(func(value int) {
		if value == '\n' && len(buffer) > 0 {
			lines = append(lines, string(buffer))
			buffer = buffer[:0]
		} else if value != '\n' {
			buffer = append(buffer, byte(value))
		}
	})

	computer.Execute()
	var alignment int
	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[y])-1; x++ {
			if isJunction(lines, point{x, y}) {
				alignment += x * y
			}
		}
	}

	fmt.Println("Part 1:", alignment)
	computer = intcode.NewComputer(program)
	computer.SetValue(0, 2)
	var result int
	computer.OnOutput(func(value int) {
		if value > 255 {
			result = value
		} else if value != '\n' {
			buffer = append(buffer, byte(value))
		} else {
			fmt.Println(string(buffer))
			buffer = buffer[:0]
		}
	})

	computer.Execute()
	scanner := bufio.NewScanner(os.Stdin)
	for !computer.Halted() {
		scanner.Scan()
		for _, c := range scanner.Text() {
			computer.Input().Write(int(c))
		}

		computer.Input().Write('\n')
		computer.Execute()
	}

	fmt.Println("Part 2:", result)
}

func isJunction(lines []string, p point) bool {
	for _, p := range []point{
		p,
		{p.x, p.y - 1},
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y + 1},
	} {
		if lines[p.y][p.x] != '#' {
			return false
		}
	}

	return true
}
