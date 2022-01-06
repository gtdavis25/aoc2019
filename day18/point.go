package main

import "fmt"

type point struct {
	x, y int
}

func (p point) adjacentPoints() []point {
	return []point{
		{p.x, p.y - 1},
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y + 1},
	}
}

func (p point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}
