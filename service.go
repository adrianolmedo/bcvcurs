package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//type Exchange struct{}

func bodyFromURL(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code from source %d", resp.StatusCode)
	}

	return resp.Body, nil
}

// rateBCV it could be named to Rate() in representation of Repository.
func rateBCV(idAtr string, body io.ReadCloser) (float64, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return 0, err
	}

	result := doc.Find("div[id='" + idAtr + "']").Find("strong").Text()
	if result == "" {
		return 0, ErrCurrencyNotFound
	}

	result = strings.TrimSpace(result)
	return strconv.ParseFloat(result, 64)
}
