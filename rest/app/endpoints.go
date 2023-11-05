package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

type DbInterface interface {
	connect() error
	getShortUrl(url string) (string, error)
	create(shortUrl, fullUrl string) error
}

func IsUrl(urlAddress string) bool {
	u, err := url.Parse(urlAddress)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func getUrl(c *gin.Context) {
	logger.Println("запрос на получение полного url-адреса")

	var urlRequest struct {
		shortUrl string
	}

	err := c.Bind(&urlRequest)
	if err != nil {
		logger.Println("ошибка при получении url-адреса из тела запроса")
	}

	fullUrl, err := database.getShortUrl(urlRequest.shortUrl)
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

	err := c.Bind(&urlRequest)
	if err != nil {
		logger.Println("ошибка при получении url-адреса из тела запроса")
	}

	shortUrl, err := makeShortUrl(urlRequest.fullUrl)
	if err != nil {
		logger.Println("ошибка при получении сокращенного url-адреса")
	}

	c.String(http.StatusOK, shortUrl)
}

// makeShortUrl todo: сделать хеширование url-адреса
func makeShortUrl(fullUrl string) (string, error) {
	return fullUrl, nil
}
