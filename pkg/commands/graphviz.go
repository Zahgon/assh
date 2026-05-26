package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var graphvizConfigCommand = &cobra.Command{
	Use:   "graphviz",
	Short: "Generate a Graphviz graph of the hosts",
	RunE:  runGraphvizConfigCommand,
}

// nolint:gochecknoinits
func init() {
	graphvizConfigCommand.Flags().BoolP("show-isolated-hosts", "", false, "Show isolated hosts")
	graphvizConfigCommand.Flags().BoolP("no-resolve-wildcard", "", false, "Do not resolve wildcards in Gateways")
	graphvizConfigCommand.Flags().BoolP("no-inheritance-links", "", false, "Do not show inheritance links")
	_ = viper.BindPFlags(graphvizConfigCommand.Flags())
}

func runGraphvizConfigCommand(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
