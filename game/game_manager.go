package game

import (
	"adventure-game/engine"
	"adventure-game/game/objects"

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
		DebugMode: g.DebugMode,
	}

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
		DebugMode: true,
	}

	g.objectManager.RegisterObject(player)

	return nil
}

func (g *GameManager) Update(deltaTime float32) {
	g.objectManager.Update(deltaTime, 0.08)
}

func (g *GameManager) Render() {
	g.objectManager.Render()
}
