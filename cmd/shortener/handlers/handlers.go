package handlers

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

var shortUrls = make(map[string]string)

func compressString(originalString string) string {
	return base64.StdEncoding.EncodeToString([]byte(originalString))[8:]
}

func mainPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	body, _ := io.ReadAll(r.Body)
	if body != nil {
		originalUrl := string(body)
		shortUrl := compressString(originalUrl)
		host := r.Host
		shortUrls[shortUrl] = originalUrl

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("http://%s/%s", host, shortUrl)))
	}

}

func mainGetHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Path[1:]
	originalUrl := shortUrls[page]
	w.Header().Set("Location", originalUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
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
