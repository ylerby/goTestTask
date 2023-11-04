package app

import "fmt"

func (s *SqlDatabase) connect() error {
	return nil
}

func (s *SqlDatabase) get(url string) (string, error) {
	return "", nil
}

func (s *SqlDatabase) create(shortUrl, fullUrl string) error {
	if !IsUrl(shortUrl) || !IsUrl(fullUrl) {
		return fmt.Errorf("переданы некорректные url")
	}
	dbMap[shortUrl] = fullUrl
	return nil
}

func (s *SqlDatabase) closeConnect() error {
	logger.Println("соединение успешно закрыто")
	return nil
}
