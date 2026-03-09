package ui

import (
	"adventure-game/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UIButton struct {
	Metadata *engine.ObjectMetadata
}

func (b UIButton) Update(deltaTime float32, drag float32) error {

	return nil
}

func (b UIButton) Render() error {
	rl.DrawRectangle(int32(b.Metadata.Position.X), int32(b.Metadata.Position.Y), int32(b.Metadata.Size.X), int32(b.Metadata.Size.Y), rl.Blue)
	return nil
}

func (b UIButton) GetId() string {
	return b.Metadata.Id
}

func (b UIButton) GetType() string {
	return b.Metadata.Id
}

func (b UIButton) GetTag() string {
	return b.Metadata.Id
}
