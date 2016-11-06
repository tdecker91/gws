package main

type MessageType int

const (
	ClientConnected MessageType = iota
	ClientDisconnected
	ClientMessage
)

type Message struct {
	ClientId string
	Type     MessageType
	Data     []byte
}

func newMessage(clientId string, mt MessageType, data []byte) *Message {
	return &Message{
		ClientId: clientId,
		Type:     mt,
		Data:     data,
	}
}
