//go:build darwin || linux || windows
// +build darwin linux windows

package mobile

import (
	"github.com/airtonix/golang-ebiten-aseprite-mobile/pkg/game"
	"github.com/hajimehoshi/ebiten/v2/mobile"
)

func init() {
	// yourgame.Game must implement ebiten.Game interface.
	// For more details, see
	// * https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Game
	mobile.SetGame(game.NewGame())
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
