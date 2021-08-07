package rootcmd

import (
	"context"
	"flag"

	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
)

// Config for the root command, including values that should be available to all subcommands.
type Config struct {
	Verbose bool
}

// RegisterFlags registers the flag fields into the provided flag.FlagSet. This
// helper function allows subcommands to register the root flags into their
// flagsets, creating "global" flags that can be passed after any subcommand at
// the commandline.
func (c *Config) RegisterFlags(fs *flag.FlagSet) {
	// fs.StringVar(&c.Token, "token", "", "secret token for object API")
	fs.BoolVar(&c.Verbose, "v", false, "log verbose output")
}

// Exec function for this command.
func (c *Config) Exec(context.Context, []string) error {
	// The root command has no meaning, so if it gets executed,
	// display the usage text to the user instead.
	return flag.ErrHelp
}

// New constructs a usable ffcli.Command and an empty Config that will be filled in after parsing.
func New() (*ffcli.Command, *Config) {
	var cfg Config

	fs := flag.NewFlagSet("objectctl", flag.ExitOnError)
	cfg.RegisterFlags(fs)

	return &ffcli.Command{
		Name:       "clitmpl",
		ShortUsage: "clitmpl [flags] <subcommand> [flags] [<arg>...]",
		FlagSet:    fs,
		Options:    []ff.Option{ff.WithEnvVarPrefix("CLITMPL")},
		Exec:       cfg.Exec,
	}, &cfg
}
