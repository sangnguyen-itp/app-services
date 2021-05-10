package main

import (
	"app-services/database"
	"app-services/implementation"
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	appService "app-services/internal/services"
)

func init() {
	if err := godotenv.Load(); err != nil {
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

	listener, err := net.Listen("tcp", ":"+os.Getenv("API_PORT"))
	if err != nil {
		panic(err)
	}

	fmt.Println("[FSA] Started App Service")
	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
