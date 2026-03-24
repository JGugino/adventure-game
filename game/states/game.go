package states

import (
	"adventure-game/engine"
	"adventure-game/game/objects"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Metadata *engine.StateMetadata
}

func (g Game) Init() {
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

	g.Metadata.ObjManager.RegisterObject(string(engine.PLAYER), player)
}

func (g Game) Update(deltaTime float32, drag float32) {
	g.Metadata.ObjManager.Update(deltaTime, drag)
}
func (g Game) Render() {
	g.Metadata.ObjManager.Render()
}
func (g Game) GetId() string {
	return g.Metadata.Id
}

func (g Game) GetObjectManager() *engine.ObjectManager {
	return g.Metadata.ObjManager
}

func (g Game) GetActive() bool {
	return g.Metadata.Active
}
