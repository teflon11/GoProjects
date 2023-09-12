package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	xTouch, yTouch int
	lenInp         int
	touchesInp     []ebiten.TouchID
)

func inputWorm() string {

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		return "RightIn"
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		return "LeftIn"
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return "EscapeIn"
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		return "PauseIn"
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		return "DebugIn"
	}

	touchesInp = ebiten.AppendTouchIDs(touchesInp[:0])
	lenInp = len(touchesInp)

	if lenInp == 0 {
		return "Not Input"
	} else {
		xTouch, yTouch = ebiten.TouchPosition(touchesInp[len(touchesInp)-1])
		if xTouch < scrWidth/3 {
			return "LeftIn"
		}
		if xTouch > scrWidth/3*2 {
			return "RightIn"
		}
		if yTouch < scrHeight/3 {
			return "EscapeIn" // Up
		}
		if yTouch > scrHeight/3*2 {
			return "PauseIn" // Down
		}
		return "DebugIn" //	Center
	}
}
