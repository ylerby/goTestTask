package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc/cmd/methods"
	service "grpc/pkg/grpc_service/api/grpc_service"
)

const grpcPort = 50051

func main() {
	dbType := flag.String("db", "sql", "database type")
	flag.Parse()

	switch *dbType {

	case "sql":
		methods.Database = &methods.SqlDatabase{}
	case "in_memory":
		methods.Database = &methods.InMemoryDatabase{}
	default:
		os.Exit(418)
	}

	err := methods.Database.Connect()
	if err != nil {
		log.Println(err)
	}
	log.Println("успешное подключение к БД")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	service.RegisterShortUrlServer(s, &methods.Server{})
	service.RegisterFullUrlServer(s, &methods.Server{})

	log.Printf("server listening at %s", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %s", err)
	}
}
