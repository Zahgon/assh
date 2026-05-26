package commands

import (
	"github.com/spf13/cobra"
)

var infoCommand = &cobra.Command{
	Use:   "info",
	Short: "Display system-wide information",
	RunE:  runInfoCommand,
}

func runInfoCommand(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// FIXME: print info about connections/running processes
// FIXME: print info about current config file version
