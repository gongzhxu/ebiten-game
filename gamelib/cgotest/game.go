package cgotest

// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG
// #include <stdlib.h>
// #include <stdint.h>
// #include "cgotest.h"
import "C"
import (
	gframe "github.com/gogf/gf/frame/g"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// Settings
	screenWidth  = 960
	screenHeight = 540
)

type Game struct {
}

func NewGame() (*Game, error) {
	gframe.Log().Infof("new game")
	return &Game{}, nil
}

func (g *Game) Update() error {
	a := 123
	b := 456

	ret := C.cgo_add(C.int(a), C.int(b))
	gframe.Log().Infof("update ret: %v", ret)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
