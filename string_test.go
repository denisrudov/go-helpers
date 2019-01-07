package go_helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringNoSpaces(t *testing.T) {

	assert.Equal(t, StringNoSpaces(`some  some 1      some`), "some some 1 some")

}

func TestParseInt(t *testing.T) {

	tests := []string{
		"3000 some string",
		"some 3000 some string",
		"some 3000 4000 some string",
		"some 3000.42 4000 some string",
	}

	for _, str := range tests {
		ret, err := ParseInt(str)
		assert.NoError(t, err)
		assert.Equal(t, ret, 3000)
	}

	ret, err := ParseInt("some no integer string")
	assert.Error(t, err)
	assert.Equal(t, ret, 0)

}
