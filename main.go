package main

import (
	"fmt"
	"linergui/window"

	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1600,900,"LinerGUI test")
	rl.SetTargetFPS(60)

	windows := []*window.Window{
		window.NewWindow(50,50,800,600,rl.SkyBlue, 15, "Test Window 1"),
		window.NewWindow(20,20,800,600,rl.SkyBlue, 15, "Test Window 2"),
	}
	hasFocusIndex := 0
	windows[hasFocusIndex].HasFocus = true

	for !rl.WindowShouldClose() {
		windows[hasFocusIndex].HasFocus = true
		if rl.IsKeyPressed(rl.KeyTab) {
			if hasFocusIndex == len(windows)-1 {
				windows[hasFocusIndex].HasFocus = false
				hasFocusIndex = 0
			} else {
				windows[hasFocusIndex].HasFocus = false
				hasFocusIndex++
			}
		}
		if rl.IsKeyPressed(rl.KeyS) {
			windows = append(windows, window.NewWindow(30,30,800,600,rl.SkyBlue, 15, fmt.Sprintf("Test Window %d", len(windows)+1)))
		}

		for _, window := range windows {
			window.Update()
		}
		rl.BeginDrawing()
			for _, window := range windows {
				if !window.HasFocus { window.Draw() }
			}
			for _, window := range windows {
				if window.HasFocus { window.Draw() }
			}
			rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()
	}
}