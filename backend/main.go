package main

import (
	"backend/model"
	"backend/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)



func main(){

	 // Create a new ServeMux for routing
     mux := http.NewServeMux()

	err := godotenv.Load()
	 if err != nil {
    log.Fatal("Error loading .env file")
  }

    config := model.Config{
    Host:     os.Getenv("DB_HOST"),
    Port:     os.Getenv("DB_PORT"),
    User:     os.Getenv("DB_USER"),
    Password: os.Getenv("DB_PASSWORD"),
    DBName:   os.Getenv("DB_NAME"),
    SSLMode:  os.Getenv("DB_SSLMODE"),
  }

   model.InitalizeDB(config)

   // routes
     routes.AuthRoutes(mux)
  // Start the server
  fmt.Println("Server listening on port 8080")
  http.ListenAndServe(":8080", mux)
}