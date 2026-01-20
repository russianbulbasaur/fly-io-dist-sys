package models

import (
	"context"
	"encoding/json"
	"log"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type ConnectedNode struct {
	ID                 string
	node               *maelstrom.Node
	newMessagesChannel chan any
}

func NewConnectedNode(id string, node *maelstrom.Node) *ConnectedNode {
	return &ConnectedNode{
		id,
		node,
		make(chan any, 1000),
	}
}

func (cNode *ConnectedNode) StartSyncing() {
	for {
		select {
		case message := <-cNode.newMessagesChannel:
			log.Println("sending ", message)
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
			if body["type"] != "gossip_ok" {
				cNode.newMessagesChannel <- message
			}
		}
	}
}

func (cNode *ConnectedNode) NewMessage(message any) {
	cNode.newMessagesChannel <- message
}
