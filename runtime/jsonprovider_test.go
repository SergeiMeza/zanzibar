package zanzibar

import (
	"testing"
	"io"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const simpleJSON = `{"foo.bar": 1, "foo.baz": "foobaz"}`

func TestJSONSimple(t *testing.T) {
	jsonBuf := new(bytes.Buffer)
	jsonBuf.WriteString(simpleJSON)

	jsonProvider, err := newJSONProvider("", []io.Reader{jsonBuf})
	require.NoError(t, err)

	t.Logf("%+v", jsonProvider)

	fooBarValue := jsonProvider.Get("foo.bar")
	var fooBarInt int64
	err = fooBarValue.Populate(&fooBarInt)
	require.NoError(t, err)
	assert.Equal(t, int64(1), fooBarInt)
}


func TestMustGetStringHappy(t *testing.T) {
	jsonBuf := new(bytes.Buffer)
	jsonBuf.WriteString(simpleJSON)

	jsonProvider, err := newJSONProvider("", []io.Reader{jsonBuf})
	require.NoError(t, err)

	t.Logf("%+v", jsonProvider)

	fooBaz := jsonProvider.MustGetString("foo.baz")
	assert.Equal(t, "foobaz", fooBaz)
}