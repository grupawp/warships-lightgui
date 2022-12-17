package board

import "github.com/fatih/color"

type customizer interface {
	EmptyColor(color.Attribute) customizer
	EmptyChar(byte) customizer
	RulerTextColor(color.Attribute) customizer
	HitChar(byte) customizer
	HitColor(color.Attribute) customizer
	MissChar(byte) customizer
	MissColor(color.Attribute) customizer
	ShipChar(byte) customizer
	ShipColor(color.Attribute) customizer
	BorderChar(byte) customizer
	BorderColor(color.Attribute) customizer
	NewConfig() *config
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

func (c *config) NewConfig() *config {
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

func (c *config) ShipChar(n byte) customizer {
	c.shipChar = n
	return c
}

func (c *config) EmptyChar(n byte) customizer {
	c.emptyChar = n
	return c
}

func (c *config) BorderChar(n byte) customizer {
	c.borderChar = n
	return c
}

func (c *config) BorderColor(col color.Attribute) customizer {
	c.borderColor = col
	return c
}

func (c *config) HitChar(n byte) customizer {
	c.hitChar = n
	return c
}

func (c *config) HitColor(col color.Attribute) customizer {
	c.hitColor = col
	return c
}

func (c *config) MissChar(n byte) customizer {
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

func NewConfig() *config {
	return &config{
		emptyChar:      ' ',
		emptyColor:     color.BgBlack,
		rulerTextColor: color.BgBlue,
		shipChar:       'X',
		shipColor:      color.FgCyan,
		borderChar:     'x',
		borderColor:    color.FgRed,
		hitChar:        'X',
		hitColor:       color.FgRed,
		missChar:       '.',
		missColor:      color.FgRed,
	}
}

func ConfigParams() customizer {
	return NewConfig()
}
