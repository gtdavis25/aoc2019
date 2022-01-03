package main

import "fmt"

type point struct {
	x, y int
}

func (p point) adjacent(direction direction) point {
	switch direction {
	case north:
		return point{p.x, p.y - 1}

	case south:
		return point{p.x, p.y + 1}

	case east:
		return point{p.x - 1, p.y}

	case west:
		return point{p.x + 1, p.y}

	default:
		panic(fmt.Sprintf("invalid direction: %d", direction))
	}
}
