package main

import (
	"flag"
	"rest/app"
)

func main() {
	dbType := flag.String("db", "sql", "database type")
	flag.Parse()
	app.Server(dbType)
}
