package main

import (
	"github.com/gonutz/prototype/draw"
	"github.com/gonutz/sudoku"
	"strconv"
)

func main() {
	var field sudoku.Game
	var cur int
	var blinkTime int
	var cursorVisible bool
	moveCursor := func(delta int) {
		cur += delta
		cursorVisible = true
		blinkTime = 30
	}
	var errMsg string

	draw.RunWindow("Sudoku", 640, 640, func(window draw.Window) {
		if window.WasKeyPressed(draw.KeyEscape) {
			errMsg = ""
			all0 := true
			for i, num := range field {
				if num != 0 {
					all0 = false
				}
				field[i] = 0
			}
			cur = 0
			if all0 {
				window.Close()
			}
		}
		num := -1
		if window.WasKeyPressed(draw.Key1) || window.WasKeyPressed(draw.KeyNum1) {
			num = 1
		}
		if window.WasKeyPressed(draw.Key2) || window.WasKeyPressed(draw.KeyNum2) {
			num = 2
		}
		if window.WasKeyPressed(draw.Key3) || window.WasKeyPressed(draw.KeyNum3) {
			num = 3
		}
		if window.WasKeyPressed(draw.Key4) || window.WasKeyPressed(draw.KeyNum4) {
			num = 4
		}
		if window.WasKeyPressed(draw.Key5) || window.WasKeyPressed(draw.KeyNum5) {
			num = 5
		}
		if window.WasKeyPressed(draw.Key6) || window.WasKeyPressed(draw.KeyNum6) {
			num = 6
		}
		if window.WasKeyPressed(draw.Key7) || window.WasKeyPressed(draw.KeyNum7) {
			num = 7
		}
		if window.WasKeyPressed(draw.Key8) || window.WasKeyPressed(draw.KeyNum8) {
			num = 8
		}
		if window.WasKeyPressed(draw.Key9) || window.WasKeyPressed(draw.KeyNum9) {
			num = 9
		}
		if window.WasKeyPressed(draw.Key0) || window.WasKeyPressed(draw.KeyNum0) ||
			window.WasKeyPressed(draw.KeySpace) || window.WasKeyPressed(draw.KeyBackspace) {
			num = 0
		}
		if window.WasKeyPressed(draw.KeyLeft) {
			moveCursor(-1)
		}
		if window.WasKeyPressed(draw.KeyRight) {
			moveCursor(+1)
		}
		if window.WasKeyPressed(draw.KeyUp) {
			moveCursor(-9)
		}
		if window.WasKeyPressed(draw.KeyDown) {
			moveCursor(+9)
		}
		if window.WasKeyPressed(draw.KeyEnter) || window.WasKeyPressed(draw.KeyNumEnter) {
			errMsg = ""
			solved, err := sudoku.Solve(field)
			if err != nil {
				errMsg = err.Error()
			} else {
				field = solved
			}
		}
		if num != -1 {
			field[cur] = num
			if num != 0 {
				cur++
			}
		}
		for cur < 0 {
			cur += 81
		}
		for cur >= 81 {
			cur -= 81
		}
		blinkTime--
		if blinkTime < 0 {
			cursorVisible = !cursorVisible
			blinkTime = 30
		}
		const (
			tileSize  = 40
			textScale = 3
			ofsX      = (640 - 9*tileSize) / 2
			ofsY      = ofsX / 2
		)
		for y := 0; y < 9; y++ {
			for x := 0; x < 9; x++ {
				color := draw.RGB(0.25, 0.25, 0.25)
				if x%2+y%2 == 1 {
					color = draw.RGB(0.15, 0.15, 0.15)
				}
				if (x/3)%2+(y/3)%2 == 1 {
					color.R += 0.1
				} else {
					color.G += 0.1
				}
				window.FillRect(ofsX+x*tileSize, ofsY+y*tileSize, tileSize, tileSize, color)
				i := x + y*9
				if field[i] != 0 {
					window.DrawScaledText(
						strconv.Itoa(field[i]),
						ofsX+x*tileSize,
						ofsY+y*tileSize,
						textScale,
						draw.White,
					)
				}
			}
		}
		if cursorVisible {
			window.DrawScaledText(
				"_",
				ofsX+(cur%9)*tileSize,
				ofsY+(cur/9)*tileSize,
				textScale,
				draw.White,
			)
		}

		window.DrawScaledText("[ENTER] solve\n [ESC]  clear", 195, 640-150, 2, draw.White)
		window.DrawScaledText(errMsg, 20, 640-50, 1.5, draw.Red)
	})
}
