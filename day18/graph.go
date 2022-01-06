package main

type graph struct {
	vertices []*vertex
}

func newGraph(m maze) graph {
	vertices := make(map[point]*vertex)
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			p := point{x, y}
			tile := m.getTile(p)
			if tile != '.' && tile != '#' {
				vertices[p] = &vertex{
					position: p,
					tile:     tile,
				}
			}
		}
	}

	for _, from := range vertices {
		seen := map[point]bool{from.position: true}
		for queue := []step{{position: from.position}}; len(queue) > 0; queue = queue[1:] {
			current := queue[0]
			if to := vertices[current.position]; to != from && to != nil {
				from.edges = append(from.edges, edge{
					to:     to,
					length: current.length(),
				})

				continue
			}

			for _, next := range current.position.adjacentPoints() {
				if m.getTile(next) == '#' || seen[next] {
					continue
				}

				seen[next] = true
				queue = append(queue, step{
					position: next,
					previous: &current,
				})
			}
		}
	}

	var g graph
	for _, v := range vertices {
		g.vertices = append(g.vertices, v)
	}

	return g
}

func (g graph) getStart() []*vertex {
	var starts []*vertex
	for _, v := range g.vertices {
		if v.tile == '@' {
			starts = append(starts, v)
		}
	}

	return starts
}
