package ui

import (
	"adventure-game/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UIText struct {
	Metadata *engine.ObjectMetadata
	Text     engine.ObjectText
	Colors   engine.ObjectColors
}

func (t UIText) Update(deltaTime float32, drag float32) error {

	return nil
}

func (t UIText) Render() error {

	textSize := rl.MeasureTextEx(rl.GetFontDefault(), t.Text.Text, t.Text.FontSize, 1)

	txtX := t.Metadata.Position.X - textSize.X/2
	txtY := t.Metadata.Position.Y - textSize.Y/2

	textPos := rl.Vector2{X: float32(txtX), Y: float32(txtY)}
	rl.DrawTextPro(rl.GetFontDefault(), t.Text.Text, textPos, rl.Vector2Zero(), 0, t.Text.FontSize, 1, t.Colors.PrimaryColor)
	return nil
}

func (t UIText) GetId() string {
	return t.Metadata.Id
}

func (t UIText) GetType() engine.ObjectType {
	return t.Metadata.Type
}

func (t UIText) GetTag() engine.ObjectTag {
	return t.Metadata.Tag
}

func (t UIText) GetPosition() rl.Vector2 {
	return t.Metadata.Position
}

func (t UIText) GetSize() rl.Vector2 {
	return t.Metadata.Size
}

func (t UIText) GetActive() bool {
	return t.Metadata.Active
}

func (t UIText) SetActive(active bool) {
	t.Metadata.Active = active
}
