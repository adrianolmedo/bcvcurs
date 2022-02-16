package main

import (
	"errors"
	"net/http"
)

const URL = "http://www.bcv.org.ve/"

// dollar handler GET: /v1/dollar
func dollar(logger Logger) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := bodyFromURL(URL)
		if err != nil {
			logger.Log("level", "error", "msg", err.Error(), "path", r.URL.Path)
			resp := newResponse(msgError, err.Error(), nil)
			rJSON(w, http.StatusBadGateway, resp)
			return
		}
		defer body.Close()

		value, err := rateBCV("dolar", body)
		if errors.Is(err, ErrCurrencyNotFound) {
			logger.Log("level", "error", "msg", err.Error(), "path", r.URL.Path)
			resp := newResponse(msgError, err.Error(), nil)
			rJSON(w, http.StatusOK, resp)
			return
		}

		if err != nil {
			logger.Log("level", "error", "msg", err.Error(), "path", r.URL.Path)
			resp := newResponse(msgError, "the gopher is eating a cable", nil)
			rJSON(w, http.StatusInternalServerError, resp)
			return
		}

		money := Money{
			Value:  value,
			Iso:    "USD",
			Symbol: "$",
		}

		resp := newResponse(msgOK, "", money)
		rJSON(w, http.StatusOK, resp)
	}
}

/*func foo() map[string]string{
	return map[string]string{
		"dollar": "dolar",
		"juan": "yuan",
	}
}
//f := foo()
//v, ok := f["dollar"]
*/

// home handler GET: /v1/
func home() HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := newResponse(msgOK, "home", nil)
		rJSON(w, http.StatusOK, resp)
	}
}
