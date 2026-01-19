package handlers

import (
	"encoding/json"
	"stromer/models"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func NewTopologyHandler(n *maelstrom.Node) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		parseTopology(body["topology"].(map[string]any), n)
		delete(body, "topology")
		body["type"] = "topology_ok"
		return n.Reply(msg, body)
	}
}

func parseTopology(topology map[string]any, n *maelstrom.Node) {
	if _, ok := topology[n.ID()]; !ok {
		return
	}
	receivedNodesMap := make(map[string]bool)
	nodesList := topology[n.ID()].([]any)
	for _, node := range nodesList {
		receivedNodesMap[node.(string)] = true
	}

	connectedNodes.mu.Lock()
	defer connectedNodes.mu.Unlock()

	// check for dead nodes
	for id, node := range connectedNodes.connectedNodes {
		if _, exists := receivedNodesMap[id]; !exists {
			node.Stop()
			delete(connectedNodes.connectedNodes, id)
			continue
		}
	}

	// new nodes
	for id, _ := range receivedNodesMap {
		if _, exists := connectedNodes.connectedNodes[id]; !exists {
			// spin new node
			node := models.NewConnectedNode(id, n)
			go node.StartSyncing()
			connectedNodes.connectedNodes[id] = node
		}
	}
}
