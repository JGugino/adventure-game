package engine

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ObjectType string
type ObjectTag string

const (
	//INFO: Object Types
	CONTROLLABLE ObjectType = "controllable"
	NPC          ObjectType = "npc"
	ENVIROMENT   ObjectType = "enviroment"
	UI           ObjectType = "ui"

	//INFO: Object Tags
	DEFAULT       ObjectTag = "default"
	PLAYER        ObjectTag = "player"
	UI_BACKGROUND ObjectTag = "ui_background"
	UI_MIDGROUND  ObjectTag = "ui_midground"
	UI_FOREGROUND ObjectTag = "ui_foreground"
)

// DOCS: Basic metadata for the object (id, tag, pos, etc.)
type ObjectMetadata struct {
	Id       string     `json:"id"`
	Type     ObjectType `json:"type"`
	Tag      ObjectTag  `json:"tag"`
	Position rl.Vector2 `json:"position"`
	Size     rl.Vector2 `json:"size"`
}

// DOCS: Metadata for values used for moving an object (speed, velocity, etc.)
type ObjectMovement struct {
	Speed         float32    `json:"speed"`
	Velocity      rl.Vector2 `json:"velocity"`
	VelocityLimit rl.Vector2 `json:"velocityLimit"`
}

// DOCS: Callback for a clickable object (mostly used for ui)
type ObjectClickable struct {
	Callback func() `json:"callback"`
}

// DOCS: Colors to use for basic object rendering
type ObjectColors struct {
	PrimaryColor   rl.Color `json:"primaryColor"`
	SecondaryColor rl.Color `json:"secondaryColor"`
}

// DOCS: Text to display on an object
type ObjectText struct {
	Text     string  `json:"text"`
	FontSize float32 `json:"fontSize"`
}

// DOCS: Interface that defines required funcs for an Object
type Object interface {
	Update(deltaTime float32, drag float32) error
	Render() error
	GetId() string
	GetType() ObjectType
	GetTag() ObjectTag
	GetPosition() rl.Vector2
	GetSize() rl.Vector2
}

type ObjectManager struct {
	Objects   map[string][]Object
	DebugMode bool
}

func (o *ObjectManager) Init() {
	o.Objects = make(map[string][]Object)
}

func (o *ObjectManager) Update(deltaTime float32, drag float32) {
	//INFO: Update all registered Objects
	for _, objs := range o.Objects {
		for _, obj := range objs {
			err := obj.Update(deltaTime, drag)

			if err != nil {
				LogError("Object failed to update")
				LogError(err.Error())
				return
			}
		}
	}
}

func (o *ObjectManager) Render() {

	//INFO: Default Render
	for _, object := range o.Objects[string(DEFAULT)] {
		err := object.Render()

		if o.DebugMode {
			pos := object.GetPosition()
			size := object.GetSize()

			rl.DrawRectangleLines(int32(pos.X), int32(pos.Y), int32(size.X), int32(size.Y), rl.DarkGreen)
		}

		if err != nil {
			LogError(fmt.Sprintf("Object has failed to render - layer: %s", UI_BACKGROUND))
			LogError(err.Error())
			return
		}
	}

	//INFO: Player Render
	for _, object := range o.Objects[string(PLAYER)] {
		err := object.Render()

		if err != nil {
			LogError(fmt.Sprintf("Object has failed to render - layer: %s", UI_BACKGROUND))
			LogError(err.Error())
			return
		}
	}

	//INFO: ### UI ###

	//INFO: UI Background Render
	for _, object := range o.Objects[string(UI_BACKGROUND)] {
		err := object.Render()

		if err != nil {
			LogError(fmt.Sprintf("Object has failed to render - layer: %s", UI_BACKGROUND))
			LogError(err.Error())
			return
		}
	}
	//INFO: UI Midground Render
	for _, object := range o.Objects[string(UI_MIDGROUND)] {
		err := object.Render()

		if err != nil {
			LogError(fmt.Sprintf("Object has failed to render - layer: %s", UI_MIDGROUND))
			LogError(err.Error())
			return
		}
	}
	//INFO: UI Foreground Render
	for _, object := range o.Objects[string(UI_FOREGROUND)] {
		err := object.Render()

		if err != nil {
			LogError(fmt.Sprintf("Object has failed to render - layer: %s", UI_FOREGROUND))
			LogError(err.Error())
			return
		}
	}
}

// DOCS: Registers a new object to the object manager
func (o *ObjectManager) RegisterObject(objectsLayer string, object Object) {
	if !o.objectExists(o.Objects[objectsLayer], object.GetId()) {
		if o.DebugMode {
			LogInfo(fmt.Sprintf("Object registered: %s", object.GetId()))
		}

		o.Objects[objectsLayer] = append(o.Objects[objectsLayer], object)
		return
	}

	if o.DebugMode {
		LogError(fmt.Sprintf("Object already registered: %s, %s", object.GetId(), objectsLayer))
	}
}

// DOCS: Removes an object from the object manager
func (o *ObjectManager) RemoveObject(objectLayer string, object Object) {
	objIndex := o.getObjectIndexById(o.Objects[objectLayer], object.GetId())

	if objIndex != -1 {
		if o.DebugMode {
			LogInfo(fmt.Sprintf("Object removed: %s", object.GetId()))
		}

		//Replace selected object with last object
		o.Objects[objectLayer][objIndex] = o.Objects[objectLayer][len(o.Objects)-1]

		//Sets objects equal to objects mius the last object
		o.Objects[objectLayer] = o.Objects[objectLayer][:len(o.Objects)-1]
	}

	if o.DebugMode {
		LogError(fmt.Sprintf("Object doesn't exist: %s, %s", object.GetId(), objectLayer))
	}
}

// DOCS: Return true or false depending on if the object exists
func (o *ObjectManager) objectExists(objects []Object, objId string) bool {
	for _, obj := range objects {

		if obj.GetId() == objId {
			return true
		}

	}

	return false
}

// DOCS: Gets the index of an item inside of a slice with a matching id, return -1 if not fuund
func (o *ObjectManager) getObjectIndexById(objects []Object, id string) int {
	for i, item := range objects {
		if item.GetId() == id {
			return i
		}

	}

	return -1
}
