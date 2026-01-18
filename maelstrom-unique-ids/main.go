package main


import  (
   "encoding/json"
   "log"
   "fmt"
   "sync"
   maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)


type Counter struct {
     counter int
     mu sync.Mutex
}


func main() {
   counter := Counter{
       counter: 0,
   }
   n := maelstrom.NewNode()
   n.Handle("generate",func(msg maelstrom.Message) error {
       var body map[string]any
       if err := json.Unmarshal(msg.Body, &body); err != nil {
            return err       
       }
       counter.mu.Lock()
       defer counter.mu.Unlock()
       id := counter.counter + 1
       counter.counter += 1
       body["type"] = "generate_ok"
       body["id"] = fmt.Sprintf("%s%d",msg.Dest,id)
       return n.Reply(msg,body)
   })

   if err := n.Run(); err != nil {
       log.Fatal(err)
   }
}
