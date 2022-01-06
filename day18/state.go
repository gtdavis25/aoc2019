package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type state struct {
	positions []*vertex
	keys      keySet
	distance  int
}

func (s state) nextStates() []state {
	var states []state
	for i, p := range s.positions {
		seen := make(map[*vertex]bool)
		queue := edgeQueue{edge{to: p, length: 0}}
		for len(queue) > 0 {
			current := heap.Pop(&queue).(edge)
			if seen[current.to] {
				continue
			}

			seen[current.to] = true
			if current.to.tile.isKey() && !s.keys.has(current.to.tile) {
				newPositions := make([]*vertex, len(s.positions))
				copy(newPositions, s.positions)
				newPositions[i] = current.to
				states = append(states, state{
					positions: newPositions,
					keys:      s.keys.add(current.to.tile),
					distance:  s.distance + current.length,
				})

				continue
			}

			for _, next := range current.to.edges {
				if seen[next.to] || next.to.tile.isDoor() && !s.keys.canOpen(next.to.tile) {
					continue
				}

				heap.Push(&queue, edge{
					to:     next.to,
					length: current.length + next.length,
				})
			}
		}

	}

	return states
}

func (s state) String() string {
	w := new(strings.Builder)
	for _, p := range s.positions {
		w.WriteByte(byte(p.tile))
	}

	fmt.Fprintf(w, " %s", s.keys)
	return w.String()
}
