package pubsub

import (
	"context"
	"sync"
)

type subscriber struct {
	receiver chan Message
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

func (p *publisher) Create(ctx context.Context, desc string) {
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

func (p *publisher) Subscribe(dest string) (<-chan Message, error) {
	p.Lock()
	t, ok := p.topics[dest]
	p.Unlock()
	if !ok {
		return nil, ErrNotFound
	}
	s := &subscriber{
		receiver: make(chan Message, 100),
	}
	t.subscribe(s)

	return s.receiver, nil
}

func (p *publisher) Remove(ctx context.Context, dest string) {
	p.Lock()
	t, ok := p.topics[dest]
	if ok {
		delete(p.topics, dest)
		t.close()
	}
	p.Unlock()
}
