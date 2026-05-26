package hooks

var possibleShells = []string{
	"/bin/sh", "/bin/bash", "/bin/zsh",
	"/usr/bin/sh", "/usr/bin/bash", "/usr/bin/zsh",
	"/usr/local/bin/sh", "/usr/local/bin/bash", "/usr/local/bin/zsh",

	"C:\\Program Files\\Git\\bin\\bash.exe",
	"C:\\Program Files\\Git\\bin\\sh.exe",
	"C:\\Windows\\System32\\bash.exe",
}

// ExecDriver is a driver that execs some texts to the terminal
type ExecDriver struct {
	line string
}

// NewExecDriver returns a ExecDriver instance
func NewExecDriver(line string) (ExecDriver, error) {
	_ = "STUB: not implemented"
	return *new(ExecDriver), nil
}

// Run execs a line to the terminal
func (d ExecDriver) Run(args RunArgs) error { _ = "STUB: not implemented"; return nil }

// #nosec

// Close is mandatory for the interface, here it does nothing
func (d ExecDriver) Close() error { _ = "STUB: not implemented"; return nil }

func renderCommand(line string, tmplArgs RunArgs) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func findAvailableShell(shells []string) string { _ = "STUB: not implemented"; return "" }

func isExecutable(path string) bool { _ = "STUB: not implemented"; return false }

// on Windows, the executable bit is not necessary

// on Linux actually check if the file is executable
