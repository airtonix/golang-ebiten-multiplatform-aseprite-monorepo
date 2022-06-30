package game

import (
	"fmt"
	"image"

	"github.com/airtonix/golang-ebiten-aesprite-mobile/pkg/game/core"
	"github.com/hajimehoshi/ebiten/v2"

	_ "image/png"
)

type Game struct {
	Figure core.Aesprite
}

func NewGame() *Game {
	core.LoadAesprites()
	fmt.Println("Loading figure")
	sprite := core.GetSprite("16x16Deliveryman")
	fmt.Println("Figure Loaded")

	game := &Game{
		Figure: sprite,
	}

	game.Figure.Sprite.Play("idle")

	return game

}

func (game *Game) Update() error {

	game.Figure.Sprite.Update(float32(1.0 / 60.0))

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {

	opts := &ebiten.DrawImageOptions{}

	sub := game.Figure.Img.SubImage(image.Rect(game.Figure.Sprite.CurrentFrameCoords()))

	screen.DrawImage(sub.(*ebiten.Image), opts)

}
func (app *Game) Layout(w, h int) (int, int) { return 320, 180 }
