package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var pingCommand = &cobra.Command{
	Use:   "ping",
	Short: "Send packets to the SSH server and display statistics",
	RunE:  runPingCommand,
}

// nolint:gochecknoinits
func init() {
	pingCommand.Flags().BoolP("no-rewrite", "", false, "Do not automatically rewrite outdated configuration")
	pingCommand.Flags().IntP("port", "p", 0, "SSH destination port")
	pingCommand.Flags().UintP("count", "c", 0, "Stop after sending 'count' packets")
	pingCommand.Flags().Float64P("wait", "i", 1, "Wait 'wait' seconds between sending each packet")
	pingCommand.Flags().BoolP("o", "", false, "Exit successfully after receiving one reply packet")
	pingCommand.Flags().Float64P("waittime", "W", 1, "Time in seconds to wait for a reply for each packet sent")
	_ = viper.BindPFlags(pingCommand.Flags())
}

func runPingCommand(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// fixme: resolve port name

// FIXME: switch on error type

// FIXME: catch Ctrl+C
