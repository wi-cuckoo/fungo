package consulrr

import (
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
)

const scheme = "consul"

func init() {
	// register consul:///resolver.consul.grpc.io
	resolver.Register(manual.NewBuilderWithScheme(scheme))
}
