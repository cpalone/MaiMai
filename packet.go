package maimai

import (
	"encoding/json"
	"errors"
)

// PacketType indicates the type of a packet's payload.
type PacketType string

// PacketEvent is the skeleton of a packet, its payload is composed of another type or types.
type PacketEvent struct {
	ID    string          `json:"id"`
	Type  PacketType      `json:"type"`
	Data  json.RawMessage `json:"data,omitempty"`
	Error string          `json:"error,omitempty"`
}

// Message is a unit of data associated with a text message sent on the service.
type Message struct {
	ID              string `json:"id"`
	Parent          string `json:"parent"`
	PreviousEditID  string `json:"previous_edit_id,omitempty"`
	Time            int64  `json:"time"`
	Sender          User   `json:"sender"`
	Content         string `json:"content"`
	EncryptionKeyID string `json:"encryption_key_id,omitempty"`
	Edited          int    `json:"edited,omitempty"`
	Deleted         int    `json:"deleted,omitempty"`
}

// PingEvent encodes the server's information on when this ping occurred and when the next will.
type PingEvent struct {
	Time int64 `json:"time"`
	Next int64 `json:"next"`
}

// User encodes the information about a user in the room. Name may be duplicated within a room
type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ServerID  string `json:"server_id"`
	ServerEra string `json:"server_era"`
}

type SendCommand struct {
	Content string `json:"content"`
	Parent  string `json:"parent"`
}

// SendEvent is a packet type that contains a Message only.
type SendEvent Message

// These give named constants to the packet types.
const (
	PingReplyType = "ping-reply"
	PingEventType = "ping-event"

	SendType      = "send"
	SendEventType = "send-event"
)

// Payload unmarshals the packet payload into the proper Event type and returns it.
func (p *PacketEvent) Payload() (interface{}, error) {
	var payload interface{}
	switch p.Type {
	case PingEventType:
		payload = &PingEvent{}
	case SendEventType:
		payload = &SendEvent{}
	case SendType:
		payload = &SendCommand{}
	default:
		return p.Data, errors.New("Unexpected packet type.")
	}
	err := json.Unmarshal(p.Data, &payload)
	return payload, err
}
