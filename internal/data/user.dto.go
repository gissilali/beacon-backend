package data

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) FromMapClaims(claims jwt.MapClaims) error {
	user := claims["user"].(map[string]interface{})
	if id, ok := user["id"].(float64); ok {
		u.ID = int64(id)
	}

	if name, ok := user["name"].(string); ok {
		u.Name = name
	}

	if email, ok := user["email"].(string); ok {
		u.Email = email
	}

	if createdAt, ok := user["created_at"].(string); ok {
		parsedTime, err := time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return err
		}
		u.CreatedAt = parsedTime
	}

	return nil
}

func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}

func (u *User) HashMatchesPassword(password string) (bool, error) {
	fmt.Println(u.Password, password, "Qui Vivra Verra")
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
