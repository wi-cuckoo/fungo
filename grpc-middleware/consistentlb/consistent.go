package consistentlb

import (
	"context"
	"os"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/connectivity"
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
	cc balancer.ClientConn
	sc balancer.SubConn
	ct *consistent.Consistent
}

func (b *consistentBalancer) HandleResolvedAddrs(addrs []resolver.Address, err error) {
	panic("not implemented")
}

func (b *consistentBalancer) HandleSubConnStateChange(sc balancer.SubConn, s connectivity.State) {
	panic("not implemented")
}

func (b *consistentBalancer) UpdateResolverState(s resolver.State) {
	panic("not implemented")
}

func (b *consistentBalancer) UpdateSubConnState(sc balancer.SubConn, state balancer.SubConnState) {
	panic("not implemented")
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
