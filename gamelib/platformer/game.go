package platformer

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"

	"github.com/gongzhxu/ebiten-game/gamelib/internal/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	rplatformer "github.com/hajimehoshi/ebiten/v2/examples/resources/images/platformer"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	// Settings
	screenWidth  = 960
	screenHeight = 540
)

var (
	leftSprite      *ebiten.Image
	rightSprite     *ebiten.Image
	idleSprite      *ebiten.Image
	backgroundImage *ebiten.Image
)

func init() {
	// Preload images
	img, _, err := image.Decode(bytes.NewReader(rplatformer.Right_png))
	if err != nil {
		panic(err)
	}
	rightSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(rplatformer.Left_png))
	if err != nil {
		panic(err)
	}
	leftSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(rplatformer.MainChar_png))
	if err != nil {
		panic(err)
	}
	idleSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(rplatformer.Background_png))
	if err != nil {
		panic(err)
	}
	backgroundImage = ebiten.NewImageFromImage(img)
}

const (
	unit    = 16
	groundY = 380
)

type char struct {
	x  int
	y  int
	vx int
	vy int
}

func (c *char) tryJump() {
	c.vy = -10 * unit
}

func (c *char) update() {
	c.x += c.vx
	c.y += c.vy

	if c.y > groundY*unit {
		c.y = groundY * unit
	}

	if c.vx > 0 {
		c.vx -= 4
	} else if c.vx < 0 {
		c.vx += 4
	}

	if c.vy < 20*unit {
		c.vy += 8
	}
}

func (c *char) draw(screen *ebiten.Image) {
	s := idleSprite
	switch {
	case c.vx > 0:
		s = rightSprite
	case c.vx < 0:
		s = leftSprite
	}

	if c.x > screenWidth*unit {
		c.x = 0
	}

	if c.x/unit < -70 {
		c.x = screenWidth * unit
	}

	if c.y < 70 {
		c.y = 70
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(float64(c.x)/unit, float64(c.y)/unit)
	screen.DrawImage(s, op)
}

type Game struct {
	gopher     *char
	touchInput *input.TouchInput
	x16        int
}

func NewGame() (*Game, error) {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Platformer (Ebiten Demo)")
	return &Game{
		touchInput: input.NewTouchInput(),
	}, nil
}

func (g *Game) Update() error {
	if g.gopher == nil {
		g.gopher = &char{x: 50 * unit, y: groundY * unit}
	}

	// Controls
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.gopher.vx = -4 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.gopher.vx = 4 * unit
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.gopher.tryJump()
	}

	g.touchInput.Update()
	if dir, ok := g.touchInput.Dir(); ok {
		if dir == input.DirLeft {
			g.gopher.vx = -4 * unit
		} else if dir == input.DirRight {
			g.gopher.vx = 4 * unit
		} else {
			g.gopher.tryJump()
		}
	}

	g.gopher.update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draws Background Image
	g.x16++
	if g.x16 > screenWidth {
		g.x16 = 0
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)

	op.GeoM.Translate(0, 0)
	screen.DrawImage(backgroundImage.
		SubImage(image.Rect(g.x16*2, 0, screenWidth*2, screenHeight*2)).(*ebiten.Image), op)
	op.GeoM.Translate(float64(screenWidth-g.x16), 0)
	screen.DrawImage(backgroundImage.
		SubImage(image.Rect(0, 0, g.x16*2, screenHeight*2)).(*ebiten.Image), op)

	// Draws the Gopher
	g.gopher.draw(screen)

	// Show the message
	msg := fmt.Sprintf("TPS: %0.2f\nPress the space key to jump.", ebiten.CurrentTPS())
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
