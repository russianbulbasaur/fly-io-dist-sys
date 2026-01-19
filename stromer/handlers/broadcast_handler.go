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

		messages.Mu.Lock()
		if !messages.Messages[body["message"]] {
			messages.Messages[body["message"]] = true
			connectedNodes.mu.Lock()
			for _, node := range connectedNodes.connectedNodes {
				if node.ID == msg.Src {
					continue
				}
				node.NewMessage(body["message"])
			}
			connectedNodes.mu.Unlock()
		}
		messages.Mu.Unlock()

		delete(body, "message")
		body["type"] = "broadcast_ok"
		return n.Reply(msg, body)
	}
}
