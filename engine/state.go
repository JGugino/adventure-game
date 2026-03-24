package engine

import (
	"errors"
)

const (
	NO_STATE = "no-state"
)

type StateMetadata struct {
	Id           string         `json:"id"`
	StateManager *StateManager  `json:"stateManager"`
	ObjManager   *ObjectManager `json:"objects"`
	Active       bool           `json:"active"`
}

type State interface {
	Init()
	Update(deltaTime float32, drag float32)
	Render()
	GetId() string
	GetObjectManager() *ObjectManager
	GetActive() bool
}

type StateManager struct {
	States        map[string]State `json:"states"`
	CurrentState  string           `json:"currentState"`
	PreviousState string           `json:"previousState"`
}

func (s *StateManager) Init(startingState string) {
	LogInfo("Init State Manager")
	s.CurrentState = startingState
	s.PreviousState = NO_STATE

	s.States[s.CurrentState].Init()
}

func (s *StateManager) Update(deltaTime float32, drag float32) {
	if s.CurrentState != NO_STATE {
		s.States[s.CurrentState].Update(deltaTime, drag)
	}
}

func (s *StateManager) Render() {
	if s.CurrentState != NO_STATE {
		s.States[s.CurrentState].Render()
	}
}

func (s *StateManager) ChangeState(id string) (string, error) {
	if state, ok := s.States[id]; ok {
		s.PreviousState = s.CurrentState
		s.CurrentState = id

		s.States[id].Init()
		return state.GetId(), nil
	}

	return NO_STATE, errors.New("state-doesnt-exist")
}

func (s *StateManager) RegisterState(id string, state State) error {
	if s.States == nil {
		s.States = make(map[string]State)
	}

	if _, ok := s.States[id]; !ok {
		s.States[id] = state
		return nil
	}

	return errors.New("state-already-exists")
}
