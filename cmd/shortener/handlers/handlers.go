package handlers

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var shortUrls = make(map[string]string)

func compressString(originalString string) string {
	return base64.StdEncoding.EncodeToString([]byte(originalString))[8:]
}

func mainPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	body, _ := io.ReadAll(r.Body)
	if body != nil {
		originalURL := string(body)
		shortURL := compressString(originalURL)
		host := r.Host
		shortUrls[shortURL] = originalURL

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("http://%s/%s", host, shortURL)))
	}

}

func mainGetHandler(w http.ResponseWriter, r *http.Request) {
	firstPathSegment := getFirstPathSegment(r.URL)
	originalURL := shortUrls[firstPathSegment]
	w.Header().Set("Location", originalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func getFirstPathSegment(url *url.URL) string {
	return url.Path[1:]
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mainGetHandler(w, r)
	case http.MethodPost:
		mainPostHandler(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
