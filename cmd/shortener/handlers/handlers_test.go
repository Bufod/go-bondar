package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainPostHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name string
		body string
		want want
	}{
		{
			name: "test #1",
			body: "https://practicum.yandex.ru/",
			want: want{
				code:        201,
				response:    "http://example.com/Ly9wcmFjdGljdW0ueWFuZGV4LnJ1Lw==",
				contentType: "text/plain",
			},
		},
	}

	router := SetupRouter()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.body))

			w := httptest.NewRecorder()
			router.ServeHTTP(w, request)

			res := w.Result()
			assert.Equal(t, test.want.code, res.StatusCode)

			defer res.Body.Close()

			resBody, err := io.ReadAll(res.Body)
			// resURL, _ := url.Parse(string(resBody))

			require.NoError(t, err)
			assert.Equal(t, test.want.response, string(resBody))
			assert.Equal(t, test.want.contentType, res.Header.Get("Content-Type"))
			// assert.Equal(t, test.body, shortUrls[getFirstPathSegment(resURL)])
		})
	}
}

func TestMainGetHandler(t *testing.T) {
	type want struct {
		code     int
		location string
	}
	tests := []struct {
		name    string
		segment string
		want    want
	}{
		{
			name:    "test #1",
			segment: "Ly9wcmFjdGljdW0ueWFuZGV4LnJ1Lw==",
			want: want{
				code:     307,
				location: "http://practicum.yandex.ru",
			},
		},
	}

	router := SetupRouter()

	for _, test := range tests {
		shortUrls[test.segment] = test.want.location
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", test.segment), nil)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, request)

			res := w.Result()
			assert.Equal(t, test.want.code, res.StatusCode)

			defer res.Body.Close()

			assert.Equal(t, test.want.location, res.Header.Get("Location"))
		})
	}
}

func TestCompressString(t *testing.T) {
	tests := []struct {
		original string
		wants    string
	}{
		{
			original: "https://practicum.yandex.ru/",
			wants:    "Ly9wcmFjdGljdW0ueWFuZGV4LnJ1Lw==",
		},
		{
			original: "http://test.ru/",
			wants:    "L3Rlc3QucnUv",
		},
	}

	for index, test := range tests {
		t.Run("compressString "+string(rune(index)), func(t *testing.T) {
			assert.Equal(t, test.wants, compressString(test.original))
		})
	}
}
