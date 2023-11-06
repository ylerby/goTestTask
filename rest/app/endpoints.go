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

type DbInterface interface {
	connect() error
	getFullUrl(url string) (string, error)
	create(shortUrl, fullUrl string) error
}

func IsUrl(urlAddress string) bool {
	logger.Println(urlAddress)
	u, err := url.Parse(urlAddress)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func getUrl(c *gin.Context) {
	logger.Println("запрос на получение полного url-адреса")

	var urlRequest struct {
		ShortUrl string `json:"short_url"`
	}

	err := c.BindJSON(&urlRequest)
	if err != nil {
		logger.Println("ошибка при получении url-адреса из тела запроса")
	}
	fullUrl, err := database.getFullUrl(urlRequest.ShortUrl)
	if err != nil {
		logger.Println(err)
		c.String(http.StatusOK, "значение не найдено")
	}

	c.String(http.StatusOK, fullUrl)
}

func postUrl(c *gin.Context) {
	logger.Println("запрос на получение сокращенного url-адреса")

	var urlRequest struct {
		FullUrl string `json:"full_url"`
	}

	err := c.BindJSON(&urlRequest)
	if err != nil {
		logger.Println("ошибка при получении url-адреса из тела запроса")
	}

	shortUrl := makeShortUrl(urlRequest.FullUrl)
	if err != nil {
		logger.Println("ошибка при создании сокращенного url-адреса")
	}

	err = database.create(shortUrl, urlRequest.FullUrl)
	if err != nil {
		logger.Println(err)
	}

	c.String(http.StatusOK, shortUrl)
}

func makeShortUrl(fullUrl string) string {
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
	hash := sha1.Sum([]byte(fullUrl))
	hashString := hex.EncodeToString(hash[:])

	shortUrl := getUniqueHash(hashString, charSet)

	logger.Println("Изначальный url = ", fullUrl)
	logger.Println("Сокращенный url = ", shortUrl)
	return shortUrl
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
