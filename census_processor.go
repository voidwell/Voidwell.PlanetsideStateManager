package main

import "log"

type CensusProcessor struct {
	state *StateManager
}

func NewCensusProcessor(m *StateManager) *CensusProcessor {
	return &CensusProcessor{
		state: m,
	}
}

func (p *CensusProcessor) ProcessAchievementEarned(msg AchievementEarnedMessage) {
	log.Printf("%+v", msg)
}

func (p *CensusProcessor) ProcessBattleRankUp(msg BattleRankUpMessage) {
	log.Printf("%+v", msg)
}

func (p *CensusProcessor) ProcessContinentLock(msg ContinentLockMessage) {
	log.Printf("%+v", msg)
}

func (p *CensusProcessor) ProcessContinentUnlock(msg ContinentUnlockMessage) {
	log.Printf("%+v", msg)
}

func (p *CensusProcessor) ProcessDeath(msg DeathMessage) {
	log.Printf("%+v", msg)
}

func (p *CensusProcessor) ProcessFacilityControl(msg FacilityControlMessage) {
	log.Printf("%+v", msg)
}

func (p *CensusProcessor) ProcessGainExperience(msg GainExperienceMessage) {
	log.Printf("%+v", msg)
}

func (p *CensusProcessor) ProcessMetagameEvent(msg MetagameEventMessage) {
	log.Printf("%+v", msg)
}

func (p *CensusProcessor) ProcessPlayerFacilityCapture(msg PlayerFacilityCaptureMessage) {
	log.Printf("%+v", msg)
}

func (p *CensusProcessor) ProcessPlayerFacilityDefend(msg PlayerFacilityDefendMessage) {
	log.Printf("%+v", msg)
}

func (p *CensusProcessor) ProcessPlayerLogin(msg PlayerLoginMessage) {
	log.Printf("%+v", msg)
}

func (p *CensusProcessor) ProcessPlayerLogout(msg PlayerLogoutMessage) {
	log.Printf("%+v", msg)
}

func (p *CensusProcessor) ProcessVehicleDestroy(msg VehicleDestroyMessage) {
	log.Printf("%+v", msg)
}
