package main

import "fmt"

type path struct {
	to                     *vertex
	index, distance, level int
}

func (p *path) key() string {
	return fmt.Sprintf("%d,%d,%d", p.to.position.x, p.to.position.y, p.level)
}
