package app

import (
	"github.com/gin-gonic/gin"
)

type DbInterface interface {
	connect() error
	get(url string) (string, error)
	create(shortUrl, fullUrl string) error
	closeConnect() error
}

type SqlDatabase struct{}

type InMemoryDatabase struct{}

func getUrl(c *gin.Context) {
}

func postUrl(c *gin.Context) {
}
