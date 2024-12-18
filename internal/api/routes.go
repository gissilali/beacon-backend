package api

import (
	"beacon.silali.com/internal/api/core"
	"beacon.silali.com/internal/api/handlers"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(app *core.AppContext) *echo.Echo {
	e := echo.New()

	protectedRoutes := e.Group("")
	protectedRoutes.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte("secret"),
		TokenLookup: "cookie:auth_token",
	}))

	e.GET("/v1/healthcheck", func(c echo.Context) error {
		return handlers.HealthCheck(c, app)
	})
	e.POST("/v1/auth/register", func(c echo.Context) error {
		return handlers.RegisterUser(c, app)
	})
	e.POST("/v1/auth/login", func(c echo.Context) error {
		return handlers.LoginUser(c, app)
	})

	protectedRoutes.GET("/v1/workspaces", func(c echo.Context) error {
		return handlers.GetWorkspaces(c, app)
	})
	protectedRoutes.POST("/v1/workspaces", func(c echo.Context) error {
		return handlers.CreateWorkspace(c, app)
	})
	protectedRoutes.POST("/v1/access-keys", func(c echo.Context) error {
		return handlers.CreateAccessKey(c, app)
	})
	protectedRoutes.GET("/v1/access-keys", func(c echo.Context) error {
		return handlers.GetAccessKeys(c, app)
	})
	//protectedRoutes.GET("/v1/servers", app.getServersHandler)

	return e
}
