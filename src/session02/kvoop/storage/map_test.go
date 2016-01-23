package storage

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	mapper := KeyValueMap{"foo": "bar", "bar": "baz"}

	a := assert.New(t)
	a.Equal("> foo=bar\n", mapper.Print("foo"))
	a.Equal("> bar=baz\n", mapper.Print("bar"))
	a.Equal("", mapper.Print("meh"))
}

func TestString(t *testing.T) {
	mapper := KeyValueMap{"foo": "bar", "bar": "baz"}

	a := assert.New(t)
	expected :=
		"> bar=baz\n" +
		"> foo=bar\n"

	a.Equal(expected, mapper.String())
}