package objects

import (
	"adventure-game/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Block struct {
	Metadata  *engine.ObjectMetadata
	DebugMode bool
}

func (b Block) Update(deltaTime float32, drag float32) error {
	return nil
}

func (b Block) Render() error {
	rl.DrawRectangle(int32(b.Metadata.Position.X), int32(b.Metadata.Position.Y), int32(b.Metadata.Size.X), int32(b.Metadata.Size.Y), rl.Red)
	return nil
}

func (b Block) GetId() string {
	return b.Metadata.Id
}

func (b Block) GetTag() engine.ObjectTag {
	return b.Metadata.Tag
}

func (b Block) GetType() engine.ObjectType {
	return b.Metadata.Type
}

func (b Block) GetPosition() rl.Vector2 {
	return b.Metadata.Position
}

func (b Block) GetSize() rl.Vector2 {
	return b.Metadata.Size
}

func (b Block) GetActive() bool {
	return b.Metadata.Active
}

func (b Block) SetActive(active bool) {
	b.Metadata.Active = active
}
