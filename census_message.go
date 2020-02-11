package main

import (
	"time"
)

type CensusHeartbeat struct {
	lastHeartbeat time.Time
	contents      interface{}
}

type CensusServiceState struct {
	Detail      string `json:"detail"`
	IsOnline    bool   `json:"online,string"`
	Service     string `json:"service"`
	MessageType string `json:"type"`
}

type CensusPayloadMessage struct {
	EventName string `json:"event_name"`
	WorldID   int    `json:"world_id,string"`
	ZoneID    *int   `json:"zone_id,string,omitempty"`
	Timestamp int64  `json:"timestamp,string"`
}

type AchievementEarnedMessage struct {
	CensusPayloadMessage
	CharacterID   string `json:"character_id"`
	AchievementID string `json:"achievement_id"`
}

type BattleRankUpMessage struct {
	CensusPayloadMessage
	CharacterID string `json:"character_id"`
	BattleRank  int    `json:"battle_rank,string"`
}

type ContinentLockMessage struct {
	CensusPayloadMessage
	TriggeringFaction int     `json:"triggering_faction,string"`
	MetagameEventID   int     `json:"metagame_event_id,string"`
	VsPopulation      float32 `json:"vs_population,string"`
	NcPopulation      float32 `json:"nc_population,string"`
	TrPopulation      float32 `json:"tr_population,string"`
}

type ContinentUnlockMessage struct {
	CensusPayloadMessage
	TriggeringFaction int `json:"triggering_faction,string"`
	MetagameEventID   int `json:"metagame_event_id,string"`
}

type DeathMessage struct {
	CensusPayloadMessage
	AttackerCharacterID string `json:"attacker_character_id"`
	AttackerFireModeID  *int   `json:"attacker_fire_mode_id,string,omitempty"`
	AttackerLoadoutID   *int   `json:"attacker_loadout_id,string,omitempty"`
	AttackerVehicleID   *int   `json:"attacker_vehicle_id,string,omitempty"`
	AttackerWeaponID    *int   `json:"attacker_weapon_id,string,omitempty"`
	CharacterID         string `json:"character_id"`
	CharacterLoadoutID  *int   `json:"character_loadout_id,string,omitempty"`
	IsHeadshot          bool   `json:"is_headshot,string"`
}

type FacilityControlMessage struct {
	CensusPayloadMessage
	FacilityID   int    `json:"facility_id,string"`
	NewFactionID int    `json:"new_faction_id,string"`
	OldFactionID int    `json:"old_faction_id,string"`
	DurationHeld int    `json:"duration_held,string"`
	OutfitID     string `json:"outfit_id"`
}

type GainExperienceMessage struct {
	CensusPayloadMessage
	CharacterID  string `json:"character_id"`
	ExperienceID int    `json:"experience_id,string"`
	Amount       int    `json:"Amount,string"`
	LoadoutID    *int   `json:"loadout_id,string,omitempty"`
	OtherID      string `json:"other_id"`
}

type MetagameEventMessage struct {
	CensusPayloadMessage
	InstanceID         int     `json:"instance_id,string"`
	MetagameEventID    int     `json:"metagame_event_id,string"`
	MetagameEventState string  `json:"metagame_event_state"`
	FactionVS          float32 `json:"faction_vs,string"`
	FactionNC          float32 `json:"faction_nc,string"`
	FactionTR          float32 `json:"faction_tr,string"`
	ExperienceBonus    float32 `json:"experience_bonus,string"`
}

type PlayerFacilityCaptureMessage struct {
	CensusPayloadMessage
	CharacterID string `json:"character_id"`
	FacilityID  int    `json:"facility_id,string"`
	OutfitID    string `json:"outfit_id"`
}

type PlayerFacilityDefendMessage struct {
	CensusPayloadMessage
	CharacterID string `json:"character_id"`
	FacilityID  int    `json:"facility_id,string"`
	OutfitID    string `json:"outfit_id"`
}

type PlayerLoginMessage struct {
	CensusPayloadMessage
	CharacterID string `json:"character_id"`
}

type PlayerLogoutMessage struct {
	CensusPayloadMessage
	CharacterID string `json:"character_id"`
}

type VehicleDestroyMessage struct {
	CensusPayloadMessage
	AttackerCharacterID string `json:"attacker_character_id"`
	AttackerLoadoutID   *int   `json:"attacker_loadout_id,string,omitempty"`
	AttackerVehicleID   *int   `json:"attacker_vehicle_id,string,omitempty"`
	AttackerWeaponID    *int   `json:"attacker_weapon_id,string,omitempty"`
	CharacterID         string `json:"character_id"`
	FacilityID          *int   `json:"facility_id,string,omitempty"`
	FactionID           *int   `json:"faction_id,string,omitempty"`
	VehicleID           int    `json:"vehicle_id,string"`
}

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

type CensusMapOwnership struct {
	RegionID  int `json:"RegionId,string"`
	FactionID int `json:"FactionID,string"`
}
