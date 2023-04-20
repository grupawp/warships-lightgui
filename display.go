package board

import (
	"fmt"

	"github.com/fatih/color"
)

// Display clears the terminal and prints the board.
func (b *Board) Display() {
	const (
		clearScreen = "\033[H\033[2J"
	)

	fmt.Print(clearScreen)

	for y := maxY - 1; y >= 0; y-- {
		for x := 0; x < maxX; x++ {
			switch s := b.b[x][y]; s {
			case ruler:
				c := color.New(b.c.RulerTextColor)
				c.Printf("%2d", y+1)
			case empty:
				c := color.New(b.c.EmptyColor)
				c.Printf(" %s", b.printChar(s))
			case Hit:
				c := color.New(b.c.HitColor)
				c.Printf(" %s", b.printChar(s))
			case Ship:
				c := color.New(b.c.ShipColor)
				c.Printf(" %s", b.printChar(s))
			case Miss:
				c := color.New(b.c.MissColor)
				c.Printf(" %s", b.printChar(s))
			case border:
				c := color.New(b.c.BorderColor)
				c.Printf(" %s", b.printChar(s))
			}
		}
		fmt.Printf("\n")
	}

	d := color.New(b.c.RulerTextColor)
	d.EnableColor()
	d.Printf("   A B C D E F G H I J ")
	d.DisableColor()
	for i := 0; i < (delimiter*2)-3; i++ {
		fmt.Printf(" ")
	}
	d.EnableColor()
	d.Printf("   A B C D E F G H I J ")
	d.DisableColor()
	fmt.Println()
}
