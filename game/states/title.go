package states

import (
	"adventure-game/engine"
	"adventure-game/engine/ui"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Title struct {
	Metadata *engine.StateMetadata
}

func (t Title) Init() {

	t.Metadata.ObjManager = &engine.ObjectManager{
		DebugMode: true,
	}

	t.Metadata.ObjManager.Init()

	button := ui.UIButton{
		Metadata: &engine.ObjectMetadata{
			Id:       "test-btn",
			Type:     engine.UI,
			Tag:      engine.UI_FOREGROUND,
			Position: rl.Vector2{X: float32(rl.GetScreenWidth())/2 - 200, Y: float32(rl.GetScreenHeight())/2 - 50},
			Size:     rl.Vector2{X: 400, Y: 100},
		},
		Text: engine.ObjectText{
			Text:     "Play",
			FontSize: 32,
		},
		Colors: engine.ObjectColors{
			PrimaryColor: rl.Red,
		},
		Clickable: &engine.ObjectClickable{
			Callback: func() {
				state, err := t.Metadata.StateManager.ChangeState("game")

				fmt.Println(state, err)
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

	t.Metadata.ObjManager.RegisterObject(string(engine.UI_FOREGROUND), text)
	t.Metadata.ObjManager.RegisterObject(string(engine.UI_FOREGROUND), button)

}

func (t Title) Update(deltaTime float32, drag float32) {
	if t.Metadata.ObjManager != nil {
		t.Metadata.ObjManager.Update(deltaTime, drag)
	}

}
func (t Title) Render() {
	if t.Metadata.ObjManager != nil {
		t.Metadata.ObjManager.Render()
	}
}
func (t Title) GetId() string {
	return t.Metadata.Id
}

func (t Title) GetObjectManager() *engine.ObjectManager {
	return t.Metadata.ObjManager
}

func (t Title) GetActive() bool {
	return t.Metadata.Active
}
