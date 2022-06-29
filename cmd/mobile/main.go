package mobilegame

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/mobile"

	"github.com/airtonix/golang-ebiten-aesprite-mobile/pkg/game"
)

type MobileGame struct {
	*game.Game
}

func NewMobileGame() *MobileGame {
	return &MobileGame{
		game.NewGame(),
	}
}
func (app *MobileGame) Draw(screen *ebiten.Image) { app.Game.Draw(screen) }
func (app *MobileGame) Update() error {
	app.Game.Update()
	return nil
}

func init() {
	app := NewMobileGame()
	// yourgame.Game must implement ebiten.Game interface.
	// For more details, see
	// * https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Game
	mobile.SetGame(app)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
