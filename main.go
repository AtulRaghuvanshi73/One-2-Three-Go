package main

import (
	"image/color"
	"log"
	"math"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type GameState int

const (
	StateWelcomeScreen GameState = iota
	StateCharacterSelection
	StatePlaying
)

type Game struct {
	State          GameState
	PlayerX        float64
	PlayerY        float64
	FadeOpacity    float64
	FadingIn       bool
	FadeSpeed      float64
	TitleText      string
	SelectedIndex  int
	Characters     []string
	PulseStartTime time.Time
}

func (g *Game) Update() error {
	switch g.State {
	case StateWelcomeScreen:
		// Handle fading in/out logic for the welcome screen
		if g.FadingIn {
			g.FadeOpacity += g.FadeSpeed
			if g.FadeOpacity >= 1.0 {
				g.FadeOpacity = 1.0
				g.FadingIn = false
			}
		} else {
			g.FadeOpacity -= g.FadeSpeed
			if g.FadeOpacity <= 0.0 {
				g.FadeOpacity = 0.0
				g.State = StateCharacterSelection
				g.PulseStartTime = time.Now()
			}
		}
	case StateCharacterSelection:
		// Handle character selection
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			if g.SelectedIndex > 0 {
				g.SelectedIndex--
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			if g.SelectedIndex < len(g.Characters)-1 {
				g.SelectedIndex++
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.State = StatePlaying
		}
	case StatePlaying:
		// Gameplay logic goes here
		// ...
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	switch g.State {
	case StateWelcomeScreen:
		// Draw the title text
		textWidth := float64(text.BoundString(basicfont.Face7x13, g.TitleText).Max.X)
		titleX := float64(screen.Bounds().Dx())/2 - textWidth/2
		titleY := float64(screen.Bounds().Dy()) / 3
		text.Draw(screen, g.TitleText, basicfont.Face7x13, int(titleX), int(titleY), color.RGBA{255, 255, 0, uint8(255 * g.FadeOpacity)})

	case StateCharacterSelection:
		// Draw a semi-transparent background overlay
		screen.Fill(color.RGBA{0, 0, 0, 100})

		// Draw the character selection message
		message := "Choose Your Character"
		textWidth := float64(text.BoundString(basicfont.Face7x13, message).Max.X)
		titleX := float64(screen.Bounds().Dx())/2 - textWidth/2
		titleY := float64(screen.Bounds().Dy()) / 5
		text.Draw(screen, message, basicfont.Face7x13, int(titleX), int(titleY), color.White)

		// Draw characters with gradient effect
		for i, char := range g.Characters {
			charX := float64(screen.Bounds().Dx())/2 + float64((i-1)*150)
			charY := float64(screen.Bounds().Dy()) / 2
			charColor := color.RGBA{255, 255, 255, 200}
			if i == g.SelectedIndex {
				// Pulsing effect for the selected character
				elapsed := time.Since(g.PulseStartTime)
				pulseOpacity := 200 + 55*float64(math.Sin(elapsed.Seconds()*2*math.Pi))
				charColor = color.RGBA{255, 255, 255, uint8(pulseOpacity)}
			}
			text.Draw(screen, char, basicfont.Face7x13, int(charX), int(charY), charColor)
		}

	case StatePlaying:
		// Draw gameplay screen
		text.Draw(screen, "Game Started!", basicfont.Face7x13, 20, 20, color.Black)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("One-2-Three-Go!")

	game := &Game{
		State:          StateWelcomeScreen,
		FadeOpacity:    0.0,
		FadingIn:       true,
		FadeSpeed:      0.01,
		TitleText:      strings.ToUpper("One-2-Three-Go!"),
		Characters:     []string{"Character 1", "Character 2"},
		SelectedIndex:  0,
		PulseStartTime: time.Now(),
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
