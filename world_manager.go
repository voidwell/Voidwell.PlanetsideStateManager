package main

type WorldManager struct {
	state       *StateManager
	worldStates []*WorldState
	worldZones  []*WorldZone
}

func NewWorldManager(state *StateManager) *WorldManager {
	worldZones := GenerateWorld(state.collections)

	return &WorldManager{
		state:       state,
		worldZones:  worldZones,
		worldStates: make([]*WorldState, 0),
	}
}

func (m *WorldManager) SetWorldState(worldID int, worldName string, isOnline bool) {
	s := m.getWorldState(worldID)
	if s == nil {
		s = NewWorldState(worldID, worldName, isOnline)
		m.worldStates = append(m.worldStates, s)
	}

	if isOnline {
		for _, z := range m.worldZones {
			ownership, _ := m.state.collections.GetMapOwnership(worldID, z.ZoneID)
			s.ZoneOwnership[z.ZoneID] = ownership
		}
	}
}

func (m *WorldManager) getWorldState(worldID int) *WorldState {
	for _, s := range m.worldStates {
		if s.ID == worldID {
			return s
		}
	}
	return nil
}
