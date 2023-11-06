package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var (
	logger   = log.Default()
	database DbInterface
)

func Server(dbType *string) {
	logger.Println(makeShortUrl("http://localhost:8080"))
	server := gin.Default()
	switch *dbType {

	case "sql":
		database = &SqlDatabase{}

	case "in_memory":
		database = &InMemoryDatabase{}

	default:
		logger.Println("некорректное значение", "текущая бд - ", database)
		os.Exit(418)
	}

	err := database.connect()
	if err != nil {
		logger.Println(err)
	}

	server.GET("/get_url", getUrl)
	server.POST("/post_url", postUrl)

	err = server.Run(":8080")
	if err != nil {
		logger.Println("ошибка сервера")
	}
}
