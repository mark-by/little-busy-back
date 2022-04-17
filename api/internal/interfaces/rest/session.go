package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"net/http"
)

func (s Server) logIn(c echo.Context) error {
	user := new(entity.User)

	if err := s.bindAndValidate(c, user); err != nil {
		return err
	}

	foundUser, session, err := s.userApp.LogIn(user.Username, user.Password)
	if err != nil {
		s.logger.Error("fail to login: ", err)
		return echo.NewHTTPError(http.StatusUnauthorized, "something wrong")
	}

	c.SetCookie(s.createAuthCookie(session))
	return c.JSON(http.StatusOK, foundUser)
}
