package main

import (
	"fmt"
	"strings"
)

type layer struct {
	width, height int
	rows          [][]byte
}

func newLayer(width, height int) layer {
	buffer := make([]byte, width*height)
	rows := make([][]byte, height)
	for i := range rows {
		rows[i] = buffer[i*width : (i+1)*width]
	}

	return layer{
		width:  width,
		height: height,
		rows:   rows,
	}
}

func (l layer) getValue(x, y int) byte {
	return l.rows[y][x]
}

func (l layer) setValue(x, y int, value byte) {
	l.rows[y][x] = value
}

func (l layer) count(filterfunc func(b byte) bool) int {
	var count int
	for y := 0; y < l.height; y++ {
		for x := 0; x < l.width; x++ {
			if filterfunc(l.getValue(x, y)) {
				count++
			}
		}
	}

	return count
}

func (l layer) zeroes() int {
	return l.count(func(b byte) bool { return b == 0 })
}

func (l layer) ones() int {
	return l.count(func(b byte) bool { return b == 1 })
}

func (l layer) twos() int {
	return l.count(func(b byte) bool { return b == 2 })
}

func (l layer) String() string {
	w := new(strings.Builder)
	for y := 0; y < l.height; y++ {
		if y > 0 {
			fmt.Fprintln(w)
		}

		for x := 0; x < l.width; x++ {
			switch l.getValue(x, y) {
			case 0:
				w.WriteByte(' ')

			case 1:
				w.WriteByte('#')

			default:
				w.WriteByte('?')
			}
		}
	}

	return w.String()
}
