package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct {
	playerX float64
	playerY float64
}

func (g *Game) Update() error {
	// Updated keyboard input handling
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.playerX -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerX += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.playerY -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.playerY += 2
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Your drawing logic goes here
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Set the layout dimensions
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("One-2-Three-Go!")

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
