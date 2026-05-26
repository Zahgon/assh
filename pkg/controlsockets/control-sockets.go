package controlsockets

import (
	"time"
)

// ControlSocket defines a unix-domain socket controlled by a master SSH process
type ControlSocket struct {
	path        string
	controlPath string
}

// ControlSockets is a list of ControlSocket
type ControlSockets []ControlSocket

func translateControlPath(input string) string { _ = "STUB: not implemented"; return "" }

// LookupControlPathDir returns the ControlSockets in the ControlPath directory
func LookupControlPathDir(controlPath string) (ControlSockets, error) {
	_ = "STUB: not implemented"
	return *new(ControlSockets), nil
}

// Path returns the absolute path of the socket
func (s *ControlSocket) Path() string {
	_ = "STUB: not implemented"

	// RelativePath returns a path relative to the configured ControlPath
	return ""
}

func (s *ControlSocket) RelativePath() string { _ = "STUB: not implemented"; return "" }

// CreatedAt returns the modification time of the sock file
func (s *ControlSocket) CreatedAt() (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// ActiveConnections returns the amount of active connections using a control socket
func (s *ControlSocket) ActiveConnections() (int, error) { _ = "STUB: not implemented"; return 0, nil }
