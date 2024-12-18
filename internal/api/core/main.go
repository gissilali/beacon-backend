package core

import (
	"beacon.silali.com/internal/api/config"
	"beacon.silali.com/internal/api/data"
	"beacon.silali.com/internal/api/dtos"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log"
)

type AppContext struct {
	Config  *config.Config
	Logger  *log.Logger
	Version string
	Models  data.Models
}

func New(cfg *config.Config, logger *log.Logger, models data.Models, version string) *AppContext {
	return &AppContext{
		Config:  cfg,
		Logger:  logger,
		Version: version,
		Models:  models,
	}
}

func (app *AppContext) CurrentUser(c echo.Context) (*dtos.User, error) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil, errors.New("JWT token missing or invalid")
	}

	claims, claimsIsOk := token.Claims.(jwt.MapClaims)
	if !claimsIsOk {
		return nil, errors.New("failed to cast claims as jwt.MapClaims")
	}

	user := &dtos.User{}
	if err := user.FromMapClaims(claims); err != nil {
		return nil, err
	}

	return user, nil
}
