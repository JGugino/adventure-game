package game

import (
	"adventure-game/engine"
	"adventure-game/game/states"
)

const (
	TITLE_STATE = "title"
	GAME_STATE  = "game"
)

type GameManager struct {
	stateManager engine.StateManager
	DebugMode    bool
}

func (g *GameManager) Init() error {
	engine.LogInfo("Initializing game manager")

	//INFO: Init StateManager
	g.stateManager = engine.StateManager{}

	g.stateManager.RegisterState(TITLE_STATE, states.Title{
		Metadata: &engine.StateMetadata{
			Id: TITLE_STATE,
			ObjManager: &engine.ObjectManager{
				Objects:   make(map[string][]engine.Object),
				DebugMode: true,
			},
			StateManager: &g.stateManager,
		},
	})

	g.stateManager.RegisterState(GAME_STATE, states.Game{
		Metadata: &engine.StateMetadata{
			Id: GAME_STATE,
			ObjManager: &engine.ObjectManager{
				Objects:   make(map[string][]engine.Object),
				DebugMode: true,
			},
			StateManager: &g.stateManager,
		},
	})

	g.stateManager.Init(TITLE_STATE)

	return nil
}

func (g *GameManager) Update(deltaTime float32) {
	g.stateManager.Update(deltaTime, 0.08)
}

func (g *GameManager) Render() {
	g.stateManager.Render()
}
