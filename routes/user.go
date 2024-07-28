package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AppendUserRoutes(g *echo.Group) {
	g.Group("/users")

	g.GET("/ping", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"status":  "ok",
			"message": "pong",
		})
	})
}
