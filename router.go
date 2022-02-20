package main

import (
	"net/http"
	"regexp"
	"strings"
)

// router handler GET: /
func router(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if regexp.MustCompile(`(?i)([a-zA-Z]+)`).MatchString(path) {
		path := strings.ToLower(path)

		// get all currencies
		if path == "/v1" || path == "/v1/" {
			curs, err := getAll()
			if err != nil {
				resp := newResponseError(err.Error(), nil)
				rJSON(w, http.StatusServiceUnavailable, resp)
				return
			}

			rJSON(w, http.StatusOK, newResponseOK("", curs))
			return
		}

		if path == "/v1/euro" || path == "/v1/euro/" {
			getOne(w, 0)
			return
		}

		if path == "/v1/yuan" || path == "/v1/yuan/" {
			getOne(w, 1)
			return
		}

		if path == "/v1/lira" || path == "/v1/lira/" {
			getOne(w, 2)
			return
		}

		if path == "/v1/ruble" || path == "/v1/ruble/" {
			getOne(w, 3)
			return
		}

		if path == "/v1/dollar" || path == "/v1/dollar/" {
			getOne(w, 4)
			return
		}

		resp := newResponseError("path error", nil)
		rJSON(w, http.StatusNotFound, resp)
		return
	}

	resp := newResponseError("path error", nil)
	rJSON(w, http.StatusNotFound, resp)
}

func getOne(w http.ResponseWriter, key int) {
	cur, err := getUnique(key)
	if err != nil {
		resp := newResponseError(err.Error(), nil)
		rJSON(w, http.StatusServiceUnavailable, resp)
		return
	}

	rJSON(w, http.StatusOK, newResponseOK("", cur))
}
