package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

type Game struct {
	// Add any fields your Game struct needs here
}

// Correct method signature
func (g *Game) Update(screen *ebiten.Image) error {
	// Your update logic goes here
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Your drawing logic goes here
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Set the layout dimensions
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("One-2-Three-Go!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
