package go_helpers

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestStringNoSpaces(t *testing.T) {

	assert.Equal(t, StringNoSpaces(`some  some 1      some`), "some some 1 some")

}
