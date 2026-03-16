package ui

import (
	"adventure-game/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UIButton struct {
	Metadata  *engine.ObjectMetadata
	Text      engine.ObjectText
	Colors    engine.ObjectColors
	Clickable *engine.ObjectClickable
}

func (b UIButton) Update(deltaTime float32, drag float32) error {

	//INFO: Check if mouse cursor is inside of button
	if rl.GetMousePosition().X > b.Metadata.Position.X && rl.GetMousePosition().X < b.Metadata.Position.X+b.Metadata.Size.X {
		if rl.GetMousePosition().Y > b.Metadata.Position.Y && rl.GetMousePosition().Y < b.Metadata.Position.Y+b.Metadata.Size.Y {

			//INFO: If inside and left mouse button is clicked, trigger callback
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				b.Clickable.Callback()
			}

		}
	}

	return nil
}

func (b UIButton) Render() error {
	rl.DrawRectangle(int32(b.Metadata.Position.X), int32(b.Metadata.Position.Y), int32(b.Metadata.Size.X), int32(b.Metadata.Size.Y), b.Colors.PrimaryColor)

	textSize := rl.MeasureTextEx(rl.GetFontDefault(), b.Text.Text, b.Text.FontSize, 1)

	txtX := b.Metadata.Position.X + (b.Metadata.Size.X / 2) - textSize.X/2
	txtY := b.Metadata.Position.Y + (b.Metadata.Size.Y / 2) - textSize.Y/2

	textPosition := rl.Vector2{X: float32(txtX), Y: float32(txtY)}

	rl.DrawTextPro(rl.GetFontDefault(), b.Text.Text, textPosition, rl.Vector2Zero(), 0, b.Text.FontSize, 1, rl.Black)
	return nil
}

func (b UIButton) GetId() string {
	return b.Metadata.Id
}

func (b UIButton) GetType() engine.ObjectType {
	return b.Metadata.Type
}

func (b UIButton) GetTag() engine.ObjectTag {
	return b.Metadata.Tag
}

func (b UIButton) GetPosition() rl.Vector2 {
	return b.Metadata.Position
}

func (b UIButton) GetSize() rl.Vector2 {
	return b.Metadata.Size
}
