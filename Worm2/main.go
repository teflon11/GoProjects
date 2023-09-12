package main

import (
	"errors"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	title     string = "Worm2"
	scrWidth  int    = 640
	scrHeight int    = 480
	wField    int    = 40
	hField    int    = 30
	maxLen    rune   = 300
	dpi              = 72
)

var (
	//clGreen = color.RGBA{0x75, 0xf9, 0x4d, 0xff}
	//clBlack        = color.RGBA{0x00, 0x00, 0x00, 0xff}
	//clPink    = color.RGBA{0xea, 0x36, 0x80, 0xff}
	pressStartFont font.Face
	fontSize       = 16
	clGreen        = color.RGBA{0x60, 0xa6, 0x65, 0xff}
	clBrown        = color.RGBA{0x78, 0x43, 0x15, 0xff}
	clAqua         = color.RGBA{0x00, 0xFF, 0xff, 0xff}
	clRed          = color.RGBA{0xff, 0x00, 0x00, 0xff}
	clFuchsia      = color.RGBA{0xff, 0x00, 0xff, 0xff}
	errorExit      = errors.New("regular termination")
)

type Pos struct {
	x rune
	y rune
}
type Game struct {
	wormField   [wField][hField]rune
	wormBody    [maxLen]Pos
	wormHead    rune
	wormTail    rune
	wormEat     rune
	delay       rune
	count       rune
	lenght      rune
	direct      rune
	pause       bool
	debug       bool
	printText   string
	returnError error
	currInput   string
	worker      bool
	tittleMode  string
}

func init() {
	tt1, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}

	pressStartFont, err = opentype.NewFace(tt1, &opentype.FaceOptions{
		Size:    float64(fontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Update() error {

	g.currInput = inputWorm()

	switch g.currInput {
	case "EscapeIn":
		{
			if !g.worker {
				if g.pause {
					return errorExit
				}
				g.printText = fmt.Sprintf("Длинна червяка: %d , delay: %d", g.lenght, g.delay)
				g.tittleMode = "Выход"
				g.pause = true
				g.returnError = errorExit
				g.worker = true
			}
		}
	case "RightIn":
		{
			if !g.worker {
				g.directRight()
				g.worker = true
			}
		}
	case "LeftIn":
		{
			if !g.worker {
				g.directLeft()
				g.worker = true
			}
		}
	case "PauseIn":
		{
			if !g.worker {
				if g.returnError == nil {
					g.pause = !g.pause
					g.tittleMode = "Off"
					if g.pause {
						g.printText = fmt.Sprintf("Пауза, длинна %d  delay %d", g.lenght, g.delay)
					}
				}
				g.worker = true
			}
		}
	case "DebugIn":
		{
			if !g.worker {
				g.debug = !g.debug
				g.worker = true
			}
		}
	case "Not Input":
		{
			g.worker = false
		}
	default:
		{

		}
	}
	g.count--
	if g.count > 0 {
		return nil
	}

	g.count = g.delay
	g.stepWorm()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for j := 0; j < hField; j++ {
		for i := 0; i < wField; i++ {
			text.Draw(screen, string(g.wormField[i][j]), pressStartFont, i*fontSize, j*fontSize+fontSize, clGreen)
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(hField*fontSize))
	op.Blend = ebiten.BlendClear
	text.DrawWithOptions(screen, fmt.Sprintf(g.printText), pressStartFont, op)

	op.ColorScale.ScaleWithColor(clBrown)
	op.Blend = ebiten.BlendSourceOver
	text.DrawWithOptions(screen, fmt.Sprintf(g.printText), pressStartFont, op)

	if g.tittleMode == "Старт" {
		text.Draw(screen, "Игра Worm2", pressStartFont, fontSize, 2*fontSize, clAqua)
		text.Draw(screen, "Управление:", pressStartFont, fontSize, 4*fontSize, clAqua)
		text.Draw(screen, "Left arrow - налево", pressStartFont, fontSize, 6*fontSize, clAqua)
		text.Draw(screen, "Right arrow - направо", pressStartFont, fontSize, 7*fontSize, clAqua)
		text.Draw(screen, "Space - пауза", pressStartFont, fontSize, 8*fontSize, clAqua)
		text.Draw(screen, "Escape - выход", pressStartFont, fontSize, 9*fontSize, clAqua)
		text.Draw(screen, "На телефоне используйте касания", pressStartFont, fontSize, 11*fontSize, clAqua)
		text.Draw(screen, "Слева, Справа, Сверху и Снизу.", pressStartFont, fontSize, 12*fontSize, clAqua)
		text.Draw(screen, "нажмите Space для продолжения", pressStartFont, fontSize, 14*fontSize, clAqua)
		text.Draw(screen, "Или Escape для выхода", pressStartFont, fontSize, 15*fontSize, clAqua)
	}
	if g.tittleMode == "Выход" {
		text.Draw(screen, "=== Игра окончена ===", pressStartFont, fontSize, 2*fontSize, clRed)
		text.Draw(screen, g.printText, pressStartFont, fontSize, 4*fontSize, clFuchsia)
		text.Draw(screen, "нажмите Escape для выхода", pressStartFont, fontSize, 6*fontSize, clRed)
	}

	if g.debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f\nCount: %d\ndirect: %d\nhead: %d\ntail: %d\ncurrInput: %v",
			ebiten.ActualTPS(), ebiten.ActualFPS(), g.count, g.direct, g.wormHead, g.wormTail, g.currInput))
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return scrWidth, scrHeight
}

func main() {
	ebiten.SetWindowSize(scrWidth, scrHeight)
	ebiten.SetWindowTitle(title)
	ggg := &Game{}
	ggg.initWorm()
	if err := ebiten.RunGame(ggg); err != nil {
		log.Fatal(err)
	}

}
