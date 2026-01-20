package handlers

import (
	"context"
	"encoding/json"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func NewAddHandler(n *maelstrom.Node, kv *maelstrom.KV) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		delta := int(body["delta"].(float64))
		delete(body, "delta")

		for {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
			val, err := kv.ReadInt(ctx, kvCounterKey)
			if err != nil {
				val = 0
			}
			if err = kv.CompareAndSwap(ctx, kvCounterKey, val, val+delta, true); err == nil {
				break
			}
			cancel()
		}
		body["type"] = "add_ok"

		return n.Reply(msg, body)
	}
}
