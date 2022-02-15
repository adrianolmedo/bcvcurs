package main

import (
	"bytes"
	"errors"
	"testing"
)

type stringer string

func (s stringer) String() string {
	return string(s)
}

type stringError string

func (s stringError) Error() string {
	return string(s)
}

func TestLogger(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	debug := &Debug{buf}

	if err := debug.Log("err", errors.New("err"), "m", map[string]int{"0": 0}, "a", []int{1, 2, 3}); err != nil {
		t.Fatal(err)
	}
	if want, have := `{"a":[1,2,3],"err":"err","m":{"0":0}}`+"\n", buf.String(); want != have {
		t.Errorf("\nwant %#v\nhave %#v", want, have)
	}
}

func TestLoggerMissingValue(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	debug := &Debug{buf}
	if err := debug.Log("k"); err != nil {
		t.Fatal(err)
	}
	if want, have := `{"k":"(MISSING)"}`+"\n", buf.String(); want != have {
		t.Errorf("\nwant %#v\nhave %#v", want, have)
	}
}

func TestLoggerNilStringerKey(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	debug := &Debug{buf}
	if err := debug.Log((*stringer)(nil), "v"); err != nil {
		t.Fatal(err)
	}
	if want, have := `{"NULL":"v"}`+"\n", buf.String(); want != have {
		t.Errorf("\nwant %#v\nhave %#v", want, have)
	}
}

func TestLoggerNilErrorValue(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	debug := &Debug{buf}
	if err := debug.Log("err", (*stringError)(nil)); err != nil {
		t.Fatal(err)
	}
	if want, have := `{"err":null}`+"\n", buf.String(); want != have {
		t.Errorf("\nwant %#v\nhave %#v", want, have)
	}
}

func TestLoggerNoHTMLEscape(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	debug := &Debug{buf}
	if err := debug.Log("k", "<&>"); err != nil {
		t.Fatal(err)
	}
	if want, have := `{"k":"<&>"}`+"\n", buf.String(); want != have {
		t.Errorf("\nwant %#v\nhave%#v", want, have)
	}
}

// aller implements json.Marshaler, encoding.TextMarshaler, and fmt.Stringer.
type aller struct{}

func (aller) MarshalJSON() ([]byte, error) {
	return []byte("\"json\""), nil
}

func (aller) MarshalText() ([]byte, error) {
	return []byte("text"), nil
}

func (aller) String() string {
	return "string"
}

func (aller) Error() string {
	return "error"
}

// textstringer implements encoding.TextMarshaler and fmt.Stringer.
type textstringer struct{}

func (textstringer) MarshalText() ([]byte, error) {
	return []byte("text"), nil
}

func (textstringer) String() string {
	return "string"
}

func TestLoggerStringValue(t *testing.T) {
	t.Parallel()
	tests := []struct {
		v        interface{}
		expected string
	}{
		{
			v:        aller{},
			expected: `{"v":"json"}`,
		},
		{
			v:        textstringer{},
			expected: `{"v":"text"}`,
		},
		{
			v:        stringer("string"),
			expected: `{"v":"string"}`,
		},
	}

	for _, test := range tests {
		buf := &bytes.Buffer{}
		debug := &Debug{buf}
		if err := debug.Log("v", test.v); err != nil {
			t.Fatal(err)
		}

		if want, have := test.expected+"\n", buf.String(); want != have {
			t.Errorf("\nwant %#v\nhave %#v", want, have)
		}
	}
}
