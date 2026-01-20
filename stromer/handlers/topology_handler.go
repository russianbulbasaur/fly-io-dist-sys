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
	var nodesList []string
	for key, _ := range topology {
		if key == n.ID() {
			continue
		}
		nodesList = append(nodesList, key)
	}

	receivedNodesMap := make(map[string]bool)
	for _, node := range nodesList {
		receivedNodesMap[node] = true
	}

	// new nodes
	for id, _ := range receivedNodesMap {
		if _, exists := connectedNodes[id]; !exists {
			// spin new node
			node := models.NewConnectedNode(id, n)
			go node.StartSyncing()
			connectedNodes[id] = node
		}
	}
}
