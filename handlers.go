package main

import "net/http"

// home handler GET: /
func home(w http.ResponseWriter, r *http.Request) {
	resp := newResponse(msgOK, "hello world", nil)
	JSON(w, http.StatusOK, resp)
}

// dollar handler GET: /dollar
func dollar(w http.ResponseWriter, r *http.Request) {
	resp := newResponse(msgOK, "dollar price", nil)
	JSON(w, http.StatusOK, resp)
}
