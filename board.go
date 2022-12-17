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
	maxX        = boardWidth + delimiter + boardWidth
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
			s := b.b[x][y]
			if s == Ruler {
				c := color.New(b.c.rulerTextColor)
				c.Printf("%2d", y+1)
			}
			if s == Empty {
				fmt.Printf("  ")
			}
			if s == Hit {
				c := color.New(b.c.hitColor)
				c.Printf(" %s", b.printChar(s))
			}
			if s == Ship {
				c := color.New(b.c.shipColor)
				c.Printf(" %s", b.printChar(s))
			}
			if s == Miss {
				c := color.New(b.c.missColor)
				c.Printf(" %s", b.printChar(s))
			}
			if s == Border {
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
	for i := 0; i < delimiter; i++ {
		fmt.Printf("  ")
	}
	d.EnableColor()
	d.Printf(" A B C D E F G H I J")
	d.DisableColor()
	fmt.Println()

}

func (b *board) HitOrMiss(p pos, coord string) state {
	x, y := b.stringCoordToInt(coord)

	var s state

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
