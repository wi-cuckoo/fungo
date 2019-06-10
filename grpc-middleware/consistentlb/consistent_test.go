package consistentlb_test

import (
	"context"
	"fmt"
	"net"
	"testing"

	"google.golang.org/grpc"
	testpb "google.golang.org/grpc/test/grpc_testing"
)

func TestConsistentlb(t *testing.T) {

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
