package gamelib

import (
	"fmt"

	"github.com/gongzhxu/ebiten-game/gamelib/blocks"
	"github.com/gongzhxu/ebiten-game/gamelib/carotidartillery"
	"github.com/gongzhxu/ebiten-game/gamelib/flappy"
	"github.com/gongzhxu/ebiten-game/gamelib/mascot"
	"github.com/gongzhxu/ebiten-game/gamelib/paint"
	"github.com/gongzhxu/ebiten-game/gamelib/platformer"
	"github.com/gongzhxu/ebiten-game/gamelib/twenty48"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewGame(gameTypes ...string) (game ebiten.Game, err error) {
	gameType := "flappy"
	if len(gameTypes) != 0 {
		gameType = gameTypes[0]
	}

	switch gameType {
	case "twenty48":
		game, err = twenty48.NewGame()
	case "blocks":
		game, err = blocks.NewGame()
	case "flappy":
		game, err = flappy.NewGame()
	case "mascot":
		game, err = mascot.NewGame()
	case "paint":
		game, err = paint.NewGame()
	case "platformer":
		game, err = platformer.NewGame()
	case "carotidartillery":
		game, err = carotidartillery.NewGame()
	default:
		err = fmt.Errorf("unkonwn game type: %s", gameType)
	}

	return
}
