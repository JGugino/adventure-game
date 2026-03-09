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
	PLAYER        ObjectTag = "player"
	UI_BACKGROUND ObjectTag = "ui_background"
	UI_MIDGROUND  ObjectTag = "ui_midground"
	UI_FOREGROUND ObjectTag = "ui_foreground"
)

type ObjectMetadata struct {
	Id       string     `json:"id"`
	Type     ObjectType `json:"type"`
	Tag      ObjectTag  `json:"tag"`
	Position rl.Vector2 `json:"position"`
	Size     rl.Vector2 `json:"size"`
}

type ObjectMovement struct {
	Speed         float32    `json:"speed"`
	Velocity      rl.Vector2 `json:"velocity"`
	VelocityLimit rl.Vector2 `json:"velocityLimit"`
}

type Object interface {
	Update(deltaTime float32, drag float32) error
	Render() error
	GetId() string
	GetType() ObjectType
	GetTag() ObjectTag
}

type ObjectManager struct {
	Objects   []Object
	DebugMode bool
}

func (o *ObjectManager) Update(deltaTime float32, drag float32) {
	//INFO: Update all registered Objects
	for _, object := range o.Objects {
		err := object.Update(deltaTime, drag)

		if err != nil {
			LogError("Object failed to update")
			LogError(err.Error())
			return
		}
	}
}

func (o *ObjectManager) Render() {
	//INFO: Update all registered Objects
	for _, object := range o.Objects {

		var err error = nil

		object.Render()

		if object.GetType() == UI {
			if object.GetTag() == UI_BACKGROUND {
				err = object.Render()
			}
			if object.GetTag() == UI_MIDGROUND {
				err = object.Render()
			}

			if object.GetTag() == UI_FOREGROUND {
				err = object.Render()
			}
		}

		if err != nil {
			LogError("Object failed to render")
			LogError(err.Error())
			return
		}
	}
}

// Registers a new object to the object manager
func (o *ObjectManager) RegisterObject(object Object) {
	if !o.objectExists(object.GetId()) {
		if o.DebugMode {
			LogInfo(fmt.Sprintf("Object registered: %s", object.GetId()))
		}

		o.Objects = append(o.Objects, object)
		return
	}

	if o.DebugMode {
		LogError(fmt.Sprintf("Object already registered: %s", object.GetId()))
	}
}

// Removes an object from the object manager
func (o *ObjectManager) RemoveObject(object Object) {
	objIndex := o.getObjectIndexById(object.GetId())

	if objIndex != -1 {
		if o.DebugMode {
			LogInfo(fmt.Sprintf("Object removed: %s", object.GetId()))
		}

		//Replace selected object with last object
		o.Objects[objIndex] = o.Objects[len(o.Objects)-1]

		//Sets objects equal to objects mius the last object
		o.Objects = o.Objects[:len(o.Objects)-1]
	}

	if o.DebugMode {
		LogError(fmt.Sprintf("Object doesn't exist: %s", object.GetId()))
	}
}

// Return true or false depending on if the object exists
func (o *ObjectManager) objectExists(objId string) bool {
	for _, obj := range o.Objects {

		if obj.GetId() == objId {
			return true
		}

	}

	return false
}

// Gets the index of an item inside of a slice with a matching id, return -1 if not fuund
func (o *ObjectManager) getObjectIndexById(id string) int {
	for i, item := range o.Objects {
		if item.GetId() == id {
			return i
		}

	}

	return -1
}
