package handlers

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func NewBroadcastHandler(n *maelstrom.Node) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		message := int(body["message"].(float64))
		delete(body, "message")
		body["type"] = "broadcast_ok"

		messages.Mu.Lock()
		messages.Messages = append(messages.Messages, message)
		messages.Mu.Unlock()

		go gossip(msg.Src, message)

		return n.Reply(msg, body)
	}
}

func gossip(src string, message int) {
	for _, node := range connectedNodes {
		if node.ID == src {
			continue
		}
		node.NewMessage(message)
	}
}
