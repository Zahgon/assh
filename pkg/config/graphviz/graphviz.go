package graphviz

import (
	"moul.io/assh/v2/pkg/config"
)

func nodename(input string) string { _ = "STUB: not implemented"; return "" }

// GraphSettings are used to change the Graph() function behavior.
type GraphSettings struct {
	ShowIsolatedHosts bool
	NoResolveWildcard bool
	NoInherits        bool
}

// Graph computes and returns a dot-compatible graph representation of the config.
func Graph(cfg *config.Config, settings *GraphSettings) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
