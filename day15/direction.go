package main

import "fmt"

type direction int

const (
	north direction = 1 + iota
	south
	east
	west
)

func (d direction) reverse() direction {
	switch d {
	case north:
		return south

	case south:
		return north

	case east:
		return west

	case west:
		return east

	default:
		panic(fmt.Sprintf("invalid direction: %d", d))
	}
}
