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

	fmt.Println("Part 1:", run(program, 1))
	fmt.Println("Part 2:", run(program, 2))
}

func run(program []int, input int) int {
	computer := intcode.NewComputer(program)
	computer.Input().Write(input)
	var result int
	computer.OnOutput(func(value int) { result = value })
	computer.Execute()
	return result
}
