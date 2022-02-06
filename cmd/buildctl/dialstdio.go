package main

import (
	"time"

	"github.com/moby/buildkit/cmd/buildctl/dial"
	"github.com/urfave/cli"
)

var dialStdioCommand = cli.Command{
	Name:   "dial-stdio",
	Usage:  "Proxy the stdio stream to the daemon connection. Should not be invoked manually.",
	Hidden: true,
	Action: dialStdioAction,
}

func dialStdioAction(clicontext *cli.Context) error {
	addr := clicontext.GlobalString("addr")
	timeout := time.Duration(clicontext.GlobalInt("timeout")) * time.Second
	return dial.DialStdio(addr, timeout)
}
