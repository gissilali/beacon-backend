package handlers

import (
	"beacon.silali.com/internal/api/core"
	"beacon.silali.com/internal/api/dtos"
	"beacon.silali.com/internal/api/validator"
	"crypto/rand"
	"encoding/base64"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAccessKeys(c echo.Context, app *core.AppContext) error {
	user, err := app.CurrentUser(c)
	if err != nil {
		return respondWithError(c, http.StatusInternalServerError, "Failed to get current user", err)
	}

	accessKeys, err := app.Models.AccessKey.GetUserAccessKeys(user.ID)
	if err != nil {
		return respondWithError(c, http.StatusInternalServerError, "Failed to retrieve access keys", err)
	}

	return c.JSON(http.StatusOK, accessKeys)
}

func CreateAccessKey(c echo.Context, app *core.AppContext) error {
	// Bind the request
	request := new(dtos.CreateAccessKeyRequest)
	if err := c.Bind(request); err != nil {
		return respondWithError(c, http.StatusBadRequest, dtos.ErrorCodeUnsupportedRequest, err)
	}

	// Validate the request
	v := validator.New(app.Models)
	v.ValidateCreateAccessKeyRequest(request)
	if !v.Valid() {
		return c.JSON(http.StatusUnprocessableEntity, dtos.ErrorResponse{
			Code:    dtos.ErrorInvalidForm,
			Details: v.Errors,
		})
	}

	// Get current user
	user, err := app.CurrentUser(c)
	if err != nil {
		return respondWithError(c, http.StatusInternalServerError, "Failed to get current user", err)
	}

	// Generate access key
	accessKey, err := generateAccessKey()
	if err != nil {
		return respondWithError(c, http.StatusInternalServerError, "Failed to generate access key", err)
	}

	// Create access key
	newAccessKey, err := app.Models.AccessKey.CreateAccessKey(request.Name, accessKey, user.ID)
	if err != nil {
		return respondWithError(c, http.StatusInternalServerError, "Failed to create access key", err)
	}

	// Return the new access key
	return c.JSON(http.StatusCreated, newAccessKey)
}

func generateAccessKey() (string, error) {
	token := make([]byte, 64)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(token), nil
}
