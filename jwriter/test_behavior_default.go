//go:build !launchdarkly_easyjson
// +build !launchdarkly_easyjson

package jwriter

// This function tells the writer tests that we shouldn't expect to see hex escape sequences in the output.
// Our default implementation doesn't use them, whereas easyjson does; either way is valid in JSON.
func tokenWriterWillEncodeAsHex(_ rune) bool { //nolint:deadcode,unused // linter is confused
	return false
}
