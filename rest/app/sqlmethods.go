package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	model "rest/app/model"
)

var dbSql *gorm.DB

func (s *SqlDatabase) connect() error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("ошибка получения текущей директории")
	}

	envPath := filepath.Join(wd, "..", ".env")

	err = godotenv.Load(envPath)
	if err != nil {
		return fmt.Errorf("ошибка получения environment файла")
	}

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s, sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	dbSql = db
	if err != nil {
		fmt.Errorf("ошибка при подключении к БД")
	}

	err = initSchema()
	if err != nil {
		return fmt.Errorf("ошибка при миграции")
	}
	return nil
}

func initSchema() error {
	return dbSql.AutoMigrate(&model.SqlModel{})
}

// get fixme: переделать получение поля из базы данных !
func (s *SqlDatabase) get(url string) (string, error) {
	var urlRequest model.SqlModel
	dbSql.First(&urlRequest, "url = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
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
