package blocks

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type touchState int

const (
	touchStateNone touchState = iota
	touchStatePressing
	touchStateSettled
	touchStateInvalid
)

// Dir represents a direction.
type touchDir int

const (
	DirUp touchDir = iota
	DirRight
	DirDown
	DirLeft
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func vecToDir(dx, dy int) (touchDir, bool) {
	if abs(dx) < 4 && abs(dy) < 4 {
		return 0, false
	}
	if abs(dx) < abs(dy) {
		if dy < 0 {
			return DirUp, true
		}
		return DirDown, true
	}
	if dx < 0 {
		return DirLeft, true
	}
	return DirRight, true
}

// Input manages the input state including gamepads and keyboards.
type Input struct {
	gamepadIDs                 []ebiten.GamepadID
	virtualGamepadButtonStates map[virtualGamepadButton]int
	gamepadConfig              gamepadConfig

	touches       []ebiten.TouchID
	touchState    touchState
	touchID       ebiten.TouchID
	touchInitPosX int
	touchInitPosY int
	touchLastPosX int
	touchLastPosY int
	touchDir      touchDir
}

// GamepadIDButtonPressed returns a gamepad ID where at least one button is pressed.
// If no button is pressed, GamepadIDButtonPressed returns -1.
func (i *Input) GamepadIDButtonPressed() ebiten.GamepadID {
	i.gamepadIDs = ebiten.AppendGamepadIDs(i.gamepadIDs[:0])
	for _, id := range i.gamepadIDs {
		for b := ebiten.GamepadButton(0); b <= ebiten.GamepadButtonMax; b++ {
			if ebiten.IsGamepadButtonPressed(id, b) {
				return id
			}
		}
	}
	return -1
}

func (i *Input) stateForVirtualGamepadButton(b virtualGamepadButton) int {
	if i.virtualGamepadButtonStates == nil {
		return 0
	}
	return i.virtualGamepadButtonStates[b]
}

func (i *Input) Update() {
	if !i.gamepadConfig.IsGamepadIDInitialized() {
		return
	}

	if i.virtualGamepadButtonStates == nil {
		i.virtualGamepadButtonStates = map[virtualGamepadButton]int{}
	}
	for _, b := range virtualGamepadButtons {
		if !i.gamepadConfig.IsButtonPressed(b) {
			i.virtualGamepadButtonStates[b] = 0
			continue
		}
		i.virtualGamepadButtonStates[b]++
	}

	i.touches = ebiten.AppendTouchIDs(i.touches[:0])
	switch i.touchState {
	case touchStateNone:
		if len(i.touches) == 1 {
			i.touchID = i.touches[0]
			x, y := ebiten.TouchPosition(i.touches[0])
			i.touchInitPosX = x
			i.touchInitPosY = y
			i.touchLastPosX = x
			i.touchLastPosX = y
			i.touchState = touchStatePressing
		}
	case touchStatePressing:
		if len(i.touches) >= 2 {
			break
		}
		if len(i.touches) == 1 {
			if i.touches[0] != i.touchID {
				i.touchState = touchStateInvalid
			} else {
				x, y := ebiten.TouchPosition(i.touches[0])
				i.touchLastPosX = x
				i.touchLastPosY = y
			}
			break
		}
		if len(i.touches) == 0 {
			dx := i.touchLastPosX - i.touchInitPosX
			dy := i.touchLastPosY - i.touchInitPosY
			d, ok := vecToDir(dx, dy)
			if !ok {
				i.touchState = touchStateNone
				break
			}
			i.touchDir = d
			i.touchState = touchStateSettled
		}
	case touchStateSettled:
		i.touchState = touchStateNone
	case touchStateInvalid:
		if len(i.touches) == 0 {
			i.touchState = touchStateNone
		}
	}
}

func (i *Input) IsRotateRightJustPressed() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyX) {
		return true
	}

	if i.touchDir == DirUp && i.touchState == touchStateSettled {
		return true
	}

	return i.stateForVirtualGamepadButton(virtualGamepadButtonButtonB) == 1
}

func (i *Input) IsRotateLeftJustPressed() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		return true
	}
	return i.stateForVirtualGamepadButton(virtualGamepadButtonButtonA) == 1
}

func (i *Input) StateForLeft() int {
	if i.touchDir == DirLeft && i.touchState == touchStateSettled {
		return 1
	}

	if v := inpututil.KeyPressDuration(ebiten.KeyArrowLeft); 0 < v {
		return v
	}

	return i.stateForVirtualGamepadButton(virtualGamepadButtonLeft)
}

func (i *Input) StateForRight() int {
	if i.touchDir == DirRight && i.touchState == touchStateSettled {
		return 1
	}

	if v := inpututil.KeyPressDuration(ebiten.KeyArrowRight); 0 < v {
		return v
	}
	return i.stateForVirtualGamepadButton(virtualGamepadButtonRight)
}

func (i *Input) StateForDown() int {
	if i.touchDir == DirDown && i.touchState == touchStateSettled {
		return 1
	}

	if v := inpututil.KeyPressDuration(ebiten.KeyArrowDown); 0 < v {
		return v
	}
	return i.stateForVirtualGamepadButton(virtualGamepadButtonDown)
}
