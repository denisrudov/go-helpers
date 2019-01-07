package go_helpers

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Removes more than one space from a string
func StringNoSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// Parses string and found first
// occurrence of an integer in a given string
// "some string 3000 4000 some other text " == int(3000)
func ParseInt(s string) (ret int, err error) {
	ret = 0
	compile := regexp.MustCompile(`(\d+)`)
	submatch := compile.FindStringSubmatch(s)
	if submatch != nil && len(submatch) > 0 {
		ret, err = strconv.Atoi(submatch[0])
		return
	}
	err = errors.New("no integer found")
	return
}
