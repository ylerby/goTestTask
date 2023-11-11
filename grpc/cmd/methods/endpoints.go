package methods

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	service "grpc/pkg/grpc_service/api/grpc_service"
	"log"
	"net/url"
	"strings"
)

type DbInterface interface {
	Connect() error
	getFullUrl(url string) (string, error)
	create(shortUrl, fullUrl string) error
}

var Database DbInterface

type Server struct {
	service.UnimplementedFullUrlServer
	service.UnimplementedShortUrlServer
}

func IsUrl(urlAddress string) bool {
	u, err := url.Parse(urlAddress)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (s *Server) GetFullUrl(_ context.Context, req *service.GetFullUrlRequest) (*service.GetFullUrlResponse, error) {
	shortUrl := req.GetShortUrl()
	fullUrl, err := Database.getFullUrl(shortUrl)
	if err != nil {
		return nil, err
	}
	return &service.GetFullUrlResponse{FullUrl: fullUrl}, nil
}

func (s *Server) GetShortUrl(_ context.Context, req *service.GetShortUrlRequest) (*service.GetShortUrlResponse, error) {
	fullUrl := req.GetFullUrl()
	shortUrl, err := MakeShortUrl(fullUrl)
	if err != nil {
		return nil, err
	}

	err = Database.create(shortUrl, fullUrl)
	if err != nil {
		return nil, err
	}
	return &service.GetShortUrlResponse{ShortUrl: fullUrl}, nil
}

func MakeShortUrl(fullUrl string) (string, error) {
	if !IsUrl(fullUrl) {
		return "", fmt.Errorf("переданы некорректные url")
	}
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
	hash := sha1.Sum([]byte(fullUrl))
	hashString := hex.EncodeToString(hash[:])

	shortUrl := getUniqueHash(hashString, charSet)

	log.Println("Изначальный url = ", fullUrl)
	log.Println("Сокращенный url = ", shortUrl)
	return shortUrl, nil
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
