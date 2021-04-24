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

	onScreenButtons []Button

	buttonRelations map[int]func()

	paused bool

	nextID int
}

func (g *Game) pauseGame() {
	g.paused = true
}

func (g *Game) ManageInput() {
	if r.IsMouseButtonPressed(r.MouseLeftButton) {
		x := r.GetMouseX()
		y := r.GetMouseY()

		for i := 0; i < len(g.onScreenButtons); i++ {
			if (x <= g.onScreenButtons[i].rect.ToInt32().X+g.onScreenButtons[i].rect.ToInt32().Width && x >= g.onScreenButtons[i].rect.ToInt32().X) &&
				(y <= g.onScreenButtons[i].rect.ToInt32().Y+g.onScreenButtons[i].rect.ToInt32().Height && y >= g.onScreenButtons[i].rect.ToInt32().Y) {
				g.ButtonPressed(g.onScreenButtons[i].id)
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
	g.onScreenButtons = append(g.onScreenButtons, button)
}

// Update - Update game
func (g *Game) Update() {

	g.ManageInput()

	if r.IsKeyPressed(r.KeyP) {
		g.paused = false
	}

}

// Draw - Draw game
func (g *Game) Draw() {
	r.BeginDrawing()

	r.ClearBackground(r.RayWhite)

	r.DrawText("Hello World!", 240, 180, 48, r.DarkGray)

	for i := 0; i < len(g.onScreenButtons); i++ {
		r.DrawRectangleRec(g.onScreenButtons[i].rect, r.SkyBlue)
		r.DrawTextRec(r.GetFontDefault(), g.onScreenButtons[i].text, g.onScreenButtons[i].rect, 20, 5, false, r.DarkGray)
	}

	if g.paused {
		r.DrawRectangle(g.screenWidth/2-r.MeasureText("GAME PAUSED", 40)/2-50, g.screenHeight/2-60, r.MeasureText("GAME PAUSED", 40)+100, 60, r.LightGray)
		r.DrawText("GAME PAUSED", g.screenWidth/2-r.MeasureText("GAME PAUSED", 40)/2, g.screenHeight/2-40, 40, r.Gray)
	}

	r.EndDrawing()
}
