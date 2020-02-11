package main

import (
	"log"
	"os"
	"os/signal"
)

type StateManager struct {
	censusManager *CensusManager
	worldManager  *WorldManager
	playerManager *PlayerManager
	alertManager  *AlertManager
	collections   *CensusCollections
}

func main() {
	censusServiceKey := os.Getenv("CensusKey")
	censusNamespace := os.Getenv("CensusNamespace")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	manager := &StateManager{}

	manager.censusManager = NewCensusManager(manager, censusServiceKey, censusNamespace)
	manager.collections = NewCensusCollections(censusServiceKey, censusNamespace)

	defer manager.censusManager.Disconnect()

	go manager.censusManager.Connect()

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			return
		}
	}
}
