package consistentlb

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"
	"stathat.com/c/consistent"
)

var hashElt string

// BalancerName is the name of the consistent balancer.
const BalancerName = "consistent"

func newconsistentBuilder() balancer.Builder {
	return &consistentBuilder{}
}

type consistentBuilder struct{}

func (*consistentBuilder) Build(cc balancer.ClientConn, opt balancer.BuildOptions) balancer.Balancer {
	return &consistentBalancer{
		cc: cc,
		ct: consistent.New(),
	}
}

func (*consistentBuilder) Name() string {
	return BalancerName
}

type consistentBalancer struct {
	cc     balancer.ClientConn
	sc     balancer.SubConn
	ct     *consistent.Consistent
}

func (b *consistentBalancer) HandleResolvedAddrs(addrs []resolver.Address, err error) {
	if err != nil {
		grpclog.Infof("consistentBalancer: HandleResolvedAddrs called with error %v", err)
		return
	}
	// re-hash to get the match address
	tmpAddrs := make([]string, len(addrs))
	for i, a := range addrs {
		tmpAddrs[i] = a.Addr
	}
	b.ct.Set(tmpAddrs)
	target, _ := b.ct.Get(hashElt)
	fmt.Println(tmpAddrs, target, b.sc)
	if b.sc != nil {
		b.sc.UpdateAddresses([]resolver.Address{{Addr: target}})
		b.sc.Connect()
		return
	}
	b.sc, err = b.cc.NewSubConn([]resolver.Address{{Addr: target}}, balancer.NewSubConnOptions{})
	if err != nil {
		grpclog.Errorf("consistentBalancer: failed to NewSubConn: %v", err)
		return
	}
	b.cc.UpdateBalancerState(connectivity.Idle, &picker{sc: b.sc})
	b.sc.Connect()
}

func (b *consistentBalancer) HandleSubConnStateChange(sc balancer.SubConn, s connectivity.State) {
	grpclog.Infof("consistentBalancer: HandleSubConnStateChange: %p, %v", sc, s)
	if b.sc != sc {
		grpclog.Infof("consistentBalancer: ignored state change because sc is not recognized")
		return
	}
	if s == connectivity.Shutdown {
		b.sc = nil
		return
	}

	switch s {
	case connectivity.Ready, connectivity.Idle:
		b.cc.UpdateBalancerState(s, &picker{sc: sc})
	case connectivity.Connecting:
		b.cc.UpdateBalancerState(s, &picker{err: balancer.ErrNoSubConnAvailable})
	case connectivity.TransientFailure:
		b.cc.UpdateBalancerState(s, &picker{err: balancer.ErrTransientFailure})
	}
}

func (b *consistentBalancer) Close() {}

type picker struct {
	err error
	sc  balancer.SubConn
}

func (p *picker) Pick(ctx context.Context, opts balancer.PickOptions) (balancer.SubConn, func(balancer.DoneInfo), error) {
	if p.err != nil {
		return nil, nil, p.err
	}
	return p.sc, nil, nil
}

func init() {
	hashElt, _ = os.Hostname()
	balancer.Register(newconsistentBuilder())
}
