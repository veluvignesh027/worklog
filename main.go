package main

import (
	"log"
	"os"

	"example.com/vb/database"
	"example.com/vb/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not exist")
	}

	err := database.ConnectAllDatabase()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := gin.Default()
	handlers.InitHandlers(router)

	listenAddr := os.Getenv("HTTP_LISTER_ADDR")
	log.Fatal(router.Run(listenAddr))
	defer database.CloseAllDatabase()
}
