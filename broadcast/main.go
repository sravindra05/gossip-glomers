package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type nodeWrapper struct {
	*maelstrom.Node
	mu       sync.RWMutex
	messages map[int]bool
	topology map[string]any
}

func (n *nodeWrapper) propagate(message int) error {
	neighbours := n.topology[n.ID()].([]interface{})

	requestBody := map[string]any{
		"type":    "broadcast",
		"message": message,
	}

	for _, neighbour := range neighbours {
		err := n.RPC(neighbour.(string), requestBody, func(msg maelstrom.Message) error {
			return nil
		})
		if err != nil {
			return fmt.Errorf("error sending message to neighbours: %w", err)
		}
	}

	return nil
}

func main() {
	n := nodeWrapper{Node: maelstrom.NewNode(), messages: map[int]bool{}}

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// https://stackoverflow.com/a/29690346
		message := int(body["message"].(float64))

		// if we have seen this message before then do nothing
		if _, ok := n.messages[message]; ok {
			body["type"] = "broadcast_ok"
			delete(body, "message")
			return n.Reply(msg, body)
		}

		n.mu.Lock()
		n.messages[message] = true
		n.mu.Unlock()

		body["type"] = "broadcast_ok"
		delete(body, "message")

		// Can we make this better?
		err := n.propagate(message)
		if err != nil {
			return err
		}

		return n.Reply(msg, body)
	})

	n.Handle("broadcast_ok", func(msg maelstrom.Message) error {
		return nil
	})

	n.Handle("read", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		n.mu.RLock()
		defer n.mu.RUnlock()

		var messages []int

		for key, _ := range n.messages {
			messages = append(messages, key)
		}

		body["type"] = "read_ok"
		body["messages"] = messages

		return n.Reply(msg, body)
	})

	n.Handle("topology", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "topology_ok"

		topology, ok := body["topology"].(map[string]any)
		if !ok {
			return errors.New("failed to decode topology")
		}
		n.topology = topology

		delete(body, "topology")

		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
