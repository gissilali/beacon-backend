package main

import (
	"fmt"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (app *application) routes() {
	e := echo.New()

	e.GET("/v1/healthcheck", app.healthcheckHandler)
	e.GET("/v1/servers", app.healthcheckHandler)
	// auth
	e.POST("/v1/auth/register", app.registerUserHandler)
	e.POST("/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	protectedRoutes := e.Group("")
	protectedRoutes.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte("secret"),
		TokenLookup: "cookie:auth_token",
	}))

	protectedRoutes.GET("/v1/workspaces", app.getWorkspacesHandler)
	protectedRoutes.POST("/v1/workspaces", app.createWorkspaceHandler)
	protectedRoutes.GET("/v1/servers", app.getServersHandler)

	app.logger.Fatal(e.Start(fmt.Sprintf(":%d", app.config.port)))
}
