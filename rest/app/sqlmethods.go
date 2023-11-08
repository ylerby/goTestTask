package app

import (
	//"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	model "rest/app/model"
)

type SqlDatabase struct {
	dbSql *gorm.DB
}

func (s *SqlDatabase) connect() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
    os.Getenv("DB_HOST"),
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_NAME"),
    os.Getenv("DB_PORT"))

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  s.dbSql = db
  if err != nil {
    return fmt.Errorf("ошибка при подключении к БД")
  }

  err = s.dbSql.AutoMigrate(&model.SqlModel{})
  if err != nil {
    return fmt.Errorf("ошибка при миграции")
  }

  return nil
}

func (s *SqlDatabase) getFullUrl(url string) (string, error) {    
	var urlRequest model.SqlModel
    	err := s.dbSql.Find(&urlRequest, "short_url = ?", url).Error    
    	if err != nil {
       		return "", fmt.Errorf("сокращенный url не найден")    
       	}
    	return urlRequest.FullUrl, nil
}

func (s *SqlDatabase) create(shortUrl, fullUrl string) error {
	if !IsUrl(fullUrl) {
		return fmt.Errorf("переданы некорректные url")
	}

	var existingUrl model.SqlModel
	result := s.dbSql.Where("short_url = ?", shortUrl).FirstOrCreate(&existingUrl, model.SqlModel{
		ShortUrl: shortUrl,
		FullUrl:  fullUrl,
	})
	if result.Error != nil {
		return fmt.Errorf("ошибка при создании/поиске записи")
	}
	
	logger.Println("val = ", existingUrl.FullUrl)
	return nil
}
