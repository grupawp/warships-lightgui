package board

import (
	"fmt"
	"strconv"
	"strings"
)

type pos int

const (
	Left pos = iota
	Right
)

type state int

const (
	Empty state = iota
	Hit
	Miss
	Ship
	Border
	Ruler
)

const (
	delimiter   = 5
	boardWidth  = 10
	boardHeight = 10
	maxX        = boardWidth + delimiter + boardWidth + 1
	maxY        = boardHeight
)

type Board struct {
	b [maxX][maxY]state
	c *config
}

func (b *Board) printChar(s state) string {
	switch s {
	case Empty:
		return string(b.c.emptyChar)
	case Ship:
		return string(b.c.shipChar)
	case Hit:
		return string(b.c.hitChar)
	case Miss:
		return string(b.c.missChar)
	case Border:
		return string(b.c.borderChar)
	case Ruler:
		return ""
	default:
		return string('.')
	}
}

func (b *Board) Import(coords []string) {
	for _, c := range coords {
		b.Set(Left, c, Ship)
	}
}

func (b *Board) Export(p pos) []string {
	var result []string

	if p == Left {
		for y := maxY - 1; y >= 0; y-- {
			for x := 0; x < boardHeight; x++ {
				if b.b[x][y] == Ship {
					coords := fmt.Sprintf("%c%d", 'A'+byte(x-1), y+1)
					result = append(result, coords)
				}
			}
		}
	}

	if p == Right {
		for y := maxY - 1; y >= 0; y-- {
			for x := boardWidth + delimiter; x < maxX; x++ {
				if b.b[x][y] == Ship {
					coords := fmt.Sprintf("%c%d", 'A'+byte(x-boardWidth-delimiter-1), y+1)
					result = append(result, coords)
				}
			}
		}
	}

	return result
}

func (b *Board) HitOrMiss(p pos, coord string) state {
	var s state

	x, y := b.stringCoordToInt(coord)

	if p == Left {
		s = b.b[x][y]
	} else {
		s = b.b[x+boardWidth+delimiter][y]
	}

	switch s {
	case Ship:
		b.Set(p, coord, Hit)
		return Hit
	case Hit:
		return Hit
	default:
		b.Set(p, coord, Miss)
		return Miss
	}
}

func (b *Board) stringCoordToInt(coord string) (int, int) {
	if len(coord) != 2 && len(coord) != 3 {
		return 0, 0
	}

	x := strings.ToUpper(coord)[0] - 'A'

	if x > 10 {
		return 0, 0
	}

	y, err := strconv.Atoi(coord[1:])
	if err != nil {
		return 0, 0
	}

	y--

	if y < 0 || y > maxY-1 {
		return 0, 0
	}

	x++

	return int(x), y
}

func (b *Board) Set(p pos, coord string, s state) {
	x, y := b.stringCoordToInt(coord)

	if p == Left {
		switch s {
		case Miss:
			if b.b[x][y] != Empty {
				return
			}
		case Hit:
			if b.b[x][y] != Ship {
				return
			}
		case Ship:
			if b.b[x][y] != Empty {
				return
			}
		}
	}

	if p == Right {
		x = x + boardWidth + delimiter
		if b.b[x][y] != Empty {
			return
		}
	}

	b.b[x][y] = s
}

func New(c *config) *Board {
	b := &Board{
		c: c,
	}

	for y := maxY - 1; y >= 0; y-- {
		b.b[0][y] = Ruler
		b.b[boardWidth+delimiter][y] = Ruler
	}

	return b
}
