package pubsub

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type pipe struct {
	sync.Map
}

type portal struct {
	done chan struct{}
	pch  chan Payload
}

func (p *pipe) Register(ctx context.Context, id string) (<-chan Payload, error) {
	if _, ok := p.Load(id); ok {
		return nil, errors.New("portal has registered")
	}
	po := &portal{
		done: make(chan struct{}),
		pch:  make(chan Payload, 100),
	}
	p.Store(id, po)
	return po.pch, nil
}

func (p *pipe) Issue(ctx context.Context, id string, payload Payload) error {
	val, ok := p.Load(id)
	if !ok {
		return ErrNotFound
	}
	po := val.(*portal)
	if len(po.pch) > 99 {
		return errors.New("full channel")
	}
	// time.Sleep(time.Millisecond * 100)
	select {
	case <-po.done:
		fmt.Println("got done signal to quit")
	default:
		po.pch <- payload
	}
	return nil
}

func (p *pipe) Unregister(ctx context.Context, id string) error {
	val, ok := p.Load(id)
	if !ok {
		return ErrNotFound
	}
	po := val.(*portal)
	close(po.done)
	p.Delete(id)
	return nil
}

// NewRegNotify ...
func NewRegNotify() RegNotify {
	return &pipe{}
}
