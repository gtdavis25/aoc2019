package intcode

import (
	"fmt"
	"io"
	"strings"
)

func ReadProgram(r io.Reader) ([]int, error) {
	var input string
	fmt.Fscan(r, &input)
	words := strings.Split(input, ",")
	program := make([]int, len(words))
	for i := range program {
		if _, err := fmt.Sscan(words[i], &program[i]); err != nil {
			return nil, fmt.Errorf("at index %d: %s", i, err)
		}
	}

	return program, nil
}
