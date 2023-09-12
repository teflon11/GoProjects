package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	widthG  int = 320
	heightG int = 240
)

//	type pos struct {
//		x int
//		y int
//	}
type Game struct {
	touchID      []ebiten.TouchID
	xT0, yT0, l0 int
	xT1, yT1, l1 int
	xT2, yT2, l2 int
	xT3, yT3, l3 int
	xT4, yT4, l4 int
	xT5, yT5, l5 int
	//idx1, idx2 int
}

func (g *Game) Update() error {

	g.touchID = inpututil.AppendJustPressedTouchIDs(g.touchID)
	g.l0 = len(g.touchID)
	for _, id := range g.touchID {
		g.xT0, g.yT0 = ebiten.TouchPosition(id)
	}

	g.touchID = inpututil.AppendJustPressedTouchIDs(g.touchID[:0])
	g.l1 = len(g.touchID)
	for _, id := range g.touchID {
		g.xT1, g.yT1 = ebiten.TouchPosition(id)
	}

	g.touchID = inpututil.AppendJustReleasedTouchIDs(g.touchID)
	g.l2 = len(g.touchID)
	for _, id := range g.touchID {
		g.xT2, g.yT2 = inpututil.TouchPositionInPreviousTick(id)
	}

	g.touchID = inpututil.AppendJustReleasedTouchIDs(g.touchID[:0])
	g.l3 = len(g.touchID)
	for _, id := range g.touchID {
		g.xT3, g.yT3 = inpututil.TouchPositionInPreviousTick(id)
	}

	g.touchID = ebiten.AppendTouchIDs(g.touchID)
	g.l4 = len(g.touchID)
	for _, id := range g.touchID {
		g.xT4, g.yT4 = ebiten.TouchPosition(id)
	}

	g.touchID = ebiten.AppendTouchIDs(g.touchID[:0])
	g.l5 = len(g.touchID)
	for _, id := range g.touchID {
		g.xT5, g.yT5 = ebiten.TouchPosition(id)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("(%d, %d) - %d\n(%d, %d) - %d\n(%d, %d) - %d\n(%d, %d) - %d\n(%d, %d) - %d\n(%d, %d) - %d",
		g.xT0, g.yT0, g.l0, g.xT1, g.yT1, g.l1, g.xT2, g.yT2, g.l2, g.xT3, g.yT3, g.l3, g.xT4, g.yT4, g.l4, g.xT5, g.yT5, g.l5))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return widthG, heightG
}

func main() {
	ebiten.SetWindowSize(widthG*2, heightG*2)
	ebiten.SetWindowTitle("Hello, Touch1!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
