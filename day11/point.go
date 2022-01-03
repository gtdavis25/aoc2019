package main

type point struct {
	x, y int
}

func (p point) turnLeft() point {
	return point{x: p.y, y: -p.x}
}

func (p point) turnRight() point {
	return point{x: -p.y, y: p.x}
}

func (p point) add(other point) point {
	return point{x: p.x + other.x, y: p.y + other.y}
}
