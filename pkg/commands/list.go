package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listConfigCommand = &cobra.Command{
	Use:   "list",
	Short: "List all hosts from assh config",
	RunE:  runListConfigCommand,
}

// nolint:gochecknoinits
func init() {
	listConfigCommand.Flags().BoolP("expand", "e", false, "Expand all fields")
	_ = viper.BindPFlags(listConfigCommand.Flags())
}

func runListConfigCommand(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// ansi coloring
