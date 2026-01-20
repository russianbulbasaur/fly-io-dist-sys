package main

import (
	"stromer/handlers"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	setupHandlers(n)
	if err := n.Run(); err != nil {
		panic(err)
	}
}

func setupHandlers(n *maelstrom.Node) {
	n.Handle("echo", handlers.NewEchoHandler(n))
	n.Handle("generate", handlers.NewGenerateHandler(n))
	n.Handle("broadcast", handlers.NewBroadcastHandler(n))
	n.Handle("read", handlers.NewReadHandler(n))
	n.Handle("topology", handlers.NewTopologyHandler(n))
	n.Handle("gossip", handlers.NewGossipHandler(n))
}
