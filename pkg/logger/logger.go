package logger

import "go.uber.org/zap/zapcore"

// MustLogLevel returns a log level based on both user input and parent SSH process
func MustLogLevel(debug, verbose bool) zapcore.Level {
	_ = "STUB: not implemented"
	return *new(zapcore.Level)
}
