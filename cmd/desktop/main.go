package main

import (
	"github.com/airtonix/golang-ebiten-aesprite-mobile/pkg/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.RunGame(game.NewGame())
}
