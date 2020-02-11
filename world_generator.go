package main

func GenerateWorld(c *CensusCollections) []*WorldZone {
	worldZones := make([]*WorldZone, 0)

	zones, _ := c.GetMapZones()

	for _, zone := range zones {
		worldZone := &WorldZone{
			ZoneID:    zone.ZoneID,
			Name:      zone.Name,
			Warpgates: make([]*ZoneRegion, 0),
			Regions:   make([]*ZoneRegion, 0),
		}

		mapRegions, _ := c.GetMapRegions(zone.ZoneID)
		mapLinks, _ := c.GetMapLinks(zone.ZoneID)

		for _, region := range mapRegions {
			stateRegion := &ZoneRegion{
				RegionID:     region.MapRegionID,
				FacilityID:   region.FacilityID,
				FacilityName: region.FacilityName,
				FacilityType: region.FacilityType,
				PosX:         region.LocationX,
				PosY:         region.LocationY,
				PosZ:         region.LocationZ,
			}

			if stateRegion.FacilityType == "Warpgate" {
				worldZone.Warpgates = append(worldZone.Warpgates, stateRegion)
			}

			worldZone.Regions = append(worldZone.Regions, stateRegion)
		}

		for _, link := range mapLinks {
			facilityA := worldZone.findRegionByFacilityID(link.FacilityIDA)
			facilityB := worldZone.findRegionByFacilityID(link.FacilityIDB)

			if facilityA != nil && facilityB != nil {
				facilityA.Links = append(facilityA.Links, facilityB)
				facilityB.Links = append(facilityB.Links, facilityB)
			}
		}
	}

	return worldZones
}

type WorldZone struct {
	ZoneID    int
	Name      string
	Warpgates []*ZoneRegion
	Regions   []*ZoneRegion
}

func (z *WorldZone) findRegionByFacilityID(facilityID int) *ZoneRegion {
	for _, region := range z.Regions {
		if region.FacilityID == facilityID {
			return region
		}
	}
	return nil
}

type ZoneRegion struct {
	RegionID       int
	FacilityID     int
	FacilityName   string
	FacilityType   string
	FacilityTypeID int
	PosX           *float32
	PosY           *float32
	PosZ           *float32
	Links          []*ZoneRegion
}
