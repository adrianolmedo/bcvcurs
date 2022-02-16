package main

import (
	"encoding"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"time"
)

var ErrLoggerMissingValue = errors.New("(MISSING)")

type Logger interface {
	Log(keyvals ...interface{}) error
}

// Debug is a simple JSON logger with time by default, e.g.:
// {"time":"2022-02-16T00:44:58.275093472-04:00"}
type Debug struct {
	io.Writer
	timefmt string
}

// NewDebug return a new Logger instance.
//
// Usage:
//
// b := NewDebug(&bytes.Buffer{})
// d.Log("level", "error", "msg", "error message description")
// {"level":"error","msg":"error message description","time":"2022-02-16T00:44:58.275093472-04:00"}
//
// Change time format:
//
// timefmt := func(s *Debug) {
//		s.timefmt = "2006-01-02 15:04:05"
// }
//
// d := NewDebug(&bytes.Buffer{}, timefmt)
// d.Log()
// {"time":"2022-02-16 00:46:46"}
func NewDebug(w io.Writer, opts ...func(*Debug)) Logger {
	d := &Debug{w, ""}
	for _, opt := range opts {
		opt(d)
	}
	return d
}

func (d *Debug) Log(keyvals ...interface{}) error {
	n := (len(keyvals) + 1) / 2 // +1 to handle case when len is odd
	m := make(map[string]interface{}, n)
	for i := 0; i < len(keyvals); i += 2 {
		k := keyvals[i]
		var v interface{} = ErrLoggerMissingValue
		if i+1 < len(keyvals) {
			v = keyvals[i+1]
		}
		merge(m, k, v)
	}

	m["time"] = time.Now()
	if d.timefmt != "" {
		m["time"] = time.Now().Format(d.timefmt)
	}

	enc := json.NewEncoder(d.Writer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(m)
	if err != nil {
		return err
	}

	fmt.Println(d.Writer)
	return nil
}

// merge helper for JSON debugger.
func merge(dst map[string]interface{}, k, v interface{}) {
	var key string
	switch x := k.(type) {
	case string:
		key = x
	case fmt.Stringer:
		key = safeString(x)
	default:
		key = fmt.Sprint(x)
	}

	// We want json.Marshaler and encoding.TextMarshaller to take priority over
	// err.Error() and v.String(). But json.Marshall (called later) does that by
	// default so we force a no-op if it's one of those 2 case.
	switch x := v.(type) {
	case json.Marshaler:
	case encoding.TextMarshaler:
	case error:
		v = safeError(x)
	case fmt.Stringer:
		v = safeString(x)
	}

	dst[key] = v
}

// safeString helper for debug JSON output.
func safeString(str fmt.Stringer) (s string) {
	defer func() {
		if panicVal := recover(); panicVal != nil {
			if v := reflect.ValueOf(str); v.Kind() == reflect.Ptr && v.IsNil() {
				s = "NULL"
			} else {
				panic(panicVal)
			}
		}
	}()
	s = str.String()
	return
}

// safeError helper for debug JSON output.
func safeError(err error) (s interface{}) {
	defer func() {
		if panicVal := recover(); panicVal != nil {
			if v := reflect.ValueOf(err); v.Kind() == reflect.Ptr && v.IsNil() {
				s = nil
			} else {
				panic(panicVal)
			}
		}
	}()
	s = err.Error()
	return
}
