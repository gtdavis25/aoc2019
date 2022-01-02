package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 2; i++ {
		scanner.Scan()
		lines = append(lines, scanner.Text())
	}

	wires := make([][]point, 2)
	for i := 0; i < 2; i++ {
		wires[i] = plotPath(lines[i])
	}

	var intersections []point
	s1 := make(map[point]bool)
	for _, p := range wires[0] {
		s1[p] = true
	}

	for _, p := range wires[1] {
		if s1[p] {
			intersections = append(intersections, p)
		}
	}

	var min = manhattanDistance(intersections[0])
	for _, p := range intersections[1:] {
		if d := manhattanDistance(p); d < min {
			min = d
		}
	}

	fmt.Println("Part 1:", min)
	min = signalDelay(intersections[0], wires)
	for _, p := range intersections[1:] {
		if d := signalDelay(p, wires); d < min {
			min = d
		}
	}

	fmt.Println("Part 2:", min)
}

func plotPath(line string) []point {
	var path []point
	var position point
	for _, instruction := range strings.Split(line, ",") {
		var direction byte
		var distance int
		fmt.Sscanf(instruction, "%c%d", &direction, &distance)
		var d point
		switch direction {
		case 'U':
			d.y = -1

		case 'L':
			d.x = -1

		case 'R':
			d.x = 1

		case 'D':
			d.y = 1

		default:
			panic(fmt.Sprintf("invalid direction: %c", direction))
		}

		for ; distance > 0; distance-- {
			position = position.add(d)
			path = append(path, position)
		}
	}

	return path
}

func signalDelay(p point, wires [][]point) int {
	return indexOf(p, wires[0]) + indexOf(p, wires[1]) + 2
}

func indexOf(p point, points []point) int {
	for i := range points {
		if p == points[i] {
			return i
		}
	}

	panic("point missing")
}

func manhattanDistance(p point) int {
	return abs(p.x) + abs(p.y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
