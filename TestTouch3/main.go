package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	touchId0   []ebiten.TouchID
	xt, yt, dt int
	len0       int
	out        string
}

func (g *Game) Update() error {
	g.out = ""
	g.touchId0 = ebiten.AppendTouchIDs(g.touchId0[:0])
	g.len0 = len(g.touchId0)
	for _, idT := range g.touchId0 {
		g.xt, g.yt = ebiten.TouchPosition(idT)
		g.dt = inpututil.TouchPressDuration(idT)
		g.out += fmt.Sprintf("(%d, %d)\t:\t%d\n", g.xt, g.yt, g.dt)
	}
	g.out += fmt.Sprintf("(len - %d)\n", g.len0)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(g.xt), float32(g.yt), 20, color.RGBA{8, 20, 157, 0}, false)
	ebitenutil.DebugPrint(screen, g.out)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, Touch!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
