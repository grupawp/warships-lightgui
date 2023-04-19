package board

import "github.com/fatih/color"

type customizer interface {
	NewConfig() *config
	RulerTextColor(color.Attribute) customizer
	BorderChar(byte) customizer
	BorderColor(color.Attribute) customizer
	EmptyChar(byte) customizer
	EmptyColor(color.Attribute) customizer
	HitChar(byte) customizer
	HitColor(color.Attribute) customizer
	MissChar(byte) customizer
	MissColor(color.Attribute) customizer
	ShipChar(byte) customizer
	ShipColor(color.Attribute) customizer
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

// NewConfig returns a new config.
func (c *config) NewConfig() *config {
	return c
}

// RulerTextColor sets the color of the ruler text.
func (c *config) RulerTextColor(col color.Attribute) customizer {
	c.rulerTextColor = col
	return c
}

// BorderChar sets the character used to represent a border around sunken ships.
func (c *config) BorderChar(n byte) customizer {
	c.borderChar = n
	return c
}

// BorderColor sets the color of a border around sunken ships.
func (c *config) BorderColor(col color.Attribute) customizer {
	c.borderColor = col
	return c
}

// EmptyChar sets the character used to represent empty coordinates.
func (c *config) EmptyChar(n byte) customizer {
	c.emptyChar = n
	return c
}

// EmptyColor sets the color of empty coordinates.
func (c *config) EmptyColor(col color.Attribute) customizer {
	c.emptyColor = col
	return c
}

// HitChar sets the character used to represent hits.
func (c *config) HitChar(n byte) customizer {
	c.hitChar = n
	return c
}

// HitColor sets the color of coordinates containing hits.
func (c *config) HitColor(col color.Attribute) customizer {
	c.hitColor = col
	return c
}

// MissChar sets the character used to represent misses.
func (c *config) MissChar(n byte) customizer {
	c.missChar = n
	return c
}

// MissColor sets the color of coordinates containing misses.
func (c *config) MissColor(col color.Attribute) customizer {
	c.missColor = col
	return c
}

// ShipColor sets the color of coordinates containing ship segments.
func (c *config) ShipColor(col color.Attribute) customizer {
	c.shipColor = col
	return c
}

// ShipChar sets the character used to represent ship segments.
func (c *config) ShipChar(n byte) customizer {
	c.shipChar = n
	return c
}

// NewConfig returns a new config with default values.
func NewConfig() *config {
	return &config{
		emptyChar:      ' ',
		emptyColor:     color.Reset,
		rulerTextColor: color.BgBlue,
		shipChar:       'X',
		shipColor:      color.FgCyan,
		borderChar:     '.',
		borderColor:    color.FgRed,
		hitChar:        'X',
		hitColor:       color.FgRed,
		missChar:       '.',
		missColor:      color.FgRed,
	}
}

/*
ConfigParams returns a new config with default values.

Deprecated: use NewConfig instead.
*/
func ConfigParams() customizer {
	return NewConfig()
}
