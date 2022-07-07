package window

import (
	"image/color"

	"github.com/gen2brain/raylib-go/raylib"
)

type Window struct {
	frames int
	border rl.Rectangle
	borderColor color.RGBA
	Position rl.Vector2 // this is the border's position
	borderWidth float32
	IsHidden bool
	topBar rl.Rectangle
	content rl.Rectangle
	okTextColor bool
	title string
}

func (w *Window) Draw() {
	if !w.IsHidden {
		rl.DrawRectangleRec(w.border, w.borderColor)
		rl.DrawRectangleRec(w.content, rl.RayWhite)
		rl.DrawRectangleRec(w.topBar, color.RGBA{232,232,232,255})
		rl.DrawText(w.title, int32(w.topBar.X) + 5, int32(w.topBar.Y) +5, 20, rl.Black)
		rl.DrawText("[x]", int32(w.topBar.Width + w.topBar.X) - 25, int32(w.topBar.Y)+5, 20, rl.Red)
		if w.okTextColor {
			rl.DrawText("[Continue]", w.content.ToInt32().X + w.content.ToInt32().Width - 105, w.content.ToInt32().Y + w.content.ToInt32().Height - 25, 20, rl.Gray)
		} else {
			rl.DrawText("[Continue]", w.content.ToInt32().X + w.content.ToInt32().Width - 105, w.content.ToInt32().Y + w.content.ToInt32().Height - 25, 20, rl.DarkGray)
		}
	}
}

func (w *Window) Update() {
	if !w.IsHidden {
		w.frames++
		if w.frames % 20 == 0 {
			w.okTextColor = !w.okTextColor
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			w.Position.X -= 6
		}
		if rl.IsKeyDown(rl.KeyRight) {
			w.Position.X += 6
		}

		if rl.IsKeyDown(rl.KeyUp) {
			w.Position.Y -= 6
		}
		if rl.IsKeyDown(rl.KeyDown) {
			w.Position.Y += 6
		}

		
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		w.IsHidden = !w.IsHidden
	}
	w.border.X, w.border.Y = w.Position.X, w.Position.Y
	w.topBar.X, w.topBar.Y = w.Position.X + w.borderWidth/2, w.Position.Y + w.borderWidth/2
	w.content.X, w.content.Y = w.Position.X + w.borderWidth/2, w.Position.Y + w.borderWidth/2 + 30
}

func NewWindow(posx float32, posy float32, width float32, height float32, borderCol color.RGBA, borderWidth float32, t string) *Window {
	return &Window{
		title: t,
		Position: rl.Vector2{X: posx, Y: posy},
		borderWidth: borderWidth,
		border: rl.NewRectangle(posx, posy, width, height),
		content: rl.NewRectangle(posx + borderWidth, posy + borderWidth + 30, width - borderWidth, height - borderWidth - 30),
		topBar: rl.NewRectangle(posx - borderWidth, posy + borderWidth, width - borderWidth, 30),
		borderColor: borderCol,
		IsHidden: false,
	}
}