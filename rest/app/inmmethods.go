package app

import "fmt"

type InMemoryDatabase struct {
	dbMap map[string]string
}

func (i *InMemoryDatabase) connect() error {
	i.dbMap = map[string]string{}
	return nil
}

func (i *InMemoryDatabase) getFullUrl(url string) (string, error) {
	value, ok := i.dbMap[url]
	logger.Println(i.dbMap)
	if !ok {
		return "", fmt.Errorf("значение не найдено")
	}
	return value, nil
}

func (i *InMemoryDatabase) create(shortUrl, fullUrl string) error {
	_, ok := i.dbMap[shortUrl]
	if ok {
		return fmt.Errorf("сокращенный url уже существует")
	}

	i.dbMap[shortUrl] = fullUrl
	return nil
}
