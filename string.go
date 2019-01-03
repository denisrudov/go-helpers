package go_helpers

import "strings"

// Removes more than one space from a string
func StringNoSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
