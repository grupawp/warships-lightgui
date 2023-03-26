package board

import (
	"fmt"

	"github.com/fatih/color"
)

func (b *Board) Display() {
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
				c := color.New(b.c.emptyColor)
				c.Printf(" %s", b.printChar(s))
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
