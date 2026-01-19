package models

import (
	"context"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type ConnectedNode struct {
	ID                 string
	Pointer            int
	ctx                context.Context
	cancelFunc         context.CancelFunc
	ticker             *time.Ticker
	node               *maelstrom.Node
	newMessagesChannel chan any
}

func NewConnectedNode(id string, node *maelstrom.Node) *ConnectedNode {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &ConnectedNode{
		id,
		0,
		ctx,
		cancelFunc,
		time.NewTicker(time.Second * 3),
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
		case <-cNode.ticker.C:
			cNode.sync()
		}
	}
}

func (cNode *ConnectedNode) NewMessage(message any) {
	cNode.newMessagesChannel <- message
}

func (cNode *ConnectedNode) sync() {
	// use node to sync
	for {
		select {
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

func (cNode *ConnectedNode) Stop() {
	cNode.cancelFunc()
}
