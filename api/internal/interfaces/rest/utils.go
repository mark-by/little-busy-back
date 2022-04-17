package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) bindAndValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(i); err != nil {
		return err
	}

	return nil
}
