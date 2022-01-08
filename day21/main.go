package main

import (
	"aoc2019/intcode"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	program, err := intcode.ReadProgram(file)
	if err != nil {
		log.Fatal(err)
	}

	computer := intcode.NewComputer(program)
	var buffer []byte
	var result int
	computer.OnOutput(func(value int) {
		if value == '\n' {
			fmt.Println(string(buffer))
			buffer = buffer[:0]
		} else if value > 255 {
			result = value
		} else {
			buffer = append(buffer, byte(value))
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

	fmt.Println("Result:", result)
}
