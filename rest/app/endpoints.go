package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

type DbInterface interface {
	connect() error
	get(url string) (string, error)
	create(shortUrl, fullUrl string) error
	closeConnect() error
}

type SqlDatabase struct{}

type InMemoryDatabase struct{}

func IsUrl(urlAddress string) bool {
	u, err := url.Parse(urlAddress)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func getUrl(c *gin.Context) {
	logger.Println("запрос на получение полного url-адреса")

	var urlRequest struct {
		shortUrl string
	}

	c.Bind(&urlRequest)

	fullUrl, err := database.get(urlRequest.shortUrl)
	if err != nil {
		logger.Println("ошибка при получении полного url-адреса")
	}

	c.String(http.StatusOK, fullUrl)
}

func postUrl(c *gin.Context) {
	logger.Println("запрос на получение сокращенного url-адреса")

	var urlRequest struct {
		fullUrl string
	}

	c.Bind(&urlRequest)

	shortUrl, err := makeShortUrl(urlRequest.fullUrl)
	if err != nil {
		logger.Println("ошибка при получении сокращенного url-адреса")
	}

	c.String(http.StatusOK, shortUrl)
}

func makeShortUrl(fullUrl string) (string, error) {
	return "", nil
}
