package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var buildConfigCommand = &cobra.Command{
	Use:   "build",
	Short: "Build .ssh/config",
	RunE:  runBuildConfigCommand,
}

var buildJSONConfigCommand = &cobra.Command{
	Use:   "json",
	Short: "Returns the JSON output",
	RunE:  runBuildJSONConfigCommand,
}

// nolint:gochecknoinits
func init() {
	buildConfigCommand.Flags().BoolP("no-automatic-rewrite", "", false, "Disable automatic ~/.ssh/config file regeneration")
	buildConfigCommand.Flags().BoolP("expand", "e", false, "Expand all fields")
	buildConfigCommand.Flags().BoolP("ignore-known-hosts", "", false, "Ignore known-hosts file")
	_ = viper.BindPFlags(buildConfigCommand.Flags())

	buildJSONConfigCommand.Flags().BoolP("expand", "e", false, "Expand all fields")
	_ = viper.BindPFlags(buildJSONConfigCommand.Flags())
}

func runBuildConfigCommand(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func runBuildJSONConfigCommand(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
