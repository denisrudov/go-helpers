package go_helpers

import (
	"strconv"
	"strings"
)

type JsonInt int

func replaceQuotes(in []byte) string {
	return strings.Replace(string(in), `"`, "", -1)
}

// Unmarshal Int Value
func (i *JsonInt) UnmarshalJSON(in []byte) error {

	replaced := replaceQuotes(in)
	var err error

	if intValue, err := strconv.Atoi(replaced); err == nil {
		*i = JsonInt(intValue)
	}

	return err

}

type JsonFloat float64

// Unmarshal Float value
func (f *JsonFloat) UnmarshalJSON(in []byte) error {
	replaced := replaceQuotes(in)

	var err error

	if float, err := strconv.ParseFloat(replaced, 64); err == nil {
		*f = JsonFloat(float)
	}

	return err
}
