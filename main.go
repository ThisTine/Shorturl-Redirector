package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ThisTine/shorturl-redirector/db"
	"github.com/ThisTine/shorturl-redirector/initialize"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if ok := os.Getenv("GO_ENV"); ok != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	app := fiber.New(fiber.Config{
		UnescapePath: true,
	})
	err := db.DBsetup()
	if err != nil {
		fmt.Println(err)
		return
	}
	initialize.Initialize(app)
	app.Listen(":8000")
}
