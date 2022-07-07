package window

import (
	"image/color"

	"github.com/gen2brain/raylib-go/raylib"
)

type Window struct {
	Border rl.Rectangle
	Window rl.Rectangle
	BorderColor color.RGBA
	Position rl.Vector2 // this is the border's position
	borderWidth float32
}

func (w *Window) Draw() {
	rl.DrawRectangleRec(w.Border, w.BorderColor)
	rl.DrawRectangleRec(w.Window, color.RGBA{235,235,235,255})
}

func (w *Window) Update() {
	if rl.IsKeyDown(rl.KeyLeft) {
		w.Position.X -= 6
	}
	if rl.IsKeyDown(rl.KeyRight) {
		w.Position.X += 6
	}

	if rl.IsKeyDown(rl.KeyUp) {
		w.Position.Y -= 6
	}
	if rl.IsKeyDown(rl.KeyUp) {
		w.Position.Y += 6
	}

	w.Border.X, w.Border.Y = w.Position.X, w.Position.Y
	w.Window.X, w.Window.Y = w.Position.X + w.borderWidth/2, w.Position.Y + w.borderWidth/2
}

func NewWindow(posx float32, posy float32, width float32, height float32, borderColor color.RGBA, borderWidth float32) *Window {
	return &Window{
		Position: rl.Vector2{X: posx, Y: posy},
		Border: rl.NewRectangle(posx, posy, width, height),
		Window: rl.NewRectangle(posx + borderWidth, posy + borderWidth, width - borderWidth, height - borderWidth),
		BorderColor: borderColor,
	}
}