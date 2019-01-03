package go_helpers

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonInt_UnmarshalJSON(t *testing.T) {

	jsonString1 := []byte(`{ "a":"1", "b":2  }`)
	jsonString2 := []byte(`{ "a":"des", "b":2  }`)

	type A struct {
		A JsonInt `json:"a"`
		B JsonInt `json:"b"`
	}

	astruct := new(A)

	err := json.Unmarshal(jsonString1, astruct)

	assert.NoError(t, err)

	assert.Equal(t, JsonInt(1), astruct.A)
	assert.Equal(t, JsonInt(2), astruct.B)

	astruct2 := new(A)

	err = json.Unmarshal(jsonString2, astruct2)

	assert.NoError(t, err)

	assert.Equal(t, JsonInt(0), astruct2.A)
	assert.Equal(t, JsonInt(2), astruct2.B)

}

func TestJsonFloat_UnmarshalJSON(t *testing.T) {

	jsonString1 := []byte(`{ "a":"1.24", "b":2 }`)

	type A struct {
		A JsonFloat `json:"a"`
		B JsonFloat `json:"b"`
	}

	js1 := new(A)

	err := json.Unmarshal(jsonString1, js1)

	assert.NoError(t, err)

	assert.Equal(t, JsonFloat(1.24), js1.A)
	assert.Equal(t, JsonFloat(2), js1.B)

}
