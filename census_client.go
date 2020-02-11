package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/recws-org/recws"
)

type CensusClient struct {
	wsURL       url.URL
	closeChan   chan struct{}
	connection  recws.RecConn
	MessageChan chan []byte
	OnConnected chan bool
}

func NewCensusClient(censusServiceKey string, censusNamespace string) *CensusClient {
	queryString := fmt.Sprintf("environment=%s&service-id=s:%s", censusNamespace, censusServiceKey)

	return &CensusClient{
		wsURL:       url.URL{Scheme: "wss", Host: "push.planetside2.com", Path: "streaming", RawQuery: queryString},
		closeChan:   make(chan struct{}),
		MessageChan: make(chan []byte),
		OnConnected: make(chan bool),
	}
}

func (client *CensusClient) Connect() {
	log.Printf("connecting to %s", client.wsURL.String())

	client.connection = recws.RecConn{
		KeepAliveTimeout: 10 * time.Second,
		SubscribeHandler: func() error {
			client.OnConnected <- true
			return nil
		},
	}

	client.connection.Dial(client.wsURL.String(), nil)

	done := make(chan struct{})

	defer client.connection.Close()

	go func() {
		defer close(done)
		for {
			_, message, err := client.connection.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			client.MessageChan <- message
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-client.closeChan:
			log.Println("close")

			err := client.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-client.closeChan:
			case <-time.After(time.Second):
			}
		}
	}
}

func (client *CensusClient) Subscribe(eventNames []string, characters []string, worlds []string) {
	subscription := NewCensusSubscriptionMessage(eventNames, characters, worlds)
	err := client.connection.WriteJSON(subscription)
	if err != nil {
		log.Println("subscription:", err)
	}
}

func (client *CensusClient) Close() {
	close(client.closeChan)
}
