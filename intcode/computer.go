package intcode

import "fmt"

type Computer struct {
	pointer int
	memory  []int
}

func NewComputer(program []int) *Computer {
	computer := new(Computer)
	computer.memory = make([]int, len(program))
	copy(computer.memory, program)
	return computer
}

func (c *Computer) Execute() {
	for op := c.memory[c.pointer]; op != 99; op = c.memory[c.pointer] {
		switch op {
		case 1:
			acc := c.GetValue(c.GetValue(c.pointer + 1))
			acc += c.GetValue(c.GetValue(c.pointer + 2))
			c.SetValue(c.GetValue(c.pointer+3), acc)
			c.pointer += 4

		case 2:
			acc := c.GetValue(c.GetValue(c.pointer + 1))
			acc *= c.GetValue(c.GetValue(c.pointer + 2))
			c.SetValue(c.GetValue(c.pointer+3), acc)
			c.pointer += 4

		default:
			panic(fmt.Sprintf("Invalid opcode: %d", op))
		}
	}
}

func (c *Computer) GetValue(address int) int {
	return c.memory[address]
}

func (c *Computer) SetValue(address, value int) {
	c.memory[address] = value
}
