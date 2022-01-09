package main

import "strings"

type tile struct {
	cells uint32
}

func (t tile) isAlive(x, y int) bool {
	if x < 0 || x >= 5 || y < 0 || y >= 5 {
		return false
	}

	var offset = 5*y + x
	return (t.cells>>offset)&1 == 1
}

func (t tile) setAlive(x, y int) tile {
	if x < 0 || x >= 5 || y < 0 || y >= 5 {
		panic("out of bounds")
	}

	var offset = 5*y + x
	t.cells |= 1 << offset
	return t
}

func (t tile) nextState() tile {
	var next tile
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			var adjacent int
			for _, p := range []struct{ x, y int }{
				{x, y - 1},
				{x - 1, y},
				{x + 1, y},
				{x, y + 1},
			} {
				if t.isAlive(p.x, p.y) {
					adjacent++
				}
			}

			if adjacent == 1 || adjacent == 2 && !t.isAlive(x, y) {
				next = next.setAlive(x, y)
			}
		}
	}

	return next
}

func parseTile(lines []string) tile {
	var t tile
	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == '#' {
				offset := 5*y + x
				t.cells |= 1 << offset
			}
		}
	}

	return t
}

func (t tile) String() string {
	w := new(strings.Builder)
	for y := 0; y < 5; y++ {
		if y > 0 {
			w.WriteByte('\n')
		}

		for x := 0; x < 5; x++ {
			if t.isAlive(x, y) {
				w.WriteByte('#')
			} else {
				w.WriteByte('.')
			}
		}
	}

	return w.String()
}
