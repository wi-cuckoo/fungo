package pubsub

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type trans struct {
	sync.RWMutex
	portals map[string]*portal
}

func (t *trans) Register(ctx context.Context, id string) (<-chan Payload, error) {
	t.RLock()
	_, ok := t.portals[id]
	t.RUnlock()
	if ok {
		return nil, ErrExisted
	}
	p := portal{
		done: make(chan struct{}),
		pch:  make(chan Payload, 100),
	}
	t.Lock()
	t.portals[id] = &p
	t.Unlock()
	return p.pch, nil
}

func (t *trans) Issue(ctx context.Context, id string, payload Payload) error {
	t.RLock()
	p, ok := t.portals[id]
	t.RUnlock()
	if !ok {
		return ErrNotFound
	}
	// time.Sleep(time.Millisecond * 100)
	select {
	case <-p.done:
		return ErrNoReceiver
	case p.pch <- payload:
	case <-time.After(time.Millisecond * 3):
		return ErrFullChan
	}
	return nil
}

func (t *trans) Unregister(ctx context.Context, id string) error {
	t.RLock()
	p, ok := t.portals[id]
	t.RUnlock()
	if !ok {
		return ErrNotFound
	}
	close(p.done)
	t.Lock()
	delete(t.portals, id)
	t.Unlock()
	return nil
}

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
	return &trans{
		portals: make(map[string]*portal, 0),
	}
}
