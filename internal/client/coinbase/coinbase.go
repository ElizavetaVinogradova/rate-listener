package coinbase

import (
	"fmt"
	"rates-listener/internal/service"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/websocket"
)

type CoinBaseClient struct {
	Conn *websocket.Conn
}

func NewCoinBaseClient(url string) (*CoinBaseClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	request := RequestMessage{
		Type: "subscribe",
		Channels: []Channel{
			{
				Name:       "ticker",
				ProductIDS: []string{"ETH-BTC", "BTC-USD", "BTC-EUR"},
			},
		},
	}

	requestMsg, err := MarshalRequest(request)
	if err != nil {
		conn.Close()
		return &CoinBaseClient{}, err
	}

	err = conn.WriteMessage(websocket.TextMessage, requestMsg)
	if err != nil {
		conn.Close()
		return &CoinBaseClient{}, fmt.Errorf("write message to coinbase: %w", err)
	}
	_, _, err = conn.ReadMessage() //read and ignore the very first message with request info
	if err != nil {
		conn.Close()
		return &CoinBaseClient{}, fmt.Errorf("read the very first coinbase's message: %w", err)
	}

	return &CoinBaseClient{Conn: conn}, nil
}

func (c *CoinBaseClient) readMessage() (TickClientDTO, error) {
	messageType, message, err := c.Conn.ReadMessage()
	if err != nil {
		return TickClientDTO{}, fmt.Errorf("read coinbase's message: %w", err)
	}
	log.Debugf("Received message type %d: %s", messageType, message)

	tickDTO, err := UnmarshalResponse(message)
	if err != nil {
		return TickClientDTO{}, err
	}
	return tickDTO, nil
}

func (c *CoinBaseClient) GetTicksBatch(batchSize int) ([]service.Tick, error) {
	messages := make([]service.Tick, batchSize)

	for i := range messages {
		tickDTO, err := c.readMessage()
		if err != nil {
			return nil, err
		}
		messages[i] = tickDTO.toServiceTick()
	}
	return messages, nil
}
