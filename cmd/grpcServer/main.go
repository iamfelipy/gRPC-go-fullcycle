package main

import (
	"database/sql"
	"net"

	"github.com/devfullcycle/14-gRPC/internal/database"
	"github.com/devfullcycle/14-gRPC/internal/pb"
	"github.com/devfullcycle/14-gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService((*categoryDb))

	grpcServer := grpc.NewServer()
	// isso foi gerado no /internal/pb/course_category_grpc.pb.go pelo protoc só preciso ligar ele ao grpcServer
	// atachei o serviço ao servidor grpc
	// preciso abrir um conexão tcp para que eu consiga falar com o grpc, abrir um porta tcp em go
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
