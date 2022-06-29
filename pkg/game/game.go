package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"

	_ "image/png"
)

type Game struct {
	Sprite *goaseprite.File
	Img    *ebiten.Image
}

func NewGame() *Game {

	game := &Game{
		Sprite: goaseprite.Open("assets/16x16Deliveryman.json"),
	}

	img, _, err := ebitenutil.NewImageFromFile(game.Sprite.ImagePath)
	if err != nil {
		panic(err)
	}

	game.Sprite.Play("idle")

	game.Img = img

	return game

}

func (game *Game) Update() error {

	game.Sprite.Update(float32(1.0 / 60.0))

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {

	opts := &ebiten.DrawImageOptions{}

	sub := game.Img.SubImage(image.Rect(game.Sprite.CurrentFrameCoords()))

	screen.DrawImage(sub.(*ebiten.Image), opts)

}
