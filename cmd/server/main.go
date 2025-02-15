//* cmd/server/main.go:
//****************************************************************************************
//* The entry point of your application.
//* This file will set up the server, load configurations,
//* and start the HTTP server.
//****************************************************************************************

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Weston-Cox/ClassPath-Backend/internal/config"
	"github.com/Weston-Cox/ClassPath-Backend/internal/database"
	"github.com/Weston-Cox/ClassPath-Backend/internal/handlers"

	"github.com/jackc/pgx/v5"
)


func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Panic("Error loading config in Main")
	}

	// Exits if server cannot connect to the database.
	connection, err := connectToDatabase(config.DatabaseURL)
	if err != nil {
		log.Fatal("Error establishing initial connection to Database")
	}
	defer connection.Close(context.Background())

	http_handler := handlers.NewHttpHandler(connection)
	http_handler.SetupRoutes()
	
	fmt.Printf("Starting server at %s\n", config.ServerAddress)
	log.Fatal(http.ListenAndServe(config.ServerAddress, nil))
}


func connectToDatabase(databaseURL string) (*pgx.Conn, error) {
	connection, err := database.Connect(databaseURL)
	if err != nil {
		return nil, err
	} else {
		fmt.Println("Database connected!")
		return connection, nil
	}

}


