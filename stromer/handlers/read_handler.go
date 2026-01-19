package handlers

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func NewReadHandler(n *maelstrom.Node) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		body["type"] = "read_ok"
		res := make([]any, 0)
		for m, _ := range messages.Messages {
			res = append(res, m)
		}
		body["messages"] = res
		return n.Reply(msg, body)
	}
}
