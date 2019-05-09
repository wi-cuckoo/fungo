package pubsub

import (
	"context"
	"sync"
)

const defaulSubQueue = 100

type subscriber struct {
	ch chan Message
}

type publisher struct {
	sync.Mutex

	topics map[string]*topic
}

// New Publisher
func New() Publisher {
	return &publisher{
		topics: make(map[string]*topic),
	}
}

func (p *publisher) Create(desc string) {
	p.Lock()
	if _, ok := p.topics[desc]; !ok {
		p.topics[desc] = newtopic(desc)
	}
	p.Unlock()
}

func (p *publisher) Publish(ctx context.Context, dest string, message Message) error {
	p.Lock()
	t, ok := p.topics[dest]
	p.Unlock()
	if !ok {
		return ErrNotFound
	}
	return t.publish(ctx, message)
}

func (p *publisher) Subscribe(dest string) (<-chan Message, Unsubscribe, error) {
	p.Lock()
	t, ok := p.topics[dest]
	p.Unlock()
	if !ok {
		return nil, nil, ErrNotFound
	}
	s := &subscriber{
		ch: make(chan Message, defaulSubQueue),
	}
	t.subscribe(s)

	unsub := func() {
		t.unsubscribe(s)
		close(s.ch)
	}

	return s.ch, unsub, nil
}

func (p *publisher) Remove(dest string) {
	p.Lock()
	t, ok := p.topics[dest]
	if ok {
		delete(p.topics, dest)
		t.close()
	}
	p.Unlock()
}
