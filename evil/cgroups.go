package evil

import (
	"errors"
	"fmt"
	"os/exec"
	"sync"
	"syscall"

	"github.com/containerd/cgroups"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type group struct {
	mux  sync.Mutex
	root cgroups.Cgroup
	subs map[string]cgroups.Cgroup
}

// New ...
func New(name string) (RunCtrl, error) {
	path := fmt.Sprintf("/%s", name)
	root, err := cgroups.Load(cgroups.V1, cgroups.StaticPath(path))
	if err != nil {
		root, err = cgroups.New(cgroups.V1, cgroups.StaticPath(path), &specs.LinuxResources{})
		if err != nil {
			return nil, err
		}
	}
	return &group{
		root: root,
		subs: map[string]cgroups.Cgroup{},
	}, nil
}

func (g *group) Start(c *Config) error {
	if len(c.ExecStart) == 0 {
		return errors.New("invalid exec")
	}
	exe := exec.Command(c.ExecStart[0], c.ExecStart[1:]...)
	if err := exe.Start(); err != nil {
		return err
	}

	// create sub group
	ctrl, err := g.root.New(c.Name, c.Res)
	if err != nil {
		return err
	}

	g.subs[c.Name] = ctrl

	return ctrl.AddTask(cgroups.Process{Pid: exe.Process.Pid})
}

func (g *group) Stop(name string) error {
	ctrl, ok := g.subs[name]
	if !ok {
		return errors.New("not found")
	}
	if procs, err := ctrl.Processes(cgroups.Cpu, false); err == nil {
		for _, p := range procs {
			if e := syscall.Kill(p.Pid, syscall.SIGQUIT); e != nil {
				syscall.Kill(p.Pid, syscall.SIGKILL) // kill -9 pid
			}

		}
	}
	return ctrl.Delete()
}

func (g *group) Status(name string) {

}
