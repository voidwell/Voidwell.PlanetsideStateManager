package main

import (
	"encoding/json"
	"regexp"
	"strconv"
)

//var censusEvents = []string{"BattleRankUp", "ContinentLock", "ContinentUnlock", "Death", "FacilityControl", "MetagameEvent", "PlayerFacilityCapture", "PlayerFacilityDefend", "PlayerLogin", "PlayerLogout", "VehicleDestroy"}
var censusEvents = []string{"PlayerLogin"}
var censusCharacters = []string{"all"}

//var censusWorlds = []string{"all"}
var censusWorlds = []string{"1"}

var regServer = regexp.MustCompile(`EventServerEndpoint_(?P<WorldName>.*)_(?P<WorldID>.*)`)

type CensusManager struct {
	state        *StateManager
	censusClient *CensusClient
	exitChan     chan struct{}
	processor    *CensusProcessor
}

func NewCensusManager(state *StateManager, censusKey string, censusNamespace string) *CensusManager {
	return &CensusManager{
		state:        state,
		censusClient: NewCensusClient(censusKey, censusNamespace),
		exitChan:     make(chan struct{}),
		processor:    NewCensusProcessor(state),
	}
}

func (m *CensusManager) Connect() {
	go m.censusClient.Connect()

	for {
		select {
		case <-m.exitChan:
			return
		case <-m.censusClient.OnConnected:
			go m.censusClient.Subscribe(censusEvents, censusCharacters, censusWorlds)
		case msg := <-m.censusClient.MessageChan:
			go m.processMessage(msg)
		}
	}
}

func (m *CensusManager) processMessage(msg []byte) {
	var jsonMap map[string]interface{}
	json.Unmarshal(msg, &jsonMap)

	if jsonMap["type"] != nil && jsonMap["type"] == "serviceStateChanged" {
		m.processServiceStateChanged(msg)
	} else {
		m.processServiceEvent(jsonMap)
	}
}

func (m *CensusManager) processServiceStateChanged(msg []byte) {
	var serviceStateMessage CensusServiceState
	err := json.Unmarshal(msg, &serviceStateMessage)
	if err != nil {
		return
	}

	matchList := regServer.FindStringSubmatch(serviceStateMessage.Detail)

	worldName := matchList[1]

	if worldID, err := strconv.ParseInt(matchList[2], 10, 64); err == nil {
		m.state.worldManager.SetWorldState(int(worldID), worldName, serviceStateMessage.IsOnline)
	}
}

func (m *CensusManager) processServiceEvent(jsonMap map[string]interface{}) {
	jPayload := jsonMap["payload"]
	if jPayload == nil {
		return
	}

	payloadBytes, _ := json.Marshal(jPayload)

	var payloadMessage CensusPayloadMessage
	err := json.Unmarshal(payloadBytes, &payloadMessage)

	if err != nil || payloadMessage.EventName == "" {
		return
	}

	switch payloadMessage.EventName {
	case "AchievementEarned":
		var eventMsg AchievementEarnedMessage
		json.Unmarshal(payloadBytes, &eventMsg)
		m.processor.ProcessAchievementEarned(eventMsg)
	case "BattleRankUp":
		var eventMsg BattleRankUpMessage
		json.Unmarshal(payloadBytes, &eventMsg)
		m.processor.ProcessBattleRankUp(eventMsg)
	case "ContinentLock":
		var eventMsg ContinentLockMessage
		json.Unmarshal(payloadBytes, &eventMsg)
		m.processor.ProcessContinentLock(eventMsg)
	case "ContinentUnlock":
		var eventMsg ContinentUnlockMessage
		json.Unmarshal(payloadBytes, &eventMsg)
		m.processor.ProcessContinentUnlock(eventMsg)
	case "Death":
		var eventMsg DeathMessage
		json.Unmarshal(payloadBytes, &eventMsg)
		m.processor.ProcessDeath(eventMsg)
	case "FacilityControl":
		var eventMsg FacilityControlMessage
		json.Unmarshal(payloadBytes, &eventMsg)
		m.processor.ProcessFacilityControl(eventMsg)
	case "GainExperience":
		var eventMsg GainExperienceMessage
		json.Unmarshal(payloadBytes, &eventMsg)
		m.processor.ProcessGainExperience(eventMsg)
	case "MetagameEvent":
		var eventMsg MetagameEventMessage
		json.Unmarshal(payloadBytes, &eventMsg)
		m.processor.ProcessMetagameEvent(eventMsg)
	case "PlayerFacilityCapture":
		var eventMsg PlayerFacilityCaptureMessage
		json.Unmarshal(payloadBytes, &eventMsg)
		m.processor.ProcessPlayerFacilityCapture(eventMsg)
	case "PlayerFacilityDefend":
		var eventMsg PlayerFacilityDefendMessage
		json.Unmarshal(payloadBytes, &eventMsg)
		m.processor.ProcessPlayerFacilityDefend(eventMsg)
	case "PlayerLogin":
		var eventMsg PlayerLoginMessage
		json.Unmarshal(payloadBytes, &eventMsg)
		m.processor.ProcessPlayerLogin(eventMsg)
	case "PlayerLogout":
		var eventMsg PlayerLogoutMessage
		json.Unmarshal(payloadBytes, &eventMsg)
		m.processor.ProcessPlayerLogout(eventMsg)
	case "VehicleDestroy":
		var eventMsg VehicleDestroyMessage
		json.Unmarshal(payloadBytes, &eventMsg)
		m.processor.ProcessVehicleDestroy(eventMsg)
	}
}

func (m *CensusManager) Disconnect() {
	m.censusClient.Close()
}
