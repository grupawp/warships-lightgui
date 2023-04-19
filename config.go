package board

import "github.com/fatih/color"

/*
Config stores colours and characters used to draw a board.
Zero values result in no colours and no characters, so it
is recommended to use NewConfig() and modify it instead. 
*/
type Config struct {
	EmptyChar      byte
	EmptyColor     color.Attribute
	RulerTextColor color.Attribute
	ShipChar       byte
	ShipColor      color.Attribute
	HitChar        byte
	HitColor       color.Attribute
	MissChar       byte
	MissColor      color.Attribute
	BorderChar     byte
	BorderColor    color.Attribute
}

// NewConfig returns a new config with default values.
func NewConfig() *Config {
	return &Config{
		EmptyChar:      ' ',
		EmptyColor:     color.Reset,
		RulerTextColor: color.BgBlue,
		ShipChar:       'X',
		ShipColor:      color.FgCyan,
		BorderChar:     '.',
		BorderColor:    color.FgRed,
		HitChar:        'X',
		HitColor:       color.FgRed,
		MissChar:       '.',
		MissColor:      color.FgRed,
	}
}
