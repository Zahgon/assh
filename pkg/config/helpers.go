package config

func isDynamicHostname(hostname string) bool { _ = "STUB: not implemented"; return false }

// BoolVal returns a boolean matching a configuration string
func BoolVal(input string) bool { _ = "STUB: not implemented"; return false }

func cleanupValue(input string) string { _ = "STUB: not implemented"; return "" }

// stringComment splits comment strings into <1024 char lines
func stringComment(name, value string) string { _ = "STUB: not implemented"; return "" }

// sliceComment splits comment strings into <1024 char lines
func sliceComment(name string, slice []string) string { _ = "STUB: not implemented"; return "" }

// splitSubN splits a string by length
// from: http://stackoverflow.com/questions/25686109/split-string-by-length-in-golang
func splitSubN(s string, n int) []string { _ = "STUB: not implemented"; return nil }
