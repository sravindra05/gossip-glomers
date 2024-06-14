package main

import (
	"encoding/json"
	"log"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type nodeWrapper struct {
	*maelstrom.Node
	mu       sync.RWMutex
	messages []int
}

func main() {
	n := nodeWrapper{Node: maelstrom.NewNode()}

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		n.mu.Lock()
		defer n.mu.Unlock()
		n.messages = append(n.messages, int(body["message"].(float64)))

		body["type"] = "broadcast_ok"
		delete(body, "message")

		return n.Reply(msg, body)
	})

	n.Handle("read", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		n.mu.RLock()
		defer n.mu.RUnlock()

		body["type"] = "read_ok"
		body["messages"] = n.messages

		return n.Reply(msg, body)
	})

	n.Handle("topology", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Unclear what to do here with this information, ignoring...
		body["type"] = "topology_ok"
		delete(body, "topology")

		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
