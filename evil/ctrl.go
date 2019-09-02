package evil

import (
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

// RunCtrl ....
type RunCtrl interface {
	Start(*Config) error    // like systemctl start xxx
	Stop(name string) error // like systemctl stop xxx
	Status(name string)     // like systemctl status xxx
}

// Config for an unit
type Config struct {
	Name      string
	Path      string // executable file path
	ExecStart []string
	Res       *specs.LinuxResources
}
