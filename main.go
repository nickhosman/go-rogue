package main

import (
	_ "image/png"
	"log"
	
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
