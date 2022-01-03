package main

type planet struct {
	position, velocity point
}

func (p *planet) potentialEnergy() int {
	return abs(p.position.x) + abs(p.position.y) + abs(p.position.z)
}

func (p *planet) kineticEnergy() int {
	return abs(p.velocity.x) + abs(p.velocity.y) + abs(p.velocity.z)
}

func (p *planet) totalEnergy() int {
	return p.potentialEnergy() * p.kineticEnergy()
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
