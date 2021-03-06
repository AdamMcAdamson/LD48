package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

type Building int

const (
	None          Building = iota // None = 0
	Research                      // 1
	MiningOutpost                 // 2
	CareRoom                      // 3 ....
	Crafting
	Port
	CoalGen
	BioGen
)

var BuildingStrings = []string{"None", "Research", "MiningOutpost", "CareRoom", "Crafting", "Port", "CoalGen", "BioGen"}

func (bp Building) String() string {
	return BuildingStrings[bp]
}

type Button struct {
	id int

	rect r.Rectangle

	text string

	pressed bool
}

type FunctionManager struct {
	room Room
}

type Room struct {
	building Building

	rect r.Rectangle
}

type Game struct {
	screenWidth  int32
	screenHeight int32

	camera r.Camera2D

	cameraSpeed int32

	worldButtons  []Button
	screenButtons []Button

	buttonRelations map[int]func()

	paused bool

	bgTexture r.Texture2D
	nextID    int
}

func (g *Game) pauseGame() {
	g.paused = true
}

func (g *Game) ManageInput() {
	if r.IsMouseButtonPressed(r.MouseLeftButton) {
		var pos r.Vector2 = r.GetScreenToWorld2D(r.GetMousePosition(), g.camera)

		for i := 0; i < len(g.worldButtons); i++ {
			if (int32(pos.X) <= g.worldButtons[i].rect.ToInt32().X+g.worldButtons[i].rect.ToInt32().Width && int32(pos.X) >= g.worldButtons[i].rect.ToInt32().X) &&
				(int32(pos.Y) <= g.worldButtons[i].rect.ToInt32().Y+g.worldButtons[i].rect.ToInt32().Height && int32(pos.Y) >= g.worldButtons[i].rect.ToInt32().Y) {
				g.ButtonPressed(g.worldButtons[i].id)
			}
		}
	}
}

func (g *Game) ButtonPressed(id int) {
	g.buttonRelations[id]()

}

func main() {
	game := Game{}
	game.Init()

	r.InitWindow(game.screenWidth, game.screenHeight, "Going Deeper")

	noise := r.GenImageWhiteNoise(2000, 2000, .8)

	game.bgTexture = r.LoadTextureFromImage(noise)

	r.SetTargetFPS(60)

	for !r.WindowShouldClose() {
		game.Update()

		game.Draw()
	}

	r.CloseWindow()
}

// Init - Initialize game
func (g *Game) Init() {
	g.screenWidth = 1280
	g.screenHeight = 720

	g.camera.Target = r.NewVector2(300, 200)
	g.camera.Offset = r.NewVector2(float32(g.screenWidth)/2, float32(g.screenHeight)/2)
	g.camera.Rotation = 0
	g.camera.Zoom = .8

	g.buttonRelations = make(map[int]func())

	g.nextID = 1

	// TESTING
	var button Button

	button.id = g.nextID
	g.nextID++
	button.text = "Pause Game"
	button.rect = r.NewRectangle(300, 50, 200, 100)
	button.pressed = false

	g.buttonRelations[button.id] = g.pauseGame
	g.worldButtons = append(g.worldButtons, button)

}

// Update - Update game
func (g *Game) Update() {

	g.ManageInput()

	if r.IsKeyPressed(r.KeyP) {
		g.paused = false
	}

	if r.IsKeyDown(r.KeyLeftShift) {
		g.cameraSpeed = 10
	} else {
		g.cameraSpeed = 4
	}

	if r.IsKeyDown(r.KeyW) && g.camera.Target.Y > 200 {
		g.camera.Target.Y -= float32(g.cameraSpeed) / g.camera.Zoom
	}
	if r.IsKeyDown(r.KeyA) && g.camera.Target.X > 300 {
		g.camera.Target.X -= float32(g.cameraSpeed) / g.camera.Zoom
	}
	if r.IsKeyDown(r.KeyS) && g.camera.Target.Y < 5000 {
		g.camera.Target.Y += float32(g.cameraSpeed) / g.camera.Zoom
	}
	if r.IsKeyDown(r.KeyD) && g.camera.Target.X < 2000 {
		g.camera.Target.X += float32(g.cameraSpeed) / g.camera.Zoom
	}

	g.camera.Zoom += float32(r.GetMouseWheelMove()) * 0.05
	if g.camera.Zoom > 2 {
		g.camera.Zoom = 2
	} else if g.camera.Zoom < 0.5 {
		g.camera.Zoom = 0.5
	}
}

// Draw - Draw game
func (g *Game) Draw() {
	r.BeginDrawing()
	{
		r.ClearBackground(r.RayWhite)

		r.BeginMode2D(g.camera)
		{
			r.DrawTexture(g.bgTexture, -1000, -600, r.NewColor(0, 100, 100, 255))
			r.DrawTexture(g.bgTexture, 1000, -600, r.NewColor(0, 100, 100, 255))
			r.DrawTexture(g.bgTexture, 1000, 1400, r.NewColor(0, 100, 100, 255))
			r.DrawRectangleGradientV(-1000, -600, 7000, 10000, r.NewColor(0, 140, 255, 245), r.NewColor(0, 0, 10, 245))
			r.DrawText("Hello World!", 240, 180, 48, r.DarkGray)
			for i := 0; i < len(g.worldButtons); i++ {
				r.DrawRectangleRec(g.worldButtons[i].rect, r.SkyBlue)
				r.DrawTextRec(r.GetFontDefault(), g.worldButtons[i].text, g.worldButtons[i].rect, 20, 5, false, r.DarkGray)
			}
		}
		r.EndMode2D()

		if g.paused {
			r.DrawRectangle(g.screenWidth/2-r.MeasureText("GAME PAUSED", 40)/2-50, g.screenHeight/2-60, r.MeasureText("GAME PAUSED", 40)+100, 60, r.LightGray)
			r.DrawText("GAME PAUSED", g.screenWidth/2-r.MeasureText("GAME PAUSED", 40)/2, g.screenHeight/2-40, 40, r.Gray)
		}

		r.DrawFPS(20, 10)

	}
	r.EndDrawing()
}
