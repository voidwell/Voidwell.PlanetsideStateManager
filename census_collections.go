package main

import (
	"encoding/json"

	"github.com/lampjaw/censusgo"
)

type CensusCollections struct {
	queryBuilder *censusgo.QueryBuilder
}

var censusEndpoint = "http://census.daybreakgames.com map_region?zone_id=2"

func NewCensusCollections(censusKey string, censusNamespace string) *CensusCollections {
	return &CensusCollections{
		queryBuilder: censusgo.NewQueryBuilder(censusKey, censusNamespace),
	}
}

func (c *CensusCollections) GetMapRegions(zoneID int) (regions []*CensusMapRegion, err error) {
	q := c.queryBuilder.NewQuery("map_region")
	q.Where("zone_id").Equals(zoneID)
	q.SetLimit(5000)

	resp, err := q.GetResults()
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &regions)

	return
}

func (c *CensusCollections) GetMapLinks(zoneID int) (links []*CensusFacilityLink, err error) {
	q := c.queryBuilder.NewQuery("facility_link")
	q.Where("zone_id").Equals(zoneID)
	q.SetLimit(5000)

	resp, err := q.GetResults()
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &links)

	return
}

func (c *CensusCollections) GetMapZones() (zones []*CensusZone, err error) {
	q := c.queryBuilder.NewQuery("zone")
	q.SetLanguage(censusgo.LangEnglish)
	q.SetLimit(100)
	q.Where("zone_id").IsLessThan(90)

	resp, err := q.GetResults()
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &zones)

	return
}

func (c *CensusCollections) GetMapOwnership(worldID int, zoneID int) (regionOwnerships []*CensusRegionOwnership, err error) {
	q := c.queryBuilder.NewQuery("map")
	q.Where("world_id").Equals(worldID)
	q.Where("zone_ids").Equals(zoneID)

	resp, err := q.GetResults()
	if err != nil {
		return
	}

	var censusMap []CensusMap
	err = json.Unmarshal(resp, &censusMap)
	if err != nil {
		return
	}

	regionOwnerships = make([]*CensusRegionOwnership, 0)

	for _, row := range censusMap[0].Regions.Row {
		regionOwnerships = append(regionOwnerships, row.RowData)
	}

	return
}
