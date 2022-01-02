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
	words := strings.Split(scanner.Text(), ",")
	program := make([]int, len(words))
	for i, word := range words {
		fmt.Sscan(word, &program[i])
	}

	var max int
	for _, perm := range getPermutations(5) {
		if result := run(program, perm); result > max {
			max = result
		}
	}

	fmt.Println("Part 1:", max)
	for _, perm := range getPermutations(5) {
		for i := range perm {
			perm[i] += 5
		}

		if result := runWithFeedback(program, perm); result > max {
			max = result
		}
	}

	fmt.Println("Part 2:", max)
}

func run(program []int, phase []int) int {
	computers := make([]*intcode.Computer, 5)
	for i := range computers {
		computer := intcode.NewComputer(program)
		computer.Input().Write(phase[i])
		computers[i] = computer
	}

	for i := 0; i < len(computers)-1; i++ {
		next := computers[i+1]
		computers[i].OnOutput(func(value int) {
			next.Input().Write(value)
		})
	}

	var result int
	computers[len(computers)-1].OnOutput(func(value int) { result = value })
	computers[0].Input().Write(0)
	for _, computer := range computers {
		computer.Execute()
	}

	return result
}

func runWithFeedback(program []int, phase []int) int {
	computers := make([]*intcode.Computer, 5)
	for i := range computers {
		computer := intcode.NewComputer(program)
		computer.Input().Write(phase[i])
		computers[i] = computer
	}

	var result int
	for i := 0; i < len(computers); i++ {
		next := computers[(i+1)%len(computers)]
		computers[i].OnOutput(func(value int) {
			result = value
			next.Input().Write(value)
		})
	}

	computers[0].Input().Write(0)
	for !computers[4].Halted() {
		for i := range computers {
			computers[i].Execute()
		}
	}

	return result
}
