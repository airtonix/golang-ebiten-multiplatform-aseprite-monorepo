package main

import (
	"github.com/airtonix/golang-ebiten-aesprite-mobile/pkg/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type DesktopGame struct {
	*game.Game
}

func NewDesktopGame() *DesktopGame {
	return &DesktopGame{
		game.NewGame(),
	}
}
func (app *DesktopGame) Layout(w, h int) (int, int) { return 320, 180 }
func (app *DesktopGame) Draw(screen *ebiten.Image)  { app.Game.Draw(screen) }
func (app *DesktopGame) Update() error {
	app.Game.Update()
	return nil
}

func main() {
	app := NewDesktopGame()
	ebiten.RunGame(app)

}
