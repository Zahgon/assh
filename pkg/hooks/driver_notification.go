//go:build !openbsd && !freebsd && !netbsd && !windows
// +build !openbsd,!freebsd,!netbsd,!windows

package hooks

// NotificationDriver is a driver that notifications some texts to the terminal
type NotificationDriver struct {
	line string
}

// NewNotificationDriver returns a NotificationDriver instance
func NewNotificationDriver(line string) (NotificationDriver, error) {
	_ = "STUB: not implemented"
	return *new(NotificationDriver), nil
}

// Run notifications a line to the terminal
func (d NotificationDriver) Run(args RunArgs) error { _ = "STUB: not implemented"; return nil }

// Close is mandatory for the interface, here it does nothing
func (d NotificationDriver) Close() error { _ = "STUB: not implemented"; return nil }
