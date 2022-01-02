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
	for i := range words {
		fmt.Sscanf(words[i], "%d", &program[i])
	}

	fmt.Println("Part 1:", run(program, 12, 2))
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			if run(program, noun, verb) == 19690720 {
				fmt.Println("Part 2:", noun*100+verb)
				return
			}
		}
	}
}

func run(program []int, noun, verb int) int {
	computer := intcode.NewComputer((program))
	computer.SetValue(1, noun)
	computer.SetValue(2, verb)
	computer.Execute()
	return computer.GetValue(0)
}
