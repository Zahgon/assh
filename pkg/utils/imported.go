// This file contains imported functions
// The license and copyright is reported for each functions in the comments.

package utils

// Imported and unmodified from https://golang.org/src/os/env.go
// Function under the BSD-License - Copyrighted by the Go Authors
// isShellSpecialVar reports whether the character identifies a special
// shell variable such as $*.
func isShellSpecialVar(c uint8) bool { _ = "STUB: not implemented"; return false }

// Imported and unmodified from https://golang.org/src/os/env.go
// Function under the BSD-License - Copyrighted by the Go Authors
// isAlphaNum reports whether the byte is an ASCII letter, number, or underscore
func isAlphaNum(c uint8) bool { _ = "STUB: not implemented"; return false }

// Imported and unmodified from https://golang.org/src/os/env.go
// Function under the BSD-License - Copyrighted by the Go Authors
// getShellName returns the name that begins the string and the number of bytes
// consumed to extract it.  If the name is enclosed in {}, it's part of a ${}
// expansion and two more bytes are needed than the length of the name.
func getShellName(s string) (string, int) { _ = "STUB: not implemented"; return "", 0 }

// Scan to closing brace

// Bad syntax; just eat the brace.

// Scan alphanumerics.
