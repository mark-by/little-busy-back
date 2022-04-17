package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"net/http"
	"time"
)

func (s Server) getUserBySession(c echo.Context) error {
	cookie, err := c.Cookie("session_id")
	if err == http.ErrNoCookie {
		return c.String(http.StatusUnauthorized, "")
	}
	if err != nil {
		return err
	}

	user, err := s.userApp.GetUserBySession(cookie.Value)
	if err != nil {
		c.SetCookie(s.deleteAuthCookie())
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (s Server) createAuthCookie(session *entity.AuthSession) *http.Cookie {
	authCookie := new(http.Cookie)
	authCookie.Name = "session_id"
	authCookie.Value = session.ID
	authCookie.Expires = time.Unix(session.ExpirationDate, 0)
	authCookie.Path = "/"
	authCookie.HttpOnly = true
	authCookie.Secure = true
	authCookie.SameSite = http.SameSiteNoneMode

	return authCookie
}

func (s Server) deleteAuthCookie() *http.Cookie {
	authCookie := new(http.Cookie)
	authCookie.Name = "session_id"
	authCookie.Value = ""
	authCookie.Expires = time.Now().AddDate(-5, 0, 0)
	authCookie.Path = "/"
	authCookie.HttpOnly = true
	authCookie.Secure = true
	authCookie.SameSite = http.SameSiteNoneMode

	return authCookie
}

func (s Server) signUp(c echo.Context) error {
	secret := c.Request().Header.Get("Authorization")

	if secret != s.config.Secret {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	user := new(entity.User)

	if err := s.bindAndValidate(c, user); err != nil {
		return err
	}

	createdUser, session, err := s.userApp.SignUp(user.Username, user.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}

	c.SetCookie(s.createAuthCookie(session))
	return c.JSON(http.StatusOK, createdUser)
}
