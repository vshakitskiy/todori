package routes

import "github.com/labstack/echo/v4"

func AppendRoutes(g *echo.Group) {
	v1 := g.Group("/v1")

	AppendUserRoutes(v1)
}
