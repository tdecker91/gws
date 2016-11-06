package main

import (
	"encoding/json"
)

type Message struct {
	data interface{}
}

func parseMessage(rawData []byte) (*Message, error) {
	m := &Message{
		data: nil,
	}

	err := json.Unmarshal(rawData, &m.data)

	return m, err
}

func (m *Message) Compose() ([]byte, error) {
	return json.Marshal(m.data)
}
