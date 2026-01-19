package models

import (
	"context"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type ConnectedNode struct {
	ID                 string
	ctx                context.Context
	cancelFunc         context.CancelFunc
	node               *maelstrom.Node
	newMessagesChannel chan any
}

func NewConnectedNode(id string, node *maelstrom.Node) *ConnectedNode {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &ConnectedNode{
		id,
		ctx,
		cancelFunc,
		node,
		make(chan any, 1000),
	}
}

func (cNode *ConnectedNode) StartSyncing() {
	for {
		select {
		case <-cNode.ctx.Done():
			close(cNode.newMessagesChannel)
			return
		case message := <-cNode.newMessagesChannel:
			body := make(map[string]any)
			body["type"] = "broadcast"
			body["message"] = message
			_, err := cNode.node.SyncRPC(cNode.ctx, cNode.ID, body)
			if err != nil {
				cNode.newMessagesChannel <- message
			}
		}
	}
}

func (cNode *ConnectedNode) NewMessage(message any) {
	cNode.newMessagesChannel <- message
}

func (cNode *ConnectedNode) Stop() {
	cNode.cancelFunc()
}
