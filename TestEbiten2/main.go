package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//ebitenutil.DebugPrint(screen, "Hello, World!")
	// ebitenutil.DrawCircle(screen, 50, 50, 20, color.RGBA{128, 20, 1, 0})
	// ebitenutil.DrawCircle(screen, 60, 60, 20, color.RGBA{8, 20, 157, 0})

	vector.DrawFilledCircle(screen, 50, 50, 20, color.RGBA{128, 20, 1, 0}, false)
	vector.DrawFilledCircle(screen, 60, 60, 20, color.RGBA{8, 20, 157, 0}, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
