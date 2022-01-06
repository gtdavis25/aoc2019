package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	var lines []string
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		lines = append(lines, scanner.Text())
	}

	maze := newMaze(lines)
	fmt.Println("Part 1:", findShortestPath(maze))
	for i, row := range []string{
		"@#@",
		"###",
		"@#@",
	} {
		line := []byte(lines[len(lines)/2-1+i])
		copy(line[len(line)/2-1:], []byte(row))
		lines[len(lines)/2-1+i] = string(line)
	}

	maze = newMaze(lines)
	fmt.Println("Part 2:", findShortestPath(maze))
}

func findShortestPath(m maze) int {
	g := newGraph(m)
	var queue stateQueue
	initial := state{positions: g.getStart()}
	seen := make(map[string]bool)
	heap.Push(&queue, initial)
	for len(queue) > 0 {
		current := heap.Pop(&queue).(state)
		if current.keys.containsAll(m.keys) {
			return current.distance
		}

		key := current.String()
		if seen[key] {
			continue
		}

		seen[key] = true
		for _, next := range current.nextStates() {
			heap.Push(&queue, next)
		}
	}

	panic("no path")
}
