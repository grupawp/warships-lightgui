package board

type point struct {
	x, y int
}

func (b *board) drawBorder(p point, position pos) {
	//
	//    XXX
	//    XOX
	//    XXX
	//
	vec := []point{
		{-1, 0},
		{-1, -1},
		{0, 1},
		{1, 1},
		{1, 0},
		{-1, 1},
		{0, -1},
		{1, -1},
	}

	for _, v := range vec {
		dx := p.x + v.x
		dy := p.y + v.y
		if position == Left {
			if dx < 0 || dx >= boardWidth+delimiter || dy < 0 || dy >= maxY {
				continue
			}
		}
		if position == Right {
			if dx < boardWidth+delimiter+1 || dx >= maxX || dy < 0 || dy >= maxY {
				continue
			}
		}

		if b.b[dx][dy] != Ship && b.b[dx][dy] != Hit {
			b.b[dx][dy] = Border
		}
	}
}

func (b *board) searchElement(x, y int, points *[]point, p pos) {
	vec := []point{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for _, i := range *points {
		if i.x == x && i.y == y {
			return
		}
	}

	*points = append(*points, point{x, y})
	connections := []point{}

	for _, v := range vec {
		dx := x + v.x
		dy := y + v.y
		if p == Left {
			if dx < 0 || dx >= boardWidth+delimiter || dy < 0 || dy >= maxY {
				continue
			}
		}
		if p == Right {
			if dx < boardWidth+delimiter || dx >= maxX || dy < 0 || dy >= maxY {
				continue
			}
		}
		if b.b[x+v.x][y+v.y] == Ship || b.b[x+v.x][y+v.y] == Hit {
			connections = append(connections, point{dx, dy})
		}
	}

	// Run the method recursively on each linked element
	for _, c := range connections {
		b.searchElement(c.x, c.y, points, p)
	}
}

func (b *board) CreateBorder(coord string, p pos) {
	x, y := b.stringCoordToInt(coord)

	if p == Right {
		x = x + boardWidth + delimiter
	}

	points := []point{}
	b.searchElement(x, y, &points, p)

	for _, i := range points {
		b.drawBorder(i, p)
	}
}
