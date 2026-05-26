package hooks

import (
	"os/exec"
)

// DaemonDriver is a driver that daemons some texts to the terminal
type DaemonDriver struct {
	line string
	cmd  *exec.Cmd
}

// NewDaemonDriver returns a DaemonDriver instance
func NewDaemonDriver(line string) (DaemonDriver, error) {
	_ = "STUB: not implemented"
	return *new(DaemonDriver), nil
}

// Run daemons a line to the terminal
func (d DaemonDriver) Run(args RunArgs) error { _ = "STUB: not implemented"; return nil }

// #nosec

// Close closes a running command
func (d DaemonDriver) Close() error { _ = "STUB: not implemented"; return nil }
