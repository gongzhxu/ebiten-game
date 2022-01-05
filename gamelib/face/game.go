package face

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"path/filepath"

	gframe "github.com/gogf/gf/frame/g"
	faceres "github.com/gongzhxu/ebiten-game/gamelib/internal/face"
	goface "github.com/gongzhxu/go-face"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// Settings
	screenWidth  = 960
	screenHeight = 540
)

const dataDir = "data"

var (
	modelsDir = filepath.Join(dataDir, "models")
	imagesDir = filepath.Join(dataDir, "images")
)

type Game struct {
	rec *goface.Recognizer
}

func NewGame() (*Game, error) {
	gframe.Log().Infof("new game")
	rec, err := goface.NewRecognizer()
	if err != nil {
		return nil, err
	}

	buf, err := faceres.Models.ReadFile("models/shape_predictor_5_face_landmarks.dat")
	if err != nil {
		return nil, err
	}

	if err = rec.Deserialize(0, buf); err != nil {
		return nil, err
	}

	buf, err = faceres.Models.ReadFile("models/dlib_face_recognition_resnet_model_v1.dat")
	if err != nil {
		return nil, err
	}

	if err = rec.Deserialize(1, buf); err != nil {
		return nil, err
	}

	buf, err = faceres.Models.ReadFile("models/mmod_human_face_detector.dat")
	if err != nil {
		return nil, err
	}

	if err = rec.Deserialize(2, buf); err != nil {
		return nil, err
	}

	return &Game{
		rec: rec,
	}, nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	imgData, err := faceres.Images.ReadFile("images/bona.jpg")
	if err != nil {
		gframe.Log().Errorf("read file err: %v", err)
		return
	}

	faces, err := g.rec.Recognize(imgData)
	if err != nil {
		gframe.Log().Errorf("recognize err: %v", err)
		return
	}

	gframe.Log().Debugf("rec faces: %+v", faces)
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		gframe.Log().Errorf("img decode err: %v", err)
		return
	}

	bonaImage := ebiten.NewImageFromImage(img)
	op := &ebiten.DrawImageOptions{}
	//screen.DrawImage(bonaImage.SubImage(faces[0].Rectangle).(*ebiten.Image), op)
	screen.DrawImage(bonaImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
