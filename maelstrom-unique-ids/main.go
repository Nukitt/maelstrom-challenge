package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	uuid "github.com/google/uuid"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	// initialise a node
	n := maelstrom.NewNode()

	// register a handler for generate
	n.Handle("generate", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// updating the message type to return back
		body["type"] = "generate_ok"
		// body["id"] = generateID()
		body["id"] = generateRandID()

		// reply with the updated message
		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}

func generateRandID() string {
	// generate a random number and a timestamp
	randomNum := rand.Intn(1000000)
	timestamp := time.Now().UnixMicro()

	// concatenate the random number and the timestamp
	id := fmt.Sprintf("%d%d", randomNum, timestamp)
	return id
}

func generateID() string {
	// generate a uuid
	id := uuid.New()
	return id.String()
}
