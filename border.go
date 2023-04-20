package board

import "log"

type point struct {
	x, y int
}

func (b *Board) drawBorder(p point, position pos) {
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

		prev := b.b[dx][dy]
		if !(prev == Ship || prev == Hit || prev == Miss) { // don't overwrite already marked
			b.b[dx][dy] = border
		}
	}
}

func (b *Board) searchElement(x, y int, points *[]point, p pos) {
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

// CreateBorder creates a border around a (sunken) ship, to indicate
// which coordinates cannot contain a ship segment and can be safely
// ignored.
func (b *Board) CreateBorder(p pos, coord string) {
	x, y, err := b.stringCoordToInt(coord)
	if err != nil {
		log.Println(err)
		return
	}

	if p == Right {
		x = x + boardWidth + delimiter
	}

	points := []point{}
	b.searchElement(x, y, &points, p)

	for _, i := range points {
		b.drawBorder(i, p)
	}
}
