package consistentlb_test

import (
	"context"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/wi-cuckoo/fungo/grpc-middleware/consistentlb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
	"google.golang.org/grpc/status"
	testpb "google.golang.org/grpc/test/grpc_testing"
	"stathat.com/c/consistent"
)

var hashElt string

func init() {
	hashElt, _ = os.Hostname()
}

func TestWithOneServer(t *testing.T) {
	r, cleanup := manual.GenerateAndRegisterManualResolver()
	defer cleanup()

	test, err := startTestServers(1)
	if err != nil {
		t.Fatalf("failed to start servers: %v", err)
	}
	defer test.cleanup()

	cc, err := grpc.Dial(r.Scheme()+":///test.server", grpc.WithInsecure(), grpc.WithBalancerName(consistentlb.BalancerName))
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer cc.Close()
	testc := testpb.NewTestServiceClient(cc)
	// The first RPC should fail because there's no address.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if _, err = testc.EmptyCall(ctx, &testpb.Empty{}); err == nil || status.Code(err) != codes.DeadlineExceeded {
		t.Fatalf("EmptyCall() = _, %v, want _, DeadlineExceeded", err)
	}

	r.UpdateState(resolver.State{Addresses: []resolver.Address{{Addr: test.addresses[0]}}})
	// The second RPC should succeed.
	if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}); err != nil {
		t.Fatalf("EmptyCall() = _, %v, want _, <nil>", err)
	}
}

func TestWithMultiServers(t *testing.T) {
	r, cleanup := manual.GenerateAndRegisterManualResolver()
	defer cleanup()

	total := 5
	test, err := startTestServers(total)
	if err != nil {
		t.Fatalf("failed to start servers: %v", err)
	}
	defer test.cleanup()

	cc, err := grpc.Dial(r.Scheme()+":///test.server", grpc.WithInsecure(), grpc.WithBalancerName(consistentlb.BalancerName))
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer cc.Close()
	testc := testpb.NewTestServiceClient(cc)
	// The first RPC should fail because there's no address.
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	if _, err := testc.EmptyCall(ctx, &testpb.Empty{}); err == nil || status.Code(err) != codes.DeadlineExceeded {
		t.Fatalf("EmptyCall() = _, %v, want _, DeadlineExceeded", err)
	}
	// update resolved addresses
	var resolvedAddrs []resolver.Address
	for i := 0; i < total; i++ {
		resolvedAddrs = append(resolvedAddrs, resolver.Address{Addr: test.addresses[i]})
	}
	r.UpdateState(resolver.State{Addresses: resolvedAddrs[:total-1]})

	// get the consistent match
	cstt := consistent.New()
	cstt.Set(test.addresses[:total-1])
	match, _ := cstt.Get(hashElt)

	var p peer.Peer
	if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}, grpc.Peer(&p)); err != nil {
		t.Fatalf("EmptyCall() = _, %v, want _, <nil>", err)
	}
	t.Logf("[INIT] peer=%s, expect=%s", p.Addr.String(), match)
	if p.Addr.String() != match {
		t.Fatalf("the match address should be %s, not %s", match, p.Addr.String())
	}

	// add node
	r.UpdateState(resolver.State{Addresses: resolvedAddrs})
	cstt.Set(test.addresses)
	match, _ = cstt.Get(hashElt)

	if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}, grpc.Peer(&p), grpc.WaitForReady(true)); err != nil {
		t.Fatalf("EmptyCall() = _, %v, want _, <nil>", err)
	}
	t.Logf("[ADD] peer=%s, expect=%s", p.Addr.String(), match)
	if p.Addr.String() != match {
		t.Fatalf("the match address should be %s, not %s", match, p.Addr.String())
	}

	// remove node
	r.UpdateState(resolver.State{Addresses: resolvedAddrs[:total-2]})
	cstt.Set(test.addresses[:total-2])
	match, _ = cstt.Get(hashElt)

	if _, err := testc.EmptyCall(context.Background(), &testpb.Empty{}, grpc.Peer(&p), grpc.WaitForReady(true)); err != nil {
		t.Fatalf("EmptyCall() = _, %v, want _, <nil>", err)
	}
	t.Logf("[REMOVE] peer=%s, expect=%s", p.Addr.String(), match)
	if p.Addr.String() != match {
		t.Fatalf("the match address should be %s, not %s", match, p.Addr.String())
	}
}

func startTestServers(count int) (_ *test, err error) {
	t := &test{}

	defer func() {
		if err != nil {
			t.cleanup()
		}
	}()
	for i := 0; i < count; i++ {
		lis, err := net.Listen("tcp", "localhost:0")
		if err != nil {
			return nil, fmt.Errorf("failed to listen %v", err)
		}

		s := grpc.NewServer()
		testpb.RegisterTestServiceServer(s, &testServer{})
		t.servers = append(t.servers, s)
		t.addresses = append(t.addresses, lis.Addr().String())

		go func(s *grpc.Server, l net.Listener) {
			s.Serve(l)
		}(s, lis)
	}

	return t, nil
}

type test struct {
	servers   []*grpc.Server
	addresses []string
}

func (t *test) cleanup() {
	for _, s := range t.servers {
		s.Stop()
	}
}

type testServer struct {
	testpb.TestServiceServer
}

func (s *testServer) EmptyCall(ctx context.Context, in *testpb.Empty) (*testpb.Empty, error) {
	return &testpb.Empty{}, nil
}

func (s *testServer) FullDuplexCall(stream testpb.TestService_FullDuplexCallServer) error {
	return nil
}
