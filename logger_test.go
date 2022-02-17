package main

import (
	"errors"
	"testing"
)

// Don't move this test from this lines.
func TestLogCaller(t *testing.T) {
	d := NewDebug()
	have, err := d.SLog("caller", logCaller(1))
	if err != nil {
		t.Fatal(err)
	}

	if want, have := `{"caller":"logger_test.go:11"}`+"\n", have; want != have {
		t.Errorf("\nwant %#v\nhave %#v", want, have)
	}
}

func TestTimeSLog(t *testing.T) {
	d := NewDebug()

	s, err := d.SLog("level", "error", "msg", "<&>")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(s)

	// ---

	d = NewDebug(func(s *Debug) {
		s.timefmt = "2006-01-02 15:04:05"
	})

	s, err = d.SLog("level", "error", "msg", "<&>")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(s)
}

// Run all tests with $ go test -run TestLog

func TestLogger(t *testing.T) {
	t.Parallel()
	debug := NewDebug()

	have, err := debug.SLog("err", errors.New("err"), "m", map[string]int{"0": 0}, "a", []int{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}

	want := `{"a":[1,2,3],"err":"err","m":{"0":0}}` + "\n"

	if want != have {
		t.Errorf("\nwant %#v\nhave %#v", want, have)
	}
}

func TestLoggerMissingValue(t *testing.T) {
	t.Parallel()
	debug := NewDebug()
	have, err := debug.SLog("k")
	if err != nil {
		t.Fatal(err)
	}

	if want, have := `{"k":"(MISSING)"}`+"\n", have; want != have {
		t.Errorf("\nwant %#v\nhave %#v", want, have)
	}
}

func TestLoggerNilStringerKey(t *testing.T) {
	t.Parallel()
	debug := NewDebug()
	have, err := debug.SLog((*stringer)(nil), "v")
	if err != nil {
		t.Fatal(err)
	}

	if want, have := `{"NULL":"v"}`+"\n", have; want != have {
		t.Errorf("\nwant %#v\nhave %#v", want, have)
	}
}

type stringer string

func (s stringer) String() string {
	return string(s)
}

type stringError string

func (s stringError) Error() string {
	return string(s)
}

func TestLoggerNilErrorValue(t *testing.T) {
	t.Parallel()
	debug := NewDebug()
	have, err := debug.SLog("err", (*stringError)(nil))
	if err != nil {
		t.Fatal(err)
	}

	if want, have := `{"err":null}`+"\n", have; want != have {
		t.Errorf("\nwant %#v\nhave %#v", want, have)
	}
}

func TestLoggerNoHTMLEscape(t *testing.T) {
	t.Parallel()
	debug := NewDebug()
	have, err := debug.SLog("k", "<&>")
	if err != nil {
		t.Fatal(err)
	}

	if want, have := `{"k":"<&>"}`+"\n", have; want != have {
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
		debug := NewDebug()
		have, err := debug.SLog("v", test.v)
		if err != nil {
			t.Fatal(err)
		}

		if want, have := test.expected+"\n", have; want != have {
			t.Errorf("\nwant %#v\nhave %#v", want, have)
		}
	}
}
