package main

import (
	"log"
	"todori/lib"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// godotenv
	err := godotenv.Load(".env")
	if err != nil {
			log.Fatal("Error loading .env file")
	}

	// echo
	e := echo.New()

	e.GET("/ping", func(ctx echo.Context) error {
		return ctx.JSON(200, map[string]interface{}{
			"status": "ok",
			"message": "pong",
		})
	})

	e.Logger.Fatal(e.Start(":" + lib.Getenv("PORT")))
}