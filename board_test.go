package board

import (
	"fmt"

	"github.com/fatih/color"
)

func ExampleBoard_Import() {
	board := New(NewConfig())
	coords := []string{"A1", "A2", "A3"}
	board.Import(coords)
}

func ExampleBoard_Export() {
	board := New(NewConfig())
	coords := []string{"A1", "A2", "A3"}
	board.Import(coords)
	exported := board.Export(Left)
	fmt.Println(exported)
	// Output: [A3 A2 A1]
}

func ExampleBoard_Set_enemy() {
	board := New(NewConfig())
	board.Set(Right, "C3", Hit)
}

func ExampleBoard_Set_player() {
	board := New(NewConfig())
	board.Set(Left, "A1", Ship)
}

func ExampleBoard_HitOrMiss() {
	board := New(NewConfig())
	board.Set(Left, "A1", Ship)
	board.HitOrMiss(Left, "A1")
}

func ExampleNew_simple() {
	cfg := NewConfig()
	New(cfg)
}

func ExampleNew_advanced() {
	cfg := NewConfig()
	cfg.HitChar = '#'
	cfg.HitColor = color.FgRed
	cfg.BorderColor = color.BgRed
	cfg.RulerTextColor = color.BgYellow
	New(cfg)
}
