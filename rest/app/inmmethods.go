package app

import "fmt"

var dbMap map[string]string

func (i *InMemoryDatabase) connect() error {
	dbMap = map[string]string{}
	return nil
}

func (i *InMemoryDatabase) create(shortUrl, fullUrl string) error {
	if !IsUrl(shortUrl) || !IsUrl(fullUrl) {
		return fmt.Errorf("переданы некорректные url")
	}
	dbMap[shortUrl] = fullUrl
	return nil
}

func (i *InMemoryDatabase) get(url string) (string, error) {
	return dbMap[url], nil
}

func (i *InMemoryDatabase) closeConnect() error {
	dbMap = nil
	return nil
}