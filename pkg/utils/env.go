package utils

// EscapeSpaces escapes all space characters with a backslash (\)
func EscapeSpaces(s string) string { _ = "STUB: not implemented"; return "" }

// GetHomeDir returns '~' as a path
func GetHomeDir() string { _ = "STUB: not implemented"; return "" }

// ExpandEnvSafe replaces ${var} or $var in the string according to the values
// of the current environment variables.
// As opposed to os.ExpandEnv, ExpandEnvSafe won't remove the dollar in '$(...)'
// See https://golang.org/src/os/env.go?s=963:994#L22 for the original function
func ExpandEnvSafe(s string) string { _ = "STUB: not implemented"; return "" }

// the following line is the only one changing

// ExpandUser expands tild and env vars in unix paths
func ExpandUser(path string) (string, error) {
	_ = "STUB: not implemented"
	// Expand variables
	return "", nil
}

// OS-agnostic slashes

// ExpandField expands environment variables in field
func ExpandField(input string) string { _ = "STUB: not implemented"; return "" }

// ExpandSliceField expands environment variables in every entries of a slice field
func ExpandSliceField(input []string) []string { _ = "STUB: not implemented"; return nil }
