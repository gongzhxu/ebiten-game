package mobile

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gongzhxu/ebiten-game/gamelib"
	"github.com/hajimehoshi/ebiten/v2/mobile"
)

func init() {
	game, err := gamelib.NewGame("cgotest")
	if err != nil {
		g.Log().Panic(err)
	}

	mobile.SetGame(game)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
