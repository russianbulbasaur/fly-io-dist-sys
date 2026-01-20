package models

import (
	"context"
	"encoding/json"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type ConnectedNode struct {
	ID                 string
	node               *maelstrom.Node
	newMessagesChannel chan int
	ticker             *time.Ticker
}

func NewConnectedNode(id string, node *maelstrom.Node) *ConnectedNode {
	return &ConnectedNode{
		id,
		node,
		make(chan int, 1000),
		time.NewTicker(time.Second * 1),
	}
}

func (cNode *ConnectedNode) StartSyncing() {
	for {
		select {
		case message := <-cNode.newMessagesChannel:
			body := make(map[string]any)
			body["type"] = "gossip"
			body["message"] = message
			ctx1, _ := context.WithTimeout(context.Background(), time.Second*1)
			msg, err := cNode.node.SyncRPC(ctx1, cNode.ID, body)
			if err != nil {
				cNode.newMessagesChannel <- message
			}
			var body1 map[string]any
			if err := json.Unmarshal(msg.Body, &body1); err != nil {
				cNode.newMessagesChannel <- message
			}
			if body1["type"] != "gossip_ok" {
				cNode.newMessagesChannel <- message
			}
		default:
		}
	}
}

func (cNode *ConnectedNode) NewMessage(message int) {
	cNode.newMessagesChannel <- message
}
