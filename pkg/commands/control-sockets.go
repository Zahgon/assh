package commands

import (
	"github.com/spf13/cobra"
)

var socketsCommand = &cobra.Command{
	Use:   "sockets",
	Short: "Manage control sockets",
}

var listSocketsCommand = &cobra.Command{
	Use:   "list",
	Short: "List active control sockets",
	RunE:  runListSocketsCommand,
}

var flushSocketsCommand = &cobra.Command{
	Use:   "flush",
	Short: "Close control sockets",
	RunE:  runFlushSocketsCommand,
}

var masterSocketCommand = &cobra.Command{
	Use:   "master",
	Short: "Open a master control socket",
	RunE:  runMasterSocketCommand,
}

// nolint:gochecknoinits
func init() {
	socketsCommand.AddCommand(listSocketsCommand)
	socketsCommand.AddCommand(flushSocketsCommand)
	socketsCommand.AddCommand(masterSocketCommand)
}

func runListSocketsCommand(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func runMasterSocketCommand(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// #nosec

func runFlushSocketsCommand(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
