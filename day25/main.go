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
	}

	program, err := intcode.ReadProgram(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	computer := intcode.NewComputer(program)
	var buffer []byte
	computer.OnOutput(func(value int) {
		if value == '\n' {
			fmt.Println(string(buffer))
			buffer = buffer[:0]
		} else {
			buffer = append(buffer, byte(value))
		}
	})

	scanner := bufio.NewScanner(os.Stdin)
	computer.Execute()
	for !computer.Halted() {
		scanner.Scan()
		line := scanner.Text()
		for _, c := range line {
			computer.Input().Write(int(c))
		}

		computer.Input().Write('\n')
		computer.Execute()
	}
}
