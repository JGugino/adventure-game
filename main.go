package main

import (
	"adventure-game/game"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WIDTH      = 1080
	HEIGHT     = 720
	TARGET_FPS = 60
)

func main() {
	rl.InitWindow(1080, 720, "Adventure Game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(TARGET_FPS)
	rl.SetExitKey(-1)

	gameManager := game.GameManager{
		DebugMode: true,
	}

	err := gameManager.Init()

	if err != nil {
		fmt.Println("An error has occured while starting the Game Manager ", err.Error())
		return
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		deltaTime := rl.GetFrameTime()

		gameManager.Update(deltaTime)

		//INFO: Clear the screen - draw game stuff after this
		rl.ClearBackground(rl.White)

		gameManager.Render()

		rl.EndDrawing()
	}
}
