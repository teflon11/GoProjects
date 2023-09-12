package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	touchId1        []ebiten.TouchID
	xt, yt, dt, len int
}

func (g *Game) Update() error {
	g.touchId1 = inpututil.AppendJustPressedTouchIDs(g.touchId1[:0])
	g.len = len(g.touchId1)
	for _, id := range g.touchId1 {
		g.xt, g.yt = ebiten.TouchPosition(id)
		g.dt = inpututil.TouchPressDuration(id)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("(%-3d - %-3d)\t-\t%-3d", g.xt, g.yt, g.dt))
	ebitenutil.DebugPrintAt(screen, strconv.Itoa(g.len), g.xt, g.yt)
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
