package main

import (
	"encoding/json"
	"net/http"
)

const (
	msgError = "error"
	msgOK    = "ok"
)

type Response struct {
	MessageOK    *MessageOK    `json:"message_ok,omitempty"`
	MessageError *MessageError `json:"message_error,omitempty"`
	Data         interface{}   `json:"data,omitempty"`
}

type MessageOK struct {
	Content string `json:"content"`
}

type MessageError struct {
	Content string `json:"content"`
}

// newResponse return standard response depending of type message.
func newResponse(msgType, content string, data interface{}) Response {
	var resp Response

	switch msgType {
	case msgOK:
		resp = Response{
			MessageOK: &MessageOK{
				Content: content,
			},
			MessageError: nil,
			Data:         data,
		}
	case msgError:
		resp = Response{
			MessageOK: nil,
			MessageError: &MessageError{
				Content: content,
			},
			Data: data,
		}
	}

	return resp
}

// newResponseOK is for generate a JSON response body, e.g.:
//
//     {
//          "message_ok": {
//               "content": "resource updated"
//          }
//     }
func newResponseOK(content string, data interface{}) Response {
	return newResponse(msgOK, content, data)
}

// newResponseError is for generate a JSON response body, e.g.:
//
//     {
//          "message_error": {
//               "content": "error path"
//          }
//     }
func newResponseError(content string, data interface{}) Response {
	return newResponse(msgError, content, data)
}

func rJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type HandleFunc func(http.ResponseWriter, *http.Request)

// mGET allows to pass a request only with the GET method.
func mGET(hf HandleFunc) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			methodNotAllowed(w, r)
		} else {
			hf(w, r)
		}
	}
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	rJSON(w, http.StatusMethodNotAllowed, newResponseError("method not allowed", nil))
}
