package app

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
)

// todo: Docker-образ
// todo: Unit-тесты
// todo: grpc

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

	shortUrl := makeShortUrl(urlRequest.fullUrl)
	if err != nil {
		logger.Println("ошибка при создании сокращенного url-адреса")
	}

	err = database.create(shortUrl, urlRequest.fullUrl)
	if err != nil {
		logger.Println(err)
	}

	c.String(http.StatusOK, shortUrl)
}

// makeShortUrl todo: сделать хеширование url-адреса
func makeShortUrl(fullUrl string) string {
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
	hash := sha1.Sum([]byte(fullUrl))
	hashString := hex.EncodeToString(hash[:])

	uniqueHash := getUniqueHash(hashString, charSet)

	return uniqueHash
}

func getUniqueHash(hash string, charSet string) string {
	var uniqueHash strings.Builder

	for _, c := range hash {
		if strings.ContainsRune(charSet, c) {
			uniqueHash.WriteRune(c)
			if uniqueHash.Len() == 10 {
				break
			}
		}
	}

	return "http://my_service/" + uniqueHash.String()
}
