package commands

import (
	"github.com/spf13/cobra"
	"moul.io/assh/v2/pkg/config"
)

var wrapperCommand = &cobra.Command{
	Use:    "wrapper",
	Short:  "Initialize assh, then run ssh/scp/rsync...",
	Hidden: true,
}

var sshWrapperCommand = &cobra.Command{
	Use:   "ssh",
	Short: "Wrap ssh",
	RunE:  runSSHWrapperCommand,
}

// nolint:gochecknoinits
func init() {
	sshWrapperCommand.Flags().AddFlagSet(config.SSHFlags())
	wrapperCommand.AddCommand(sshWrapperCommand)
}

func runSSHWrapperCommand(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// prepare variables

// check if config is up-to-date

// check if .ssh/config is outdated

// Execute Binary
// #nosec
