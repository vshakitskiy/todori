package main

import (
	"log"
	"todori/ent"
	"todori/lib"
	"todori/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// godotenv
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	ent.OpenClient()

	// echo
	e := echo.New()
	routes.AppendRoutes(e.Group("/api"))
	e.Logger.Fatal(e.Start(":" + lib.Getenv("PORT")))
}
