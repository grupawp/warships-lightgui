package board

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
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

type board struct {
	b [maxX][maxY]state
	c *config
}

func (b *board) printChar(s state) string {
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

func (b *board) Import(coords []string) {
	for _, c := range coords {
		b.Set(Left, c, Ship)
	}
}

func (b *board) Export(p pos) []string {
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

func (b *board) Display() {
	const (
		clearScreen = "\033[H\033[2J"
	)

	fmt.Print(clearScreen)

	for y := maxY - 1; y >= 0; y-- {
		for x := 0; x < maxX; x++ {
			switch s := b.b[x][y]; s {
			case Ruler:
				c := color.New(b.c.rulerTextColor)
				c.Printf("%2d", y+1)
			case Empty:
				fmt.Printf("  ")
			case Hit:
				c := color.New(b.c.hitColor)
				c.Printf(" %s", b.printChar(s))
			case Ship:
				c := color.New(b.c.shipColor)
				c.Printf(" %s", b.printChar(s))
			case Miss:
				c := color.New(b.c.missColor)
				c.Printf(" %s", b.printChar(s))
			case Border:
				c := color.New(b.c.borderColor)
				c.Printf(" %s", b.printChar(s))
			}
		}
		fmt.Printf("\n")
	}

	d := color.New(b.c.rulerTextColor)
	d.EnableColor()
	d.Printf("   A B C D E F G H I J")
	d.DisableColor()
	for i := 0; i < (delimiter*2)-3; i++ {
		fmt.Printf(" ")
	}
	d.EnableColor()
	d.Printf("    A B C D E F G H I J")
	d.DisableColor()
	fmt.Println()
}

type point struct {
	x, y int
}

func (b *board) drawBorder(p point, pos pos) {
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
		if pos == Left {
			if dx < 0 || dx >= boardWidth+delimiter || dy < 0 || dy >= maxY {
				continue
			}
		}
		if pos == Right {
			if dx < boardWidth+delimiter+1 || dx >= maxX || dy < 0 || dy >= maxY {
				continue
			}
		}

		if b.b[dx][dy] != Ship && b.b[dx][dy] != Hit {
			b.b[dx][dy] = Border
		}
	}
}

func (b *board) searchElement(x, y int, p *[]point, pos pos) {
	vec := []point{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for _, i := range *p {
		if i.x == x && i.y == y {
			return
		}
	}

	*p = append(*p, point{x, y})
	connections := []point{}

	for _, v := range vec {
		dx := x + v.x
		dy := y + v.y
		if pos == Left {
			if dx < 0 || dx >= boardWidth+delimiter || dy < 0 || dy >= maxY {
				continue
			}
		}
		if pos == Right {
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
		b.searchElement(c.x, c.y, p, pos)
	}
}

func (b *board) CreateBorder(coord string, pos pos) {
	x, y := b.stringCoordToInt(coord)

	if pos == Right {
		x = x + boardWidth + delimiter
	}

	p := []point{}
	b.searchElement(x, y, &p, pos)

	for _, i := range p {
		b.drawBorder(i, pos)
	}
}

func (b *board) HitOrMiss(p pos, coord string) state {
	var s state

	x, y := b.stringCoordToInt(coord)

	if p == Left {
		s = b.b[x][y]
	} else {
		s = b.b[x+boardWidth+delimiter][y]
	}

	if s == Ship {
		b.Set(p, coord, Hit)
		return Hit
	} else {
		b.Set(p, coord, Miss)
		return Miss
	}
}

func (b *board) stringCoordToInt(coord string) (int, int) {
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

func (b *board) Set(p pos, coord string, s state) {
	x, y := b.stringCoordToInt(coord)

	if p == Left {
		b.b[x][y] = s
	} else {
		b.b[x+boardWidth+delimiter][y] = s
	}
}

func New(c *config) *board {
	b := &board{
		c: c,
	}

	for y := maxY - 1; y >= 0; y-- {
		b.b[0][y] = Ruler
		b.b[boardWidth+delimiter][y] = Ruler
	}

	return b
}
