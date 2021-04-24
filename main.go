package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	screenWidth  int32
	screenHeight int32
}

func main() {
	game := Game{}
	game.Init()

	r.InitWindow(game.screenWidth, game.screenHeight, "Walking-Game")

	r.SetTargetFPS(60)

	for !r.WindowShouldClose() {
		game.Update()

		game.Draw()
	}
	r.CloseWindow()
}

// Init - Initialize game
func (g *Game) Init() {
	g.screenWidth = 800
	g.screenHeight = 450

}

// Update - Update game
func (g *Game) Update() {

}

// Draw - Draw game
func (g *Game) Draw() {
	r.BeginDrawing()

	r.ClearBackground(r.RayWhite)

	r.DrawText("Hello World!", 240, 180, 48, r.DarkGray)
	r.EndDrawing()
}
