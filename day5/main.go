package main

import (
	"aoc2019/intcode"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := strings.Split(scanner.Text(), ",")
	program := make([]int, len(s))
	for i := range s {
		fmt.Sscanf(s[i], "%d", &program[i])
	}

	fmt.Println("Part 1:", run(program, 1))
	fmt.Println("Part 2:", run(program, 5))
}

func run(program []int, input int) int {
	computer := intcode.NewComputer(program)
	computer.Input().Write(input)
	var result int
	computer.OnOutput(func(value int) { result = value })
	computer.Execute()
	return result
}
