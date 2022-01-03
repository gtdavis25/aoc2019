package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	fmt.Fscan(os.Stdin, &input)
	words := strings.Split(input, ",")
	program := make([]int, len(words))
	for i := range program {
		fmt.Sscan(words[i], &program[i])
	}

	droid := newDroid(program)
	caveMap := explore(droid)
	var result int
	breadthFirstSearch(point{0, 0}, caveMap, func(position point, distance int) bool {
		if caveMap[position] == oxygenSystem {
			result = distance
			return false
		}

		return true
	})

	var position point
	fmt.Println("Part 1:", result)
	for p := range caveMap {
		if caveMap[p] == oxygenSystem {
			position = p
			break
		}
	}

	breadthFirstSearch(position, caveMap, func(position point, distance int) bool {
		if distance > result {
			result = distance
		}

		return true
	})

	fmt.Println("Part 2:", result)
}

func explore(droid *repairDroid) map[point]tile {
	caveMap := map[point]tile{{0, 0}: start}
	var position point
	var moves []direction

outer:
	for {
		for _, direction := range []direction{north, south, east, west} {
			dest := position.adjacent(direction)
			if _, seen := caveMap[dest]; seen {
				continue
			}

			switch result := droid.move(direction); result {
			case moved, found:
				caveMap[dest] = empty
				if result == found {
					caveMap[dest] = oxygenSystem
				}

				position = dest
				moves = append(moves, direction)
				continue outer

			case hitWall:
				caveMap[dest] = wall

			default:
				panic(fmt.Sprintf("unexpected move result: %d", result))
			}
		}

		if len(moves) == 0 {
			break
		}

		move := moves[len(moves)-1].reverse()
		droid.move(move)
		position = position.adjacent(move)
		moves = moves[:len(moves)-1]
	}

	return caveMap
}

func breadthFirstSearch(start point, caveMap map[point]tile, continueFunc func(position point, distance int) bool) {
	seen := map[point]bool{start: true}
	queue := []struct {
		position point
		distance int
	}{{start, 0}}
	for ; len(queue) > 0; queue = queue[1:] {
		current := queue[0]
		if !continueFunc(current.position, current.distance) {
			return
		}

		for _, direction := range []direction{north, south, east, west} {
			dest := current.position.adjacent(direction)
			if seen[dest] || caveMap[dest] == wall {
				continue
			}

			seen[dest] = true
			queue = append(queue, struct {
				position point
				distance int
			}{dest, current.distance + 1})
		}
	}
}
