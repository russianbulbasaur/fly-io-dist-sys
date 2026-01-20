package handlers

import (
	"context"
	"encoding/json"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func NewKVReadHandler(n *maelstrom.Node, kv *maelstrom.KV) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		val, _ := kv.ReadInt(ctx, kvCounterKey)
		body["type"] = "read_ok"
		body["value"] = val
		return n.Reply(msg, body)
	}
}
