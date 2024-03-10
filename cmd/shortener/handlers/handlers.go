package handlers

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var shortUrls = make(map[string]string)

func compressString(originalString string) string {
	return base64.StdEncoding.EncodeToString([]byte(originalString))[8:]
}

func MainPostHandler(c *gin.Context) {
	c.Header("Content-Type", "text/plain")
	body, _ := io.ReadAll(c.Request.Body)
	if body != nil {
		originalURL := string(body)
		shortURL := compressString(originalURL)
		host := c.Request.Host
		shortUrls[shortURL] = originalURL
		c.String(http.StatusCreated, fmt.Sprintf("http://%s/%s", host, shortURL))
	}

}

func MainGetHandler(c *gin.Context) {
	shortUrl := c.Param("shortURL")
	originalURL := shortUrls[shortUrl]
	c.Redirect(http.StatusTemporaryRedirect, originalURL)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/:shortURL", MainGetHandler)
	router.POST("/", MainPostHandler)
	return router
}
