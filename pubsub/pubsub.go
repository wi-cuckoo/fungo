package pubsub

import (
	"context"
	"errors"
)

// ErrNotFound is returned when the named topic does not exist.
var ErrNotFound = errors.New("topic not found")

// ErrTimeout is returned when publish timeout
var ErrTimeout = errors.New("publish timeout")

// ErrCtxDone is returned when publish timeout
var ErrCtxDone = errors.New("context done")

// ErrExisted ....
var ErrExisted = errors.New("portal has existed")

// ErrFullChan ....
var ErrFullChan = errors.New("portal channel is full")

// ErrNoReceiver ....
var ErrNoReceiver = errors.New("portal receiver has gone")

// Message defines a published message.
type Message struct {
	// ID identifies this message.
	ID string `json:"id,omitempty"`

	// Data is the actual data in the entry.
	Data []byte `json:"data"`

	// Labels represents the key-value pairs the entry is lebeled with.
	Labels map[string]string `json:"labels,omitempty"`
}

// Unsubscribe to sub
type Unsubscribe func()

// Publisher defines a mechanism for communicating messages from a group
// of senders, called publishers, to a group of consumers.
type Publisher interface {
	// Create creates the named topic.
	Create(topic string)

	// Remove removes the named topic.
	Remove(topic string)

	// Publish publishes the message.
	Publish(ctx context.Context, topic string, message Message) error

	// Subscribe subscribes to the topic
	Subscribe(topic string) (<-chan Message, Unsubscribe, error)
}

// Payload data
type Payload interface{}

// RegNotify register portal and wait notify
// tips: just for one receiver, multi senders
type RegNotify interface {
	// Register portal
	Register(c context.Context, portal string) (<-chan Payload, error)
	// Issue payload to the portal
	Issue(c context.Context, portal string, p Payload) error
	// Unregister
	Unregister(c context.Context, portal string) error
}
