package board

import (
	"fmt"

	"github.com/fatih/color"
)

func ExampleBoard_Import() {
	board := New(NewConfig())
	coords := []string{"A1", "A2", "A3"}
	err := board.Import(coords)
	if err != nil {
		fmt.Println(err)
	}
}

func ExampleBoard_Export() {
	board := New(NewConfig())
	coords := []string{"A1", "A2", "A3"}
	_ = board.Import(coords)
	exported := board.Export(Left)
	fmt.Println(exported)
	// Output: [A3 A2 A1]
}

func ExampleBoard_Set_enemy() {
	board := New(NewConfig())
	err := board.Set(Right, "C3", Hit)
	if err != nil {
		fmt.Println(err)
	}
}

func ExampleBoard_Set_player() {
	board := New(NewConfig())
	err := board.Set(Left, "A1", Ship)
	if err != nil {
		fmt.Println(err)
	}
}

func ExampleBoard_HitOrMiss() {
	board := New(NewConfig())
	_ = board.Set(Left, "A1", Ship)
	_, err := board.HitOrMiss(Left, "A1")
	if err != nil {
		fmt.Println(err)
	}
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
