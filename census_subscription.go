package main

type CensusSubscriptionMessage struct {
	Service    string   `json:"service,omitempty"`
	Action     string   `json:"action,omitempty"`
	Characters []string `json:"characters,omitempty"`
	Worlds     []string `json:"worlds,omitempty"`
	EventNames []string `json:"eventNames,omitempty"`
}

func NewCensusSubscriptionMessage(eventNames []string, characters []string, worlds []string) *CensusSubscriptionMessage {
	return &CensusSubscriptionMessage{
		Service:    "event",
		Action:     "subscribe",
		Characters: characters,
		Worlds:     worlds,
		EventNames: eventNames,
	}
}
