package main

import (
	"flag"
	"rest/app"
	"fmt"
)

func main() {
	dbType := flag.String("db", "sql", "database type")
	dbHost := flag.String("DB_HOST", " ", "database host")
	dbName := flag.String("DB_NAME", " ", "database name")
	fmt.Println("flags = ", dbHost, dbName)
	flag.Parse()
	app.Server(dbType)
}
