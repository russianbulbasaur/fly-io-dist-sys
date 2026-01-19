package handlers

import (
	"encoding/json"
	"fmt"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type IDCounter struct {
	counter int
	mu      sync.Mutex
}

func NewGenerateHandler(n *maelstrom.Node) func(msg maelstrom.Message) error {
	counter := IDCounter{}
	return func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		body["type"] = "generate_ok"
		counter.mu.Lock()
		body["id"] = fmt.Sprintf("%s:%d", msg.Dest, counter.counter)
		counter.counter += 1
		counter.mu.Unlock()
		return n.Reply(msg, body)
	}
}
