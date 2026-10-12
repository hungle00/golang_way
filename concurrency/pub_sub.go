package main

import (
    "fmt"
    "sync"
    "time"
)

// Message represents a pub/sub message
type Message struct {
    Topic   string
    Payload interface{}
}

// Subscriber represents a message handler
type Subscriber func(Message)

// PubSub is a simple in-memory pub/sub system
type PubSub struct {
    mu          sync.RWMutex
    subscribers map[string][]Subscriber
    closed      bool
}

// NewPubSub creates a new pub/sub instance
func NewPubSub() *PubSub {
    return &PubSub{
        subscribers: make(map[string][]Subscriber),
    }
}

// Subscribe adds a subscriber to a topic
func (ps *PubSub) Subscribe(topic string, subscriber Subscriber) {
    ps.mu.Lock()
    defer ps.mu.Unlock()
    
    if ps.closed {
        return
    }
    
    ps.subscribers[topic] = append(ps.subscribers[topic], subscriber)
}

// Publish sends a message to all subscribers of a topic
func (ps *PubSub) Publish(topic string, payload interface{}) {
    ps.mu.RLock()
    defer ps.mu.RUnlock()
    
    if ps.closed {
        return
    }
    
    message := Message{
        Topic:   topic,
        Payload: payload,
    }
    
    // Send to all subscribers asynchronously
    for _, subscriber := range ps.subscribers[topic] {
        go subscriber(message)
    }
}

// Close shuts down the pub/sub system
func (ps *PubSub) Close() {
    ps.mu.Lock()
    defer ps.mu.Unlock()
    ps.closed = true
}

func main() {
    pubsub := NewPubSub()
    defer pubsub.Close()
    
    // Subscribe to user events
    pubsub.Subscribe("user.created", func(msg Message) {
        fmt.Printf("Email service: Welcome %v!\n", msg.Payload)
    })
    
    pubsub.Subscribe("user.created", func(msg Message) {
        fmt.Printf("Analytics: New user registered: %v\n", msg.Payload)
    })
    
    pubsub.Subscribe("user.deleted", func(msg Message) {
        fmt.Printf("Cleanup service: Remove user data for %v\n", msg.Payload)
    })
    
    // Publish events
    pubsub.Publish("user.created", "john@example.com")
    pubsub.Publish("user.created", "jane@example.com")
    pubsub.Publish("user.deleted", "old@example.com")
    
    // Wait for async processing
    time.Sleep(100 * time.Millisecond)
}
