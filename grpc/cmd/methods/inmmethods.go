package methods

import "fmt"

type InMemoryDatabase struct {
	dbMap map[string]string
}

func (in *InMemoryDatabase) Connect() error {
	in.dbMap = map[string]string{}
	return nil
}

func (in *InMemoryDatabase) getFullUrl(url string) (string, error) {
	value, ok := in.dbMap[url]
	if !ok {
		return "", fmt.Errorf("значение не найдено")
	}
	return value, nil
}

func (in *InMemoryDatabase) create(shortUrl, fullUrl string) error {
	in.dbMap[shortUrl] = fullUrl
	return nil
}
