package intcode_test

import (
	"aoc2019/intcode"
	"testing"
)

func TestExecute(t *testing.T) {
	program := []int{3, 0, 4, 0, 99}
	computer := intcode.NewComputer(program)
	computer.Input().Write(123)
	var result int
	computer.OnOutput(func(value int) {
		result = value
	})

	computer.Execute()
	if result != 123 {
		t.Errorf("got %d, want %d", result, 123)
	}
}
