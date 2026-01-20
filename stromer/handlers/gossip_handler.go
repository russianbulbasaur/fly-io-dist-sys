package handlers

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func NewGossipHandler(n *maelstrom.Node) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "gossip_ok"
		message := int(body["message"].(float64))
		delete(body, "message")

		messages.Mu.Lock()
		messages.Messages = append(messages.Messages, message)
		messages.Mu.Unlock()

		return n.Reply(msg, body)
	}
}
