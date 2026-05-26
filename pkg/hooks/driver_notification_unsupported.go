//go:build openbsd || freebsd || netbsd || windows
// +build openbsd freebsd netbsd windows

package hooks

type NotificationDriver struct{}

func NewNotificationDriver(_ string) (NotificationDriver, error) {
	_ = "STUB: not implemented"
	return *new(NotificationDriver), nil
}
func (NotificationDriver) Run(_ RunArgs) error { _ = "STUB: not implemented"; return nil }
func (d NotificationDriver) Close() error      { _ = "STUB: not implemented"; return nil }
