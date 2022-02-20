package main

import (
	"bytes"
	"encoding"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Logger interface {
	Log(keyvals ...interface{}) error
}

// Debug is a JSON logger.
type Debug struct {
	timefmt string
}

// NewDebug provides a simple JSON logger copied from https://github.com/go-kit/log/blob/main/json_logger.go
// but with minor modifications:
//
//     d := NewDebug()
//     d.Log("level", "error", "msg", "error message description")
//
// Output:
//
//     {"level":"error","msg":"error message description"}
//
// Enable time setting its format:
//
//     timefmt := func(s *Debug) {
//          s.timefmt = "2006-01-02 15:04:05"
//     }
//
//     d := NewDebug(timefmt)
//     d.Log("level", "error", "msg", "error message description")
//
// Output:
//
//     {"level":"error","msg":"error message description","time":"2022-02-16 00:46:46"}
func NewDebug(opts ...func(*Debug)) *Debug {
	d := &Debug{}
	for _, opt := range opts {
		opt(d)
	}
	return d
}

func (d *Debug) Log(keyvals ...interface{}) error {
	s, err := d.SLog(keyvals...)
	if err != nil {
		return err
	}
	fmt.Print(s)
	return nil
}

func (d *Debug) SLog(keyvals ...interface{}) (string, error) {
	n := (len(keyvals) + 1) / 2 // +1 to handle case when len is odd
	m := make(map[string]interface{}, n)
	for i := 0; i < len(keyvals); i += 2 {
		k := keyvals[i]
		var v interface{} = "(MISSING)"
		if i+1 < len(keyvals) {
			v = keyvals[i+1]
		}
		merge(m, k, v)
	}

	if d.timefmt != "" {
		m["time"] = time.Now().Format(d.timefmt)
	}

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(m)
	return buf.String(), err
}

// logCaller returns a string that returns a file and line from a specified depth
// in the callstack.
func logCaller(depth int) string {
	_, file, line, _ := runtime.Caller(depth)
	idx := strings.LastIndexByte(file, '/')
	// using idx+1 below handles both of following cases:
	// idx == -1 because no "/" was found, or
	// idx >= 0 and we want to start at the character after the found "/".
	return file[idx+1:] + ":" + strconv.Itoa(line)
}

// merge helper for map[string]interface in SLog func.
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

// safeString helper for map[string]interface in SLog func.
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

// safeError helper for map[string]interface in SLog func.
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
