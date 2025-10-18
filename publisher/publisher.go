package publisher

import (
	"encoding/json"
	"sync"

	"github.com/nats-io/nats.go"
)

type Publisher struct {
	natsConn *nats.Conn

	mu sync.Mutex
}

func NewPublisher(serverUrl string) (*Publisher, error) {
	conn, err := nats.Connect(serverUrl)
	if err != nil {
		return nil, err
	}

	return &Publisher{natsConn: conn}, nil
}

// Publish - send a provided message to the connected NATS server with a provided subject
//
// TODO: add protobuf
func (p *Publisher) Publish(subject string, data interface{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := p.natsConn.Publish(subject, payload); err != nil {
		return nil
	}

	return nil
}

func (p *Publisher) Close() {
	p.natsConn.Close()
}
