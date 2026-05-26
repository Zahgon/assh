package config

import "github.com/spf13/pflag"

var (
	// SSHBoolFlags contains list of available SSH boolean options
	SSHBoolFlags = []string{"1", "2", "4", "6", "A", "a", "C", "f", "G", "g", "K", "k", "M", "N", "n", "q", "s", "T", "t", "V", "v", "X", "x", "Y", "y"}
	// SSHStringFlags contains list of available SSH string options
	SSHStringFlags = []string{"b", "c", "D", "E", "e", "F", "I", "i", "L", "l", "m", "O", "o", "p", "Q", "R", "S", "W", "w"}
)

// SSHFlags contains cobra string and bool flags for SSH
func SSHFlags() *pflag.FlagSet { _ = "STUB: not implemented"; return nil }

// Populate SSHFlags
// FIXME: support count flags (-vvv == -v -v -v)
// FIXME: support joined bool flags (-it == -i -t)
