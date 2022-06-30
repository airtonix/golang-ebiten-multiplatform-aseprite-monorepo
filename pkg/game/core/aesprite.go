package core

import (
	"fmt"
	"image"

	"github.com/airtonix/golang-ebiten-aseprite-mobile/pkg/game/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/goaseprite"
)

type Aseprite struct {
	definitionFilepath string
	Sprite             *goaseprite.File
	Img                *ebiten.Image
}

func NewAseprite(definition string) Aseprite {
	aseprite := Aseprite{
		definitionFilepath: definition,
	}
	aseprite.Load()
	return aseprite
}

func (this *Aseprite) Load() error {
	fmt.Println("Loading Aseprite", this.definitionFilepath)
	// Read the json definition file
	jsonFile, err := assets.Assets.ReadFile(this.definitionFilepath)
	if err != nil {
		fmt.Println("💢 Problem reading Aseprite JSON")

		return err
	}

	this.Sprite = goaseprite.Read(jsonFile)

	// Read the bitmap image file
	imageFile, err := assets.Assets.Open(this.Sprite.ImagePath)
	if err != nil {
		fmt.Println("💢 Problem reading Aseprite Image", this.Sprite.ImagePath)
		return err
	}

	decodedImage, _, err := image.Decode(imageFile)
	if err != nil {
		fmt.Println("💢 Problem decoding Aseprite Image")
		return err
	}

	img := ebiten.NewImageFromImage(decodedImage)
	this.Img = img

	return nil
}
