package main

import (
	"errors"
	"log"
	"os"
	"time"

	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/wi-cuckoo/fungo/evil"
)

var pid int

func init() {
	pid = os.Getpid()
}

func main() {
	en, err := evil.New("cuckoo")
	if err != nil {
		panic(err.Error())
	}

	yy := []string{"yes-one", "yes-two"}
	for i, y := range yy {
		quota := int64((i + 1) * 10)
		if err := yes(en, y, quota); err != nil {
			log.Fatal(err)
		}
	}

	defer func() {
		for _, y := range yy {
			if err := en.Stop(y); err != nil {
				log.Println(err.Error())
			}
		}
	}()

	time.Sleep(time.Second * 30)
}

func yes(rc evil.RunCtrl, name string, quota int64) error {
	if quota <= 0 || quota > 50 {
		return errors.New("invalid quota (0, 100)")
	}
	quota *= 1000
	limit := int64(1 << 10) // 1M
	cfg := &evil.Config{
		Name:      name,
		ExecStart: []string{"yes", "my lord"},
		Res: &specs.LinuxResources{
			CPU: &specs.LinuxCPU{
				Quota: &quota,
			},
			Memory: &specs.LinuxMemory{
				Limit: &limit,
				// DisableOOMKiller: &oom,
			},
		},
	}
	if err := rc.Start(cfg); err != nil {
		return err
	}

	return nil
}
