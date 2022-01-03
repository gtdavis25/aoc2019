package main

import "fmt"

type point struct {
	x, y, z int
}

func (p point) add(other point) point {
	return point{
		x: p.x + other.x,
		y: p.y + other.y,
		z: p.z + other.z,
	}
}

func (p point) minus(other point) point {
	return point{
		x: p.x - other.x,
		y: p.y - other.y,
		z: p.z - other.z,
	}
}

func (p point) String() string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, p.z)
}
