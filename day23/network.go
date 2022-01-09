package main

import "aoc2019/intcode"

type network struct {
	computers []*intcode.Computer
	packets   map[int][]packet
}

func newNetwork(program []int) *network {
	computers := make([]*intcode.Computer, 50)
	packets := make(map[int][]packet)
	for i := range computers {
		computers[i] = intcode.NewComputer(program)
		computers[i].Input().Write(i)
		var buffer []int
		computers[i].OnOutput(func(value int) {
			buffer = append(buffer, value)
			if len(buffer) == 3 {
				packets[buffer[0]] = append(packets[buffer[0]], packet{buffer[1], buffer[2]})
				buffer = buffer[:0]
			}
		})
	}

	return &network{
		computers: computers,
		packets:   packets,
	}
}

func (n *network) run() {
	for i := range n.computers {
		if len(n.packets[i]) > 0 {
			p := n.packets[i][0]
			n.packets[i] = n.packets[i][1:]
			n.computers[i].Input().Write(p.x)
			n.computers[i].Input().Write(p.y)
		} else {
			n.computers[i].Input().Write(-1)
		}

		n.computers[i].Execute()
	}
}

func (n *network) idle() bool {
	for i := range n.computers {
		if len(n.packets[i]) > 0 {
			return false
		}
	}

	return true
}
