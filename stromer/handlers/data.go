package handlers

import (
	"stromer/models"
)

var messages = models.Messages{
	Messages: make([]int, 0),
}

var connectedNodes = make(map[string]*models.ConnectedNode)

const kvCounterKey string = "counter"
