package mobile

import (
	"github.com/gongzhxu/ebiten-game/gamelib/blocks"
	"github.com/hajimehoshi/ebiten/v2/mobile"
)

func init() {
	game, err := blocks.NewGame()
	if err != nil {
		panic(err)
	}

	mobile.SetGame(game)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
