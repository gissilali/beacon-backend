package main

import (
	"beacon.silali.com/internal/data"
	"beacon.silali.com/internal/validator"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (app *application) registerUserHandler(c echo.Context) error {
	request := new(data.RegisterUserRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorEnvelope{
			Code:    ErrorCodeUnsupportedRequest,
			Message: fmt.Sprintf("%s", err),
			Details: err,
		})
	}

	v := validator.New(app.models)
	v.ValidateRegisterUserRequest(request)
	if !v.Valid() {
		return c.JSON(http.StatusUnprocessableEntity, ErrorEnvelope{
			Code:    ErrorInvalidForm,
			Details: v.Errors,
		})
	}

	newUser, err := app.models.User.Create(&data.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorEnvelope{
			Code:    ErrorCodeUnsupportedRequest,
			Message: fmt.Sprintf("%s", err),
			Details: err,
		})
	}

	return c.JSON(http.StatusCreated, newUser)
}

func (app *application) createAuthenticationTokenHandler(c echo.Context) error {
	request := new(data.LoginUserRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorEnvelope{
			Code:    ErrorCodeUnsupportedRequest,
			Message: fmt.Sprintf("%s", err),
			Details: err,
		})
	}

	v := validator.New(app.models)

	v.ValidateLoginUserRequest(request)

	if !v.Valid() {
		return c.JSON(http.StatusUnprocessableEntity, ErrorEnvelope{
			Code:    ErrorInvalidForm,
			Details: v.Errors,
		})
	}

	user, err := app.attemptAuth(request.Email, request.Password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorEnvelope{
			Code:    ErrorInvalidCredentials,
			Message: fmt.Sprintf("%s", err),
		})
	}

	tokens, issueTokenErr := issueToken(user)

	if issueTokenErr != nil {
		return c.JSON(http.StatusInternalServerError, &ErrorEnvelope{
			Code:    "FAILED_TO_ISSUE_TOKEN",
			Message: fmt.Sprintf("Token issuance failed: %v", err),
		})
	}

	saveTokenErr := app.models.AuthToken.Create(tokens.RefreshToken, user.ID)
	if saveTokenErr != nil {
		return c.JSON(http.StatusInternalServerError, ErrorEnvelope{
			Code:    "FAILED_TO_SAVE_TOKEN",
			Message: fmt.Sprintf("Failed to save token: %v", err),
		})
	}

	cookie := new(http.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = tokens.AccessToken
	cookie.Expires = time.Now().Add(1 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = false

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, envelope{
		"user":      user,
		"tokens":    tokens,
		"expiresIn": 3600,
	})
}

func (app *application) attemptAuth(email string, password string) (*data.User, error) {
	user, err := app.models.User.GetByEmail(email)

	if err != nil {
		return nil, err
	}

	isOk, hashErr := user.HashMatchesPassword(password)

	if hashErr != nil || !isOk {
		return nil, fmt.Errorf("invalid credentials")
	}

	return user, nil
}

func issueToken(user *data.User) (*data.AuthTokens, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	refreshToken, refreshTokenErr := generateRefreshToken()

	if refreshTokenErr != nil {
		return nil, refreshTokenErr
	}

	return &data.AuthTokens{
		AccessToken:  tokenString,
		RefreshToken: refreshToken,
	}, nil
}

func generateRefreshToken() (string, error) {
	token := make([]byte, 64) // 64 bytes = 512 bits of randomness
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(token), nil
}
