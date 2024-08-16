package lib

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tokens"
)

const AuthCookieName = "auth"

func setAuthToken(app core.App, c echo.Context, user *models.Record) error {
	s, tokenErr := tokens.NewRecordAuthToken(app, user)
	if tokenErr != nil {
		return fmt.Errorf("Login failed")
	}

	c.SetCookie(&http.Cookie{
		Name:     AuthCookieName,
		Value:    s,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		MaxAge:   int(app.Settings().RecordAuthToken.Duration),
	})

	return nil
}

func Login(app core.App, c echo.Context, username string, password string) error {
	user, err := app.Dao().FindAuthRecordByUsername("users", username)
	if err != nil {
		return fmt.Errorf("no user found for this email/password")
	}

	valid := user.ValidatePassword(password)
	if !valid {
		return fmt.Errorf("no user found for this email/password")
	}

	return setAuthToken(app, c, user)
}

type RegisterError struct {
	Username       string
	Password       string
	PasswordRepeat string
	Unknown        string
}

func Register(app core.App, c echo.Context, username string, password string, passwordRepeat string) RegisterError {
	user, _ := app.Dao().FindAuthRecordByUsername("users", username)

	errors := RegisterError{}

	if user != nil {
		errors.Username = "Username already taken"
	}

	if password != passwordRepeat {
		errors.PasswordRepeat = "Passwords  match"
	}

	collection, err := app.Dao().FindCollectionByNameOrId("users")
	if err != nil {
		errors.Unknown = err.Error()
		return errors
	}

	newUser := models.NewRecord(collection)
	newUser.SetPassword(password)
	newUser.SetUsername(username)

	if err = app.Dao().SaveRecord(newUser); err != nil {
		errors.Unknown = err.Error()
		return errors
	}

	err = setAuthToken(app, c, newUser)

	if err != nil {
		errors.Unknown = err.Error()
	}

	return errors
}

func Logout(app core.App, c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     AuthCookieName,
		Value:    "",
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		MaxAge:   -1,
	})

	return nil
}
