package objects

import (
	"adventure-game/engine"
	"fmt"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Metadata  *engine.ObjectMetadata
	Movement  *engine.ObjectMovement
	DebugMode bool
}

func (p Player) Update(deltaTime float32, drag float32) error {

	playerMovement(deltaTime, drag, p.Metadata, p.Movement)

	if p.DebugMode {
		fmt.Printf("Pos X: %f | Pos Y: %f\n", p.Metadata.Position.X, p.Metadata.Position.Y)
		fmt.Printf("Vel X: %f | Vel Y: %f\n", p.Movement.Velocity.X, p.Movement.Velocity.Y)
	}
	return nil
}

func playerMovement(deltaTime float32, drag float32, metadata *engine.ObjectMetadata, movement *engine.ObjectMovement) {

	//INFO: Add to velocity depending on desired direction
	if rl.IsKeyDown(rl.KeyW) {
		if movement.Velocity.Y >= -movement.VelocityLimit.Y {
			movement.Velocity.Y -= movement.Speed * deltaTime
		}
	} else if rl.IsKeyDown(rl.KeyS) {
		if movement.Velocity.Y <= movement.VelocityLimit.Y {
			movement.Velocity.Y += movement.Speed * deltaTime
		}
	}

	if rl.IsKeyDown(rl.KeyA) {
		if movement.Velocity.X >= -movement.VelocityLimit.X {
			movement.Velocity.X -= movement.Speed * deltaTime
		}
	} else if rl.IsKeyDown(rl.KeyD) {
		if movement.Velocity.X <= movement.VelocityLimit.X {
			movement.Velocity.X += movement.Speed * deltaTime
		}
	}

	//INFO: Add the velocity to the players position
	metadata.Position.X += movement.Velocity.X
	metadata.Position.Y += movement.Velocity.Y

	//INFO: Apply drag to player movement
	if movement.Velocity.X < 0.0 {
		v, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", movement.Velocity.X+drag), 32)
		movement.Velocity.X = float32(v)
	} else if movement.Velocity.X > 0.0 {
		v, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", movement.Velocity.X-drag), 32)
		movement.Velocity.X = float32(v)
	}

	if movement.Velocity.Y < 0.0 {
		v, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", movement.Velocity.Y+drag), 32)
		movement.Velocity.Y = float32(v)
	} else if movement.Velocity.Y > 0.0 {
		v, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", movement.Velocity.Y-drag), 32)
		movement.Velocity.Y = float32(v)
	}
}

func (p Player) Render() error {
	rl.DrawRectangle(int32(p.Metadata.Position.X), int32(p.Metadata.Position.Y), int32(p.Metadata.Size.X), int32(p.Metadata.Size.Y), rl.Blue)
	return nil
}

func (p Player) GetId() string {
	return p.Metadata.Id
}

func (p Player) GetTag() engine.ObjectTag {
	return p.Metadata.Tag
}

func (p Player) GetType() engine.ObjectType {
	return p.Metadata.Type
}
