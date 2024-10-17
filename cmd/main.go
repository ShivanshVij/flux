package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/loopholelabs/cmdutils/pkg/command"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	return Cmd.Execute(ctx, command.Noninteractive)
}
