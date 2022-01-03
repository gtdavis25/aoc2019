package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	var asteroids []point
	for y, scanner := 0, bufio.NewScanner(os.Stdin); scanner.Scan(); y++ {
		line := scanner.Text()
		for x := range line {
			if line[x] == '#' {
				asteroids = append(asteroids, point{x, y})
			}
		}
	}

	p := asteroids[0]
	max := countVisible(asteroids, p)
	for _, asteroid := range asteroids[1:] {
		if count := countVisible(asteroids, asteroid); count > max {
			p = asteroid
			max = count
		}
	}

	fmt.Println("Part 1:", max)
	rays := make(map[float64][]point)
	for _, to := range asteroids {
		if p == to {
			continue
		}

		dx, dy := to.x-p.x, to.y-p.y
		angle := math.Atan2(float64(dx), float64(-dy))
		if angle < 0 {
			angle += math.Pi * 2
		}

		rays[angle] = append(rays[angle], to)
	}

	var angles []float64
	for angle := range rays {
		angles = append(angles, angle)
	}

	for _, angle := range angles {
		ray := rays[angle]
		for i := 1; i < len(ray); i++ {
			for j := i; j > 0 && distance(p, ray[j]) < distance(p, ray[j-1]); j-- {
				ray[j-1], ray[j] = ray[j], ray[j-1]
			}
		}
	}

	sort.Float64s(angles)
	var destroyed []point
	for len(destroyed) < 200 {
		for _, angle := range angles {
			if len(rays[angle]) > 0 {
				destroyed = append(destroyed, rays[angle][0])
				rays[angle] = rays[angle][1:]
			}
		}
	}

	fmt.Println("Part 2:", destroyed[199].x*100+destroyed[199].y)
}

func countVisible(asteroids []point, from point) int {
	var count int
	slopes := make(map[point]bool)
	for _, to := range asteroids {
		if from == to {
			continue
		}

		dx, dy := to.x-from.x, to.y-from.y
		gcd := gcd(dx, dy)
		slope := point{dx / gcd, dy / gcd}
		if !slopes[slope] {
			slopes[slope] = true
			count++
		}
	}

	return count
}

func distance(from, to point) int {
	dx, dy := to.x-from.x, to.y-from.y
	return abs(dx) + abs(dy)
}

func gcd(a, b int) int {
	if b == 0 {
		return abs(a)
	}

	return gcd(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
