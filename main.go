package main

import (
	"log"

	"github.com/abhishek-singh/config"
	_ "github.com/jinzhu/gorm/dialects/postgres"


	"github.com/abhishek-singh/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	app := gin.Default()
	godotenv.Load()
	config.ConnectDB()
	router.Routes(app)

	log.Fatal(app.Run(":8080"))
}
