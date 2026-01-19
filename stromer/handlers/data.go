package handlers

import (
	"stromer/models"
	"sync"
)

type ConnectedNodes struct {
	connectedNodes map[string]*models.ConnectedNode
	mu             sync.Mutex
}

var messages = models.Messages{
	Messages: make(map[any]bool),
}

var connectedNodes = ConnectedNodes{
	connectedNodes: make(map[string]*models.ConnectedNode),
}
