package main

type point struct {
	x, y int
}

func (p point) add(q point) point {
	return point{x: p.x + q.x, y: p.y + q.y}
}
