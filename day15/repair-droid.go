package main

import "aoc2019/intcode"

type repairDroid struct {
	computer *intcode.Computer
}

func newDroid(program []int) *repairDroid {
	droid := new(repairDroid)
	droid.computer = intcode.NewComputer(program)
	return droid
}

func (d *repairDroid) move(direction direction) moveResult {
	var result moveResult
	d.computer.Input().Write(int(direction))
	d.computer.OnOutput(func(value int) {
		result = moveResult(value)
	})

	d.computer.Execute()
	return result
}
