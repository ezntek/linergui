package main

import (
	"linergui/window"

	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1600,900,"LinerGUI test")
	rl.SetTargetFPS(60)

	testDialog := window.NewWindow(20,20,800,600,rl.SkyBlue, 15, "Test Window")

	for !rl.WindowShouldClose() {
		testDialog.Update()
		rl.BeginDrawing()
			testDialog.Draw()
			rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()
	}
}