//go:build !openbsd && !freebsd && !netbsd
// +build !openbsd,!freebsd,!netbsd

package logger

import (
	"go.uber.org/zap/zapcore"
)

// LogLevelFromParentSSHProcess inspects parent `ssh` process for eventual passed `-v` flags.
func LogLevelFromParentSSHProcess() (zapcore.Level, error) {
	_ = "STUB: not implemented"
	// FIXME: check if parent process is `ssh`
	return *new(zapcore.Level), nil
}
