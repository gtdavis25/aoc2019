package main

type cell struct {
	x, y, level int
}

func (c cell) adjacentCells() []cell {
	var adjacent []cell
	for _, p := range []struct{ x, y int }{
		{c.x, c.y - 1},
		{c.x - 1, c.y},
		{c.x + 1, c.y},
		{c.x, c.y + 1},
	} {
		if p.x < 0 {
			adjacent = append(adjacent, cell{1, 2, c.level + 1})
		} else if p.x > 4 {
			adjacent = append(adjacent, cell{3, 2, c.level + 1})
		} else if p.y < 0 {
			adjacent = append(adjacent, cell{2, 1, c.level + 1})
		} else if p.y > 4 {
			adjacent = append(adjacent, cell{2, 3, c.level + 1})
		} else if p.x == 2 && p.y == 2 {
			if c.x == 2 && c.y == 1 {
				for x := 0; x < 5; x++ {
					adjacent = append(adjacent, cell{x, 0, c.level - 1})
				}
			} else if c.x == 1 && c.y == 2 {
				for y := 0; y < 5; y++ {
					adjacent = append(adjacent, cell{0, y, c.level - 1})
				}
			} else if c.x == 3 && c.y == 2 {
				for y := 0; y < 5; y++ {
					adjacent = append(adjacent, cell{4, y, c.level - 1})
				}
			} else if c.x == 2 && c.y == 3 {
				for x := 0; x < 5; x++ {
					adjacent = append(adjacent, cell{x, 4, c.level - 1})
				}
			} else {
				panic("this shouldn't happen")
			}
		} else {
			adjacent = append(adjacent, cell{p.x, p.y, c.level})
		}
	}

	return adjacent
}
