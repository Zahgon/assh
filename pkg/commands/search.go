package commands

import (
	"github.com/spf13/cobra"
)

var searchConfigCommand = &cobra.Command{
	Use:   "search",
	Short: "Search entries by given search text",
	RunE:  searchConfig,
}

func searchConfig(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }
