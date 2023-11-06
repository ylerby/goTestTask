package app

import "fmt"

type InMemoryDatabase struct {
	dbMap map[string]string
}

func (i *InMemoryDatabase) connect() error {
	i.dbMap = map[string]string{}
	return nil
}

func (i *InMemoryDatabase) getShortUrl(url string) (string, error) {
	value, ok := i.dbMap[url]
	if !ok {
		return "", fmt.Errorf("значение не найдено")
	}
	return value, nil
}

func (i *InMemoryDatabase) create(shortUrl, fullUrl string) error {
	if !IsUrl(shortUrl) || !IsUrl(fullUrl) {
		return fmt.Errorf("переданы некорректные url")
	}
	_, ok := i.dbMap[fullUrl]
	if ok {
		return fmt.Errorf("сокращенный url для данного url уже существует")
	}

	i.dbMap[fullUrl] = shortUrl
	return nil
}
