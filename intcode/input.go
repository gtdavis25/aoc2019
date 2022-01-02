package intcode

type Input interface {
	Write(value int)
	Read() (int, bool)
}

type input struct {
	buffer []int
}

func NewInput() Input {
	return new(input)
}

func (i *input) Write(value int) {
	i.buffer = append(i.buffer, value)
}

func (i *input) Read() (int, bool) {
	if len(i.buffer) < 1 {
		return -1, false
	}

	value := i.buffer[0]
	i.buffer = i.buffer[1:]
	return value, true
}
