package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/urfave/cli"
)

const file = "~/.poems"

var flags = []cli.Flag{}

func main() {
	app := cli.NewApp()
	app.Name = "leo"
	app.Usage = "I'm a poet"
	app.Action = pick

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
}

func pick(c *cli.Context) {
}

func init() {
	rand.Seed(time.Now().Unix())
	fmt.Fprint(os.Stdout, "\033c")
}
