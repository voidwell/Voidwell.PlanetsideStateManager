package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CensusCollections struct {
	endpointURL string
}

var censusEndpoint = "http://census.daybreakgames.com map_region?zone_id=2"

func NewCensusCollections(censusKey string, censusNamespace string) *CensusCollections {
	return &CensusCollections{
		endpointURL: fmt.Sprintf("%s/s:%s/get/%s", censusEndpoint, censusKey, censusNamespace),
	}
}

func (c *CensusCollections) GetMapRegions(zoneID int) ([]CensusMapRegion, error) {
	resp, err := c.getCensusCollection("map_region", fmt.Sprintf("zone_id=%s", zoneID))
	if err != nil {
		log.Println("collection:", resp)
		return nil, err
	}

	var result []CensusMapRegion
	err = json.Unmarshal(resp, &result)

	return result, err
}

func (c *CensusCollections) GetMapLinks(zoneID int) ([]CensusFacilityLink, error) {
	resp, err := c.getCensusCollection("facility_link", fmt.Sprintf("zone_id=%s", zoneID))
	if err != nil {
		log.Println("collection:", resp)
		return nil, err
	}

	var result []CensusFacilityLink
	err = json.Unmarshal(resp, &result)

	return result, err
}

func (c *CensusCollections) GetMapZones() ([]CensusZone, error) {
	resp, err := c.getCensusCollection("zone", "")
	if err != nil {
		log.Println("collection:", resp)
		return nil, err
	}

	var result []CensusZone
	err = json.Unmarshal(resp, &result)

	return result, err
}

func (c *CensusCollections) GetMapOwnership(worldID int, zoneID int) ([]*CensusMapOwnership, error) {
	resp, err := c.getCensusCollection("map", fmt.Sprintf("world_id=%s&zone_ids=%s", worldID, zoneID))
	if err != nil {
		log.Println("collection:", resp)
		return nil, err
	}

	mapOwnership := make([]*CensusMapOwnership, 0)

	var contentBody []map[string]interface{}
	err = json.Unmarshal(resp, &contentBody)

	mapState := contentBody[0]["Regions"]

	mapStateBytes, _ := json.Marshal(mapState)
	regionsBytes := getJsonTokenContent("Regions", mapStateBytes)
	regionRowBytes := getJsonTokenContent("Row", regionsBytes)

	var regionRows []map[string]interface{}
	err = json.Unmarshal(regionRowBytes, &regionRows)

	for _, r := range regionRows {
		rowBytes, _ := json.Marshal(r)
		rowDataBytes := getJsonTokenContent("RowData", rowBytes)

		var ownershipRow *CensusMapOwnership
		json.Unmarshal(rowDataBytes, &ownershipRow)

		mapOwnership = append(mapOwnership, ownershipRow)
	}

	return mapOwnership, nil
}

func (c *CensusCollections) getCensusCollection(collection string, params string) (json.RawMessage, error) {
	if len(params) > 0 {
		params = fmt.Sprintf("&%s", params)
	}
	url := fmt.Sprintf("%s/%s?c:limit=5000%s", c.endpointURL, collection, params)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var contentBody map[string]interface{}
	err = json.Unmarshal(body, &contentBody)

	propertyIndex := fmt.Sprintf("%s_list", collection)
	collectionBody, _ := json.Marshal(contentBody[propertyIndex])

	var jsonResponse json.RawMessage
	err = json.Unmarshal(collectionBody, &jsonResponse)

	return jsonResponse, nil
}

func getJsonTokenContent(tokenID string, content []byte) []byte {
	var contentBody map[string]interface{}
	json.Unmarshal(content, &contentBody)

	tokenBody := contentBody[tokenID]
	tokenBodyBytes, _ := json.Marshal(tokenBody)

	return tokenBodyBytes
}
