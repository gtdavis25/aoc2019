package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var initialPositions []point
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		var p point
		fmt.Sscanf(scanner.Text(), "<x=%d, y=%d, z=%d>", &p.x, &p.y, &p.z)
		initialPositions = append(initialPositions, p)
	}

	planets := make([]*planet, len(initialPositions))
	for i := range planets {
		planets[i] = &planet{position: initialPositions[i]}
	}

	for i := 0; i < 1000; i++ {
		nextState(planets)
	}

	var result int
	for _, planet := range planets {
		result += planet.totalEnergy()
	}

	fmt.Println("Part 1:", result)
	initialX := make([]int, len(initialPositions))
	initialY := make([]int, len(initialPositions))
	initialZ := make([]int, len(initialPositions))
	currentX := make([]int, len(initialPositions))
	currentY := make([]int, len(initialPositions))
	currentZ := make([]int, len(initialPositions))
	for i := range initialPositions {
		initialX[i] = initialPositions[i].x
		initialY[i] = initialPositions[i].y
		initialZ[i] = initialPositions[i].z
		planets[i] = &planet{position: initialPositions[i]}
	}

	var xPeriod, yPeriod, zPeriod int
	for t := 2; xPeriod == 0 || yPeriod == 0 || zPeriod == 0; t++ {
		nextState(planets)
		for i := range planets {
			currentX[i] = planets[i].position.x
			currentY[i] = planets[i].position.y
			currentZ[i] = planets[i].position.z
		}

		if xPeriod == 0 && equals(initialX, currentX) {
			xPeriod = t
		}

		if yPeriod == 0 && equals(initialY, currentY) {
			yPeriod = t
		}

		if zPeriod == 0 && equals(initialZ, currentZ) {
			zPeriod = t
		}
	}

	fmt.Println("Part 2:", lcm(xPeriod, lcm(yPeriod, zPeriod)))
}

func nextState(planets []*planet) {
	for i := range planets[:len(planets)-1] {
		for j := i + 1; j < len(planets); j++ {
			d := point{
				x: sign(planets[j].position.x - planets[i].position.x),
				y: sign(planets[j].position.y - planets[i].position.y),
				z: sign(planets[j].position.z - planets[i].position.z),
			}

			planets[i].velocity = planets[i].velocity.add(d)
			planets[j].velocity = planets[j].velocity.minus(d)
		}
	}

	for _, planet := range planets {
		planet.position = planet.position.add(planet.velocity)
	}
}

func sign(a int) int {
	if a < 0 {
		return -1
	} else if a == 0 {
		return 0
	} else {
		return 1
	}
}

func equals(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}

	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}

	return true
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func gcd(a, b int) int {
	if b == 0 {
		return abs(a)
	}

	return gcd(b, a%b)
}
