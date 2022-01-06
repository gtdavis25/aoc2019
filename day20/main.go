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

	vertices := getVertices(lines)
	buildEdges(lines, vertices)
	start := findVertex(vertices, func(v *vertex) bool { return v.label == "AA" })
	shortest := findShortestPath(&path{to: start}, func(p *path) []*path {
		var nextPaths []*path
		for _, e := range p.to.edges {
			nextPaths = append(nextPaths, &path{
				to:       e.to,
				distance: p.distance + e.length,
			})
		}

		return nextPaths
	})

	fmt.Println("Part 1:", shortest.distance)
	shortest = findShortestPath(&path{to: start}, func(p *path) []*path {
		var nextPaths []*path
		for _, e := range p.to.edges {
			if e.isPortal && p.to.isOuter && p.level == 0 {
				continue
			} else if e.isPortal && p.to.isOuter {
				nextPaths = append(nextPaths, &path{
					to:       e.to,
					distance: p.distance + e.length,
					level:    p.level - 1,
				})
			} else if e.isPortal {
				nextPaths = append(nextPaths, &path{
					to:       e.to,
					distance: p.distance + e.length,
					level:    p.level + 1,
				})
			} else {
				nextPaths = append(nextPaths, &path{
					to:       e.to,
					distance: p.distance + e.length,
					level:    p.level,
				})
			}
		}

		return nextPaths
	})

	fmt.Println("Part 2:", shortest.distance)
}

func findShortestPath(initial *path, transition func(*path) []*path) *path {
	shortestPaths := map[string]*path{initial.key(): initial}
	var q pathQueue
	heap.Push(&q, initial)
	for len(q) > 0 {
		state := heap.Pop(&q).(*path)
		if state.to.label == "ZZ" && state.level == 0 {
			return state
		}

		for _, next := range transition(state) {
			key := next.key()
			if p, ok := shortestPaths[key]; ok && next.distance < p.distance {
				p.distance = next.distance
				heap.Fix(&q, p.index)
			} else if !ok {
				shortestPaths[key] = next
				heap.Push(&q, next)
			}
		}
	}

	return nil
}

func getVertices(lines []string) []*vertex {
	var vertices []*vertex
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] != '.' {
				continue
			}

			var label string
			if isLetter(lines[y-1][x]) {
				label = fmt.Sprintf("%c%c", lines[y-2][x], lines[y-1][x])
			} else if isLetter(lines[y][x-1]) {
				label = fmt.Sprintf("%c%c", lines[y][x-2], lines[y][x-1])
			} else if isLetter(lines[y][x+1]) {
				label = fmt.Sprintf("%c%c", lines[y][x+1], lines[y][x+2])
			} else if isLetter(lines[y+1][x]) {
				label = fmt.Sprintf("%c%c", lines[y+1][x], lines[y+2][x])
			}

			if len(label) > 0 {
				vertices = append(vertices, &vertex{
					position: point{x, y},
					label:    label,
				})
			}
		}
	}

	minX := vertices[0].position.x
	minY := vertices[0].position.y
	maxX := vertices[0].position.y
	maxY := vertices[0].position.y
	for _, v := range vertices[1:] {
		minX = min(minX, v.position.x)
		minY = min(minY, v.position.y)
		maxX = max(maxX, v.position.x)
		maxY = max(maxY, v.position.y)
	}

	for _, v := range vertices {
		if v.position.x == minX || v.position.x == maxX || v.position.y == minY || v.position.y == maxY {
			v.isOuter = true
		}
	}

	return vertices
}

func buildEdges(lines []string, vertices []*vertex) {
	for _, v := range vertices {
		if to := findVertex(vertices, func(to *vertex) bool {
			return v.label == to.label && v != to
		}); to != nil {
			v.edges = append(v.edges, edge{
				to:       to,
				length:   1,
				isPortal: true,
			})
		}

		seen := map[point]bool{v.position: true}
		for q := []step{{v.position, 0}}; len(q) > 0; q = q[1:] {
			current := q[0]
			if to := findVertex(vertices, func(v *vertex) bool {
				return v.position == current.to
			}); to != nil && to != v {
				v.edges = append(v.edges, edge{
					to:     to,
					length: current.distance,
				})

				continue
			}

			for _, p := range []point{
				{current.to.x, current.to.y - 1},
				{current.to.x - 1, current.to.y},
				{current.to.x + 1, current.to.y},
				{current.to.x, current.to.y + 1},
			} {
				if seen[p] || lines[p.y][p.x] != '.' {
					continue
				}

				seen[p] = true
				q = append(q, step{
					to:       p,
					distance: current.distance + 1,
				})
			}
		}
	}
}

func findVertex(vertices []*vertex, condition func(*vertex) bool) *vertex {
	for _, v := range vertices {
		if condition(v) {
			return v
		}
	}

	return nil
}

func isLetter(b byte) bool {
	return 'A' <= b && b <= 'Z'
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
