package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gongzhxu/ebiten-game/gamelib"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := gamelib.NewGame("cgotest")
	if err != nil {
		g.Log().Panic(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		g.Log().Panic(err)
	}
}
