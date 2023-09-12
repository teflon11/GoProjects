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

const (
	scrWidth  = 320
	scrHeight = 240
)

type Game struct {
	touchId0   []ebiten.TouchID
	xt, yt, dt int
	len0       int
	out        string
}

func getDirection(x, y int) (d string) {
	if x == 0 && y == 0 {
		return "Zero"
	}
	if x < scrWidth/3 {
		return "Left"
	}
	if x > scrWidth/3*2 {
		return "Right"
	}
	if y < scrHeight/3 {
		return "Up"
	}
	if y > scrHeight/3*2 {
		return "Down"
	}
	return "Center"
}

func (g *Game) Update() error {
	g.out = ""
	g.touchId0 = ebiten.AppendTouchIDs(g.touchId0[:0])
	g.len0 = len(g.touchId0)

	if g.len0 > 0 {
		g.xt, g.yt = ebiten.TouchPosition(g.touchId0[len(g.touchId0)-1])
		g.dt = inpututil.TouchPressDuration(g.touchId0[len(g.touchId0)-1])
	} else {
		g.xt, g.yt, g.dt = 0, 0, 0
	}

	g.out += fmt.Sprintf("(%d, %d)\t:\t%d\n", g.xt, g.yt, g.dt)
	g.out += fmt.Sprintf("(len - %d)\n", g.len0)

	switch getDirection(g.xt, g.yt) {
	case "Up":
		g.out += "Up"
	case "Right":
		g.out += "Right"
	case "Down":
		g.out += "Down"
	case "Left":
		g.out += "Left"
	case "Center":
		g.out += "Center"
	case "Zero":
		g.out += "By zeros"
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(g.xt), float32(g.yt), 20, color.RGBA{8, 20, 157, 0}, false)
	ebitenutil.DebugPrint(screen, g.out)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return scrWidth, scrHeight
}

func main() {
	ebiten.SetWindowSize(scrWidth*2, scrHeight*2)
	ebiten.SetWindowTitle("Hello, Touch!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
