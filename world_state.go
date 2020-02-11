package main

type WorldState struct {
	ID            int
	Name          string
	IsOnline      bool
	ZoneOwnership map[int][]*CensusMapOwnership
}

func NewWorldState(id int, name string, isOnline bool) *WorldState {
	return &WorldState{
		ID:            id,
		Name:          name,
		IsOnline:      isOnline,
		ZoneOwnership: make(map[int][]*CensusMapOwnership),
	}
}

func (s *WorldState) RegionFactionChange(zoneID int, regionID int, factionID int) {
	regions := s.ZoneOwnership[zoneID]
	for _, r := range regions {
		if r.RegionID == regionID {
			r.FactionID = factionID
			return
		}
	}
}
