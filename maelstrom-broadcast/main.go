package main

import (
   "encoding/json"
   "log"

   mael "github.com/jepsen-io/maelstrom/demo/go"
)


var messages []any

func main() {
  n := mael.NewNode()
  n.Handle("topology",func(msg mael.Message) error {
     var body map[string]any
     if err := json.Unmarshal(msg.Body, &body); err != nil {
        return err
     }
     delete(body,"topology")
     body["type"] = "topology_ok"
     return n.Reply(msg,body)
  }) 
  
  n.Handle("broadcast",func(msg mael.Message) error {
     var body map[string]any
     if err := json.Unmarshal(msg.Body, &body); err != nil {
        return err
     }
     messages = append(messages,body["message"])
     delete(body,"message")
     body["type"] = "broadcast_ok"
     return n.Reply(msg,body)
  })

  n.Handle("read", func(msg mael.Message) error {
     var body map[string]any
     if err := json.Unmarshal(msg.Body, &body); err != nil {
        return err
     }
     body["type"] = "read_ok"
     body["messages"] = messages
     return n.Reply(msg,body)
  })

  if err := n.Run(); err != nil {
     log.Fatal(err)
  }
}



