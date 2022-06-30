package core

import (
	"fmt"
	"image"

	"github.com/airtonix/golang-ebiten-aesprite-mobile/pkg/game/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/goaseprite"
)

type Aesprite struct {
	definitionFilepath string
	Sprite             *goaseprite.File
	Img                *ebiten.Image
}

func NewAesprite(definition string) Aesprite {
	aesprite := Aesprite{
		definitionFilepath: definition,
	}
	aesprite.Load()
	return aesprite
}

func (this *Aesprite) Load() error {
	fmt.Println("Loading Aesprite", this.definitionFilepath)
	// Read the json definition file
	jsonFile, err := assets.Assets.ReadFile(this.definitionFilepath)
	if err != nil {
		fmt.Println("💢 Problem reading Aesprite JSON")

		return err
	}

	this.Sprite = goaseprite.Read(jsonFile)

	// Read the bitmap image file
	imageFile, err := assets.Assets.Open(this.Sprite.ImagePath)
	if err != nil {
		fmt.Println("💢 Problem reading Aesprite Image", this.Sprite.ImagePath)
		return err
	}

	decodedImage, _, err := image.Decode(imageFile)
	if err != nil {
		fmt.Println("💢 Problem decoding Aesprite Image")
		return err
	}

	img := ebiten.NewImageFromImage(decodedImage)
	this.Img = img

	return nil
}
