package main

type CensusMapRegion struct {
	MapRegionID    int      `json:"map_region_id,string"`
	ZoneID         int      `json:"zone_id,string"`
	FacilityID     int      `json:"facility_id,string"`
	FacilityName   string   `json:"facility_name"`
	FacilityTypeID int      `json:"facility_type_id,string"`
	FacilityType   string   `json:"facility_type"`
	LocationX      *float32 `json:"location_x,string"`
	LocationY      *float32 `json:"location_y,string"`
	LocationZ      *float32 `json:"location_z,string"`
}

type CensusFacilityLink struct {
	ZoneID      int    `json:"zone_id,string"`
	FacilityIDA int    `json:"facility_id_a,string"`
	FacilityIDB int    `json:"facility_id_b,string"`
	Description string `json:"description"`
}

type CensusZone struct {
	ZoneID int    `json:"zone_id,string"`
	Name   string `json:"code"`
}

type CensusMap struct {
	ZoneID  int                  `json:"ZoneId,string"`
	Regions *CensusMapRegionRows `json:"Regions"`
}

type CensusMapRegionRows struct {
	IsList int             `json:"IsList,string"`
	Row    []*CensusMapRow `json:"Row"`
}

type CensusMapRow struct {
	RowData *CensusRegionOwnership `json:"RowData"`
}

type CensusRegionOwnership struct {
	RegionID  int `json:"RegionId,string"`
	FactionID int `json:"FactionID,string"`
}
