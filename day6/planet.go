package main

type planet struct {
	parent *planet
	orbits []*planet
}

func (p *planet) totalOrbits() int {
	var count int
	for _, orbit := range p.orbits {
		count += 1 + orbit.totalOrbits()
	}

	return count
}

func (p *planet) distanceTo(other *planet) int {
	count := 0
	distances := make(map[*planet]int)
	for current := p; current != nil; current = current.parent {
		distances[current] = count
		count++
	}

	count = 0
	for current := other; current != nil; current = current.parent {
		if d, ok := distances[current]; ok {
			return d + count
		}

		count++
	}

	panic("disconnected planets")
}
