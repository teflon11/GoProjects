package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	touchId0         []ebiten.TouchID
	touchId1         []ebiten.TouchID
	touchId2         []ebiten.TouchID
	xt0, yt0         int
	xt1, yt1         int
	xt2, yt2         int
	len0, len1, len2 int
}

func (g *Game) Update() error {
	g.touchId0 = ebiten.AppendTouchIDs(g.touchId0[:0])
	if len(g.touchId0) > 0 {
		g.xt0, g.yt0 = ebiten.TouchPosition(g.touchId0[0])
		g.len0 = len(g.touchId0)
	} else {
		g.xt0, g.yt0, g.len0 = 0, 0, 0
	}

	g.touchId1 = inpututil.AppendJustPressedTouchIDs(g.touchId1[:0])
	g.len1 = len(g.touchId1)
	for _, id := range g.touchId1 {
		g.xt1, g.yt1 = ebiten.TouchPosition(id)
	}

	g.touchId2 = inpututil.AppendJustReleasedTouchIDs(g.touchId2[:0])
	g.len2 = len(g.touchId2)
	for _, id := range g.touchId2 {
		g.xt2, g.yt2 = inpututil.TouchPositionInPreviousTick(id)

	}

	// if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
	// 	return ebiten.Termination
	// }

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("(%d, %d)\t-\t%d\n(%d, %d)\t-\t%d\n(%d, %d)\t-\t%d",
		g.xt0, g.yt0, g.len0, g.xt1, g.yt1, g.len1, g.xt2, g.yt2, g.len2))
	//vector.DrawFilledCircle(screen, float32(g.xt1), float32(g.yt1), 20, color.RGBA{8, 20, 157, 0}, false)

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
