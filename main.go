package main

import (
	"github.com/containerd/cgroups"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

func main() {
	shares := uint64(1)
	_, err := cgroups.New(cgroups.V1, cgroups.StaticPath("/magic"), &specs.LinuxResources{
		CPU: &specs.LinuxCPU{
			Shares: &shares,
		},
	})
	if err != nil {
		panic(err)
	}
}
