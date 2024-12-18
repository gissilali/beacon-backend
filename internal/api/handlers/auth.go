package handlers

import (
	"beacon.silali.com/internal/api/core"
	"beacon.silali.com/internal/api/data"
	"beacon.silali.com/internal/api/dtos"
	"beacon.silali.com/internal/api/validator"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func RegisterUser(c echo.Context, app *core.AppContext) error {
	request := new(dtos.RegisterUserRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Code:    dtos.ErrorInvalidCredentials,
			Message: fmt.Sprintf("%s", err),
			Details: err,
		})
	}

	v := validator.New(app.Models)
	v.ValidateRegisterUserRequest(request)
	if !v.Valid() {
		return c.JSON(http.StatusUnprocessableEntity, dtos.ErrorResponse{
			Code:    dtos.ErrorInvalidForm,
			Details: v.Errors,
		})
	}

	newUser, err := app.Models.User.Create(&dtos.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Code:    dtos.ErrorCodeUnsupportedRequest,
			Message: fmt.Sprintf("%s", err),
			Details: err,
		})
	}

	return c.JSON(http.StatusCreated, newUser)
}

func LoginUser(c echo.Context, app *core.AppContext) error {
	request := new(dtos.LoginUserRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Code:    dtos.ErrorCodeUnsupportedRequest,
			Message: fmt.Sprintf("%s", err),
			Details: err,
		})
	}

	v := validator.New(app.Models)

	v.ValidateLoginUserRequest(request)

	if !v.Valid() {
		return c.JSON(http.StatusUnprocessableEntity, dtos.ErrorResponse{
			Code:    dtos.ErrorInvalidForm,
			Details: v.Errors,
		})
	}

	user, err := attemptAuth(request.Email, request.Password, app.Models.User)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, dtos.ErrorResponse{
			Code:    dtos.ErrorInvalidCredentials,
			Message: fmt.Sprintf("%s", err),
		})
	}

	tokens, issueTokenErr := issueToken(user)

	if issueTokenErr != nil {
		return c.JSON(http.StatusInternalServerError, &dtos.ErrorResponse{
			Code:    dtos.ErrorFailedToIssueTokens,
			Message: fmt.Sprintf("Token issuance failed: %v", err),
		})
	}

	saveTokenErr := app.Models.AuthToken.Create(tokens.RefreshToken, user.ID)
	if saveTokenErr != nil {
		return c.JSON(http.StatusInternalServerError, &dtos.ErrorResponse{
			Code:    dtos.ErrorFailedToSaveTokens,
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

func attemptAuth(email string, password string, userModel data.UserModel) (*dtos.User, error) {
	user, err := userModel.GetByEmail(email)

	if err != nil {
		return nil, err
	}

	isOk, hashErr := user.HashMatchesPassword(password)

	if hashErr != nil || !isOk {
		return nil, fmt.Errorf("invalid credentials")
	}

	return user, nil
}

func issueToken(user *dtos.User) (*dtos.AuthTokens, error) {
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

	return &dtos.AuthTokens{
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
