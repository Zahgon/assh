package hooks

// WriteDriver is a driver that writes some texts to the terminal
type WriteDriver struct {
	line string
}

// NewWriteDriver returns a WriteDriver instance
func NewWriteDriver(line string) (WriteDriver, error) {
	_ = "STUB: not implemented"
	return *new(WriteDriver), nil
}

// Run writes a line to the terminal
func (d WriteDriver) Run(args RunArgs) error { _ = "STUB: not implemented"; return nil }

// Close is mandatory for the interface, here it does nothing
func (d WriteDriver) Close() error { _ = "STUB: not implemented"; return nil }
