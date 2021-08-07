package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bkono/clitmpl/cmd/rootcmd"
	"github.com/bkono/clitmpl/cmd/servercmd"
	"github.com/peterbourgon/ff/v3/ffcli"
)

func main() {
	var (
		rootCommand, rootCfg = rootcmd.New()
		serverCommand        = servercmd.New(rootCfg)
	)

	// Create and add subcommands here, prior to parsing

	rootCommand.Subcommands = []*ffcli.Command{serverCommand}

	if err := rootCommand.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error during Parse: %v\n", err)
		os.Exit(1)
	}

	finishConfig(rootCfg)

	if err := rootCommand.Run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func finishConfig(cfg *rootcmd.Config) {
	// Do client initialization here, after the config is parsed
	fmt.Printf("%+v\n", cfg)
	return
}
