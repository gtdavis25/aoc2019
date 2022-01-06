package main

import "aoc2019/intcode"

type tractorBeam struct {
	program []int
}

func (t tractorBeam) contains(p point) bool {
	computer := intcode.NewComputer(t.program)
	var result bool
	computer.OnOutput(func(value int) { result = value == 1 })
	computer.Input().Write(p.x)
	computer.Input().Write(p.y)
	computer.Execute()
	return result
}
