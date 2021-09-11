package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// message struct represents a message.
type message struct {
	username    string
	usermessage string
}

// messageDB struct represents the message database.
type messageDB struct {
	messages []message
	sync.Mutex
}

// global message database.
var messagedb messageDB

// posts a message to the server.
func postMessage(username, usermessage string) {
	messagedb.Lock()
	messagedb.messages = append(messagedb.messages, message{username, usermessage})
	messagedb.Unlock()
}

// returns a random message from the server.
func getMessage() string {
	rand.Seed(time.Now().UnixNano())

	messagedb.Lock()
	defer messagedb.Unlock()

	if len(messagedb.messages) == 0 {
		return "No messages..."
	}

	randomMessageObj := messagedb.messages[rand.Intn(len(messagedb.messages))]
	return fmt.Sprintf("%s: %s", randomMessageObj.username, randomMessageObj.usermessage)
}
