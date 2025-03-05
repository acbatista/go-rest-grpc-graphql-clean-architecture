package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	mysqldriver "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-rest-grpc-graphql-clean-architecture/graph"
	"go-rest-grpc-graphql-clean-architecture/graph/generated"
	"go-rest-grpc-graphql-clean-architecture/internal/infrastructure/repository"
	grpcserver "go-rest-grpc-graphql-clean-architecture/internal/interfaces/grpc"
	"go-rest-grpc-graphql-clean-architecture/internal/interfaces/rest"
	"go-rest-grpc-graphql-clean-architecture/internal/usecase"
	pb "go-rest-grpc-graphql-clean-architecture/proto"
)

func main() {
	// Database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysqldriver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Run migrations
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database instance: %v", err)
	}

	driver, err := mysql.WithInstance(sqlDB, &mysql.Config{})
	if err != nil {
		log.Fatalf("failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("failed to create migration instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migrations: %v", err)
	}

	// Setup dependencies
	orderRepo := repository.NewOrderRepository(db)
	listOrdersUseCase := usecase.NewListOrdersUseCase(orderRepo)
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepo)

	// Setup REST server
	orderHandler := rest.NewOrderHandler(listOrdersUseCase, createOrderUseCase)
	router := chi.NewRouter()
	router.Get("/order", orderHandler.ListOrders)
	router.Post("/order", orderHandler.CreateOrder)

	// Setup gRPC server
	grpcServer := grpc.NewServer()
	orderServer := grpcserver.NewOrderServer(listOrdersUseCase)
	pb.RegisterOrderServiceServer(grpcServer, orderServer)
	reflection.Register(grpcServer)

	// Setup GraphQL server
	resolver := graph.NewResolver(listOrdersUseCase, createOrderUseCase)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	router.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	// Start gRPC server
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("gRPC server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Start HTTP server
	log.Printf("REST server listening at http://localhost:8080")
	log.Printf("GraphQL playground available at http://localhost:8080/graphql")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
