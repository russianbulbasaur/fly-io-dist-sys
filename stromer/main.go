package main

import (
	"stromer/handlers"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	kv := maelstrom.NewSeqKV(n)
	setupHandlers(n, kv)
	if err := n.Run(); err != nil {
		panic(err)
	}
}

func setupHandlers(n *maelstrom.Node, kv *maelstrom.KV) {
	n.Handle("echo", handlers.NewEchoHandler(n))
	n.Handle("generate", handlers.NewGenerateHandler(n))
	n.Handle("broadcast", handlers.NewBroadcastHandler(n))

	// normal read
	// n.Handle("read", handlers.NewKVReadHandler(n,kv))

	// stage 4 kv read
	n.Handle("read", handlers.NewKVReadHandler(n, kv))

	n.Handle("add", handlers.NewAddHandler(n, kv))
	n.Handle("topology", handlers.NewTopologyHandler(n))
	n.Handle("gossip", handlers.NewGossipHandler(n))
}
