package main

import (
	"fmt"
	"strings"
)

type maze struct {
	width, height int
	rows          [][]tile
	keys          keySet
}

func newMaze(lines []string) maze {
	width := len(lines[0])
	height := len(lines)
	tiles := make([]tile, width*height)
	rows := make([][]tile, height)
	var keys keySet
	for y := range rows {
		rows[y] = tiles[y*width : (y+1)*width]
		for x := range rows[y] {
			rows[y][x] = tile(lines[y][x])
			if rows[y][x].isKey() {
				keys = keys.add(rows[y][x])
			}
		}
	}

	return maze{
		width:  width,
		height: height,
		rows:   rows,
		keys:   keys,
	}
}

func (m maze) getTile(p point) tile {
	if p.x < 0 || p.x >= m.width || p.y < 0 || p.y >= m.height {
		return 0
	}

	return m.rows[p.y][p.x]
}

func (m maze) String() string {
	w := new(strings.Builder)
	for y := range m.rows {
		if y > 0 {
			fmt.Fprintln(w)
		}

		for x := range m.rows[y] {
			w.WriteByte(byte(m.rows[y][x]))
		}
	}

	return w.String()
}
