package core

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/airtonix/golang-ebiten-aseprite-mobile/pkg/game/assets"
	"golang.org/x/exp/maps"
)

var (
	sprites = map[string]Aseprite{}
)

func LoadAseprites() error {
	const dir = "images"

	fmt.Println(fmt.Sprintf("📂 Loading aseprites from: %s", dir))

	ents, err := assets.Assets.ReadDir(dir)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("📂 Reading: %d", len(ents)))

	for _, ent := range ents {
		name := ent.Name()
		ext := filepath.Ext(name)
		basename := name[:len(name)-len(ext)]

		filepath := path.Join(dir, name)

		if ext != ".json" {
			continue
		}

		sprites[basename] = NewAseprite(filepath)

	}
	fmt.Println(fmt.Sprintf("Stored %d aseprites: %s",
		len(sprites),
		strings.Join(maps.Keys(sprites), ","),
	))
	return nil
}

func GetSprite(name string) Aseprite {
	fmt.Println(fmt.Sprintf("👀 looking for sprite named %s", name))
	sprite, hasSprite := sprites[name]

	if !hasSprite {
		available := maps.Keys(sprites)
		fmt.Println(fmt.Sprintf("💢 Missing sprite named %s", name))
		fmt.Println(fmt.Sprintf("🫱 Available: [%d] %s",
			len(available),
			strings.Join(available, ","),
		))
		panic("🤯 Error")
	}

	return sprite
}
