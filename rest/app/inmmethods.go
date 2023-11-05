package app

import "fmt"

type InMemoryDatabase struct {
	dbMap map[string]string
}

func (i *InMemoryDatabase) connect() error {
	i.dbMap = map[string]string{}
	return nil
}

func (i *InMemoryDatabase) create(shortUrl, fullUrl string) error {
	if !IsUrl(shortUrl) || !IsUrl(fullUrl) {
		return fmt.Errorf("переданы некорректные url")
	}
	i.dbMap[fullUrl] = shortUrl
	return nil
}

func (i *InMemoryDatabase) getShortUrl(url string) (string, error) {
	value, ok := i.dbMap[url]
	if !ok {
		return "", fmt.Errorf("значение не найдено")
	}
	return value, nil
}
