package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var (
	InMemoryDb  = map[string]string{}
	SqlDatabase *sql.DB
)

var logger = log.Default()

func Server(dbType *string) {
	switch *dbType {
	case "sql":
		logger.Println("sql")
	case "in_memory":
		logger.Println("in memory database")
	default:
		logger.Println("некорректное значение")
		os.Exit(-1)
	}

	server := gin.Default()
	server.GET("/get_url", getUrl)
	server.POST("/post_url", postUrl)
	server.Run(":8080")
}
