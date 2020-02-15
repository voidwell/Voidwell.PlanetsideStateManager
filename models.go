package main

import "time"

type CensusHeartbeat struct {
	lastHeartbeat time.Time
	contents      interface{}
}
