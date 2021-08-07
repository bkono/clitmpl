package servercmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/bkono/clitmpl/cmd/rootcmd"
	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
)

// Config for the root command, including values that should be available to all subcommands.
type Config struct {
	rootCfg *rootcmd.Config
	port    int
}

// RegisterFlags registers the flag fields into the provided flag.FlagSet.
func (c *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.IntVar(&c.port, "port", 5544, "port to listen on")
}

// Exec function for this command.
func (c *Config) Exec(context.Context, []string) error {
	fmt.Printf("server: %+v\n", c)

	// Implement server logic here
	return flag.ErrHelp
}

// New constructs a usable ffcli.Command and an empty Config that will be filled in after parsing.
func New(rootCfg *rootcmd.Config) *ffcli.Command {
	cfg := Config{rootCfg: rootCfg}

	fs := flag.NewFlagSet("clitmpl server", flag.ExitOnError)
	cfg.RegisterFlags(fs)
	rootCfg.RegisterFlags(fs) // note, this allows the root flags to avoid being positional. they will work before or after the subcommand

	return &ffcli.Command{
		Name:       "server",
		ShortUsage: "clitmpl server [flags]",
		FlagSet:    fs,
		Options:    []ff.Option{ff.WithEnvVarPrefix("CLITMPL")},
		Exec:       cfg.Exec,
	}
}
