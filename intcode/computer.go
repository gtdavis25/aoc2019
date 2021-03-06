package intcode

import "fmt"

type Computer struct {
	pointer, paramModes, base int
	memory                    []int
	input                     Input
	outputFunc                func(int)
}

func NewComputer(program []int) *Computer {
	computer := new(Computer)
	computer.memory = make([]int, len(program))
	copy(computer.memory, program)
	computer.input = NewInput()
	return computer
}

func (c *Computer) Execute() {
	for {
		op := c.GetValue(c.pointer)
		c.paramModes = op / 100
		op %= 100
		switch op {
		case 1:
			acc := c.readArgument(c.pointer + 1)
			acc += c.readArgument(c.pointer + 2)
			c.writeArgument(c.pointer+3, acc)
			c.pointer += 4

		case 2:
			acc := c.readArgument(c.pointer + 1)
			acc *= c.readArgument(c.pointer + 2)
			c.writeArgument(c.pointer+3, acc)
			c.pointer += 4

		case 3:
			if v, ok := c.input.Read(); ok {
				c.writeArgument(c.pointer+1, v)
				c.pointer += 2
			} else {
				return
			}

		case 4:
			if c.outputFunc == nil {
				panic("output is nil")
			}

			value := c.readArgument(c.pointer + 1)
			c.outputFunc(value)
			c.pointer += 2

		case 5:
			v := c.readArgument(c.pointer + 1)
			if v != 0 {
				c.pointer = c.readArgument(c.pointer + 2)
			} else {
				c.pointer += 3
			}

		case 6:
			v := c.readArgument(c.pointer + 1)
			if v == 0 {
				c.pointer = c.readArgument(c.pointer + 2)
			} else {
				c.pointer += 3
			}

		case 7:
			arg1 := c.readArgument(c.pointer + 1)
			arg2 := c.readArgument(c.pointer + 2)
			if arg1 < arg2 {
				c.writeArgument(c.pointer+3, 1)
			} else {
				c.writeArgument(c.pointer+3, 0)
			}

			c.pointer += 4

		case 8:
			arg1 := c.readArgument(c.pointer + 1)
			arg2 := c.readArgument(c.pointer + 2)
			if arg1 == arg2 {
				c.writeArgument(c.pointer+3, 1)
			} else {
				c.writeArgument(c.pointer+3, 0)
			}

			c.pointer += 4

		case 9:
			offset := c.readArgument(c.pointer + 1)
			c.base += offset
			c.pointer += 2

		case 99:
			return

		default:
			panic(fmt.Sprintf("Invalid opcode: %d", op))
		}
	}
}

func (c *Computer) readArgument(address int) int {
	paramMode := c.paramModes % 10
	c.paramModes /= 10
	switch paramMode {
	case 0:
		return c.GetValue(c.GetValue(address))

	case 1:
		return c.GetValue(address)

	case 2:
		return c.GetValue(c.GetValue(address) + c.base)

	default:
		panic(fmt.Sprintf("invalid parameter mode %d", paramMode))
	}
}

func (c *Computer) writeArgument(address, value int) {
	paramMode := c.paramModes % 10
	c.paramModes /= 10
	switch paramMode {
	case 0:
		c.SetValue(c.GetValue(address), value)

	case 2:
		c.SetValue(c.GetValue(address)+c.base, value)
	}
}

func (c *Computer) GetValue(address int) int {
	if address >= len(c.memory) {
		return 0
	}

	return c.memory[address]
}

func (c *Computer) SetValue(address, value int) {
	for address >= len(c.memory) {
		c.memory = append(c.memory, 0)
	}

	c.memory[address] = value
}

func (c *Computer) Input() Input {
	return c.input
}

func (c *Computer) OnOutput(outputFunc func(int)) {
	c.outputFunc = outputFunc
}

func (c *Computer) Halted() bool {
	return c.GetValue(c.pointer) == 99
}
