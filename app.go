package main

import (
	"fmt"
	"gRPC/app_services/implementation"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"net"
	"os"

	"gRPC/app_services/database"
	appService "gRPC/app_services/internal/services"
)

func init() {
	if err := godotenv.Load("app_services/app.env"); err != nil {
		panic(err)
	}
}

func main() {
	// open connection DB
	database.OpenConnection()

	// Inject dependencies
	appServer := implementation.NewAppImpl(database.DBContainer.DB)

	// init authenticate_services service
	srv := grpc.NewServer()
	appService.RegisterAppServiceServer(srv, appServer)

	listener, err := net.Listen("tcp", ":" + os.Getenv("API_PORT"))
	if err != nil {
		panic(err)
	}

	fmt.Println("[FSA] Started App Service")
	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
