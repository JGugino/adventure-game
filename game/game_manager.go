package game

import (
	"adventure-game/engine"
	"adventure-game/engine/ui"
	"adventure-game/game/objects"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameManager struct {
	objectManager engine.ObjectManager
	DebugMode     bool
}

func (g *GameManager) Init() error {
	engine.LogInfo("Initializing game manager")

	//INFO: Init object manager
	g.objectManager = engine.ObjectManager{
		DebugMode: true,
	}

	g.objectManager.Init()

	//INFO: Register the player object
	player := objects.Player{
		Metadata: &engine.ObjectMetadata{
			Id:       "player",
			Type:     engine.CONTROLLABLE,
			Tag:      engine.PLAYER,
			Position: rl.Vector2{X: 0, Y: 0},
			Size:     rl.Vector2{X: 36, Y: 36},
		},
		Movement: &engine.ObjectMovement{
			Speed:         10.0,
			Velocity:      rl.Vector2Zero(),
			VelocityLimit: rl.Vector2{X: 2.0, Y: 2.0},
		},
		DebugMode: false,
	}

	button := ui.UIButton{
		Metadata: &engine.ObjectMetadata{
			Id:       "test-btn",
			Type:     engine.UI,
			Tag:      engine.UI_FOREGROUND,
			Position: rl.Vector2{X: 60, Y: 60},
			Size:     rl.Vector2{X: 400, Y: 100},
		},
		Text: engine.ObjectText{
			Text:     "Example Button",
			FontSize: 32,
		},
		Colors: engine.ObjectColors{
			PrimaryColor: rl.Red,
		},
		Clickable: &engine.ObjectClickable{
			Callback: func() {
				fmt.Println("Button Clicked")
			},
		},
	}

	text := ui.UIText{
		Metadata: &engine.ObjectMetadata{
			Id:       "title-text",
			Type:     engine.UI,
			Tag:      engine.UI_FOREGROUND,
			Position: rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: 66},
			Size:     rl.Vector2{X: 48, Y: 48},
		},
		Text: engine.ObjectText{
			Text:     "Adventure Game",
			FontSize: 48,
		},
		Colors: engine.ObjectColors{
			PrimaryColor: rl.Black,
		},
	}

	g.objectManager.RegisterObject(string(engine.PLAYER), player)
	g.objectManager.RegisterObject(string(engine.UI_FOREGROUND), text)
	g.objectManager.RegisterObject(string(engine.UI_FOREGROUND), button)

	return nil
}

func (g *GameManager) Update(deltaTime float32) {
	g.objectManager.Update(deltaTime, 0.08)
}

func (g *GameManager) Render() {
	g.objectManager.Render()
}
