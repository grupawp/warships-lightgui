package board

import "github.com/fatih/color"

type customizer interface {
	EmptyColor(color.Attribute) customizer
	Empty(byte) customizer
	RulerTextColor(color.Attribute) customizer
	Hit(byte) customizer
	HitColor(color.Attribute) customizer
	Miss(byte) customizer
	MissColor(color.Attribute) customizer
	Ship(byte) customizer
	ShipColor(color.Attribute) customizer
	Border(byte) customizer
	BorderColor(color.Attribute) customizer
	Commit() *config
}

type config struct {
	emptyChar      byte
	emptyColor     color.Attribute
	rulerTextColor color.Attribute
	shipChar       byte
	shipColor      color.Attribute
	hitChar        byte
	hitColor       color.Attribute
	missChar       byte
	missColor      color.Attribute
	borderChar     byte
	borderColor    color.Attribute
}

func (c *config) Commit() *config {
	return c
}

func (c *config) EmptyColor(col color.Attribute) customizer {
	c.emptyColor = col
	return c
}

func (c *config) ShipColor(col color.Attribute) customizer {
	c.shipColor = col
	return c
}

func (c *config) Ship(n byte) customizer {
	c.shipChar = n
	return c
}

func (c *config) Empty(n byte) customizer {
	c.emptyChar = n
	return c
}

func (c *config) Border(n byte) customizer {
	c.borderChar = n
	return c
}

func (c *config) BorderColor(col color.Attribute) customizer {
	c.borderColor = col
	return c
}

func (c *config) Hit(n byte) customizer {
	c.hitChar = n
	return c
}

func (c *config) HitColor(col color.Attribute) customizer {
	c.hitColor = col
	return c
}

func (c *config) Miss(n byte) customizer {
	c.missChar = n
	return c
}

func (c *config) MissColor(col color.Attribute) customizer {
	c.missColor = col
	return c
}

func (c *config) RulerTextColor(col color.Attribute) customizer {
	c.rulerTextColor = col
	return c
}

func NewConfig() customizer {
	return &config{
		emptyChar:      ' ',
		emptyColor:     color.BgBlack,
		rulerTextColor: color.BgBlue,
		shipChar:       'X',
		shipColor:      color.FgCyan,
		borderChar:     ' ',
		borderColor:    color.BgRed,
		hitChar:        'X',
		hitColor:       color.FgRed,
		missChar:       '.',
		missColor:      color.FgGreen,
	}
}
