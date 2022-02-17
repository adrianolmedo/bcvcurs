package main

import (
	"net/http"
	"regexp"
	"strings"
)

const URL = "http://www.bcv.org.ve/"

// router handler GET: /
func router(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if regexp.MustCompile(`(?i)([a-zA-Z]+)`).MatchString(path) {
		path := strings.ToLower(path)

		// get all currencies
		if path == "/v1" || path == "/v1/" {
			curs, err := getAll()
			if err != nil {
				resp := newResponse(msgError, err.Error(), nil)
				rJSON(w, http.StatusServiceUnavailable, resp)
				return
			}

			resp := newResponse(msgOK, "", curs)
			rJSON(w, http.StatusOK, resp)
			return
		}

		// get euro
		if path == "/v1/euro" || path == "/v1/euro/" {
			getOne(w, 0)
			return
		}

		// get yuan
		if path == "/v1/yuan" || path == "/v1/yuan/" {
			getOne(w, 1)
			return
		}

		// get lira
		if path == "/v1/lira" || path == "/v1/lira/" {
			getOne(w, 2)
			return
		}

		// get rublo
		if path == "/v1/rublo" || path == "/v1/rublo/" {
			getOne(w, 3)
			return
		}

		// get dollar
		if path == "/v1/dollar" || path == "/v1/dollar/" {
			getOne(w, 4)
			return
		}

		resp := newResponse(msgOK, "path error", nil)
		rJSON(w, http.StatusNotFound, resp)
		return
	}

	resp := newResponse(msgOK, "path error", nil)
	rJSON(w, http.StatusNotFound, resp)
}

func getOne(w http.ResponseWriter, key int) {
	cur, err := getUnique(key)
	if err != nil {
		resp := newResponse(msgError, err.Error(), nil)
		rJSON(w, http.StatusServiceUnavailable, resp)
		return
	}

	resp := newResponse(msgOK, "", cur)
	rJSON(w, http.StatusOK, resp)
}
