package handlers

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func NewGossipHandler(n *maelstrom.Node) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		messages.Mu.Lock()
		messages.Messages[body["message"]] = true
		log.Println(messages.Messages)
		messages.Mu.Unlock()

		delete(body, "message")
		body["type"] = "gossip_ok"
		return n.Reply(msg, body)
	}
}
