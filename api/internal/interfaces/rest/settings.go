package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"net/http"
)

func (s Server) getSettings(c echo.Context) error {
	settings, err := s.settingsApp.Get()
	if err != nil {
		s.logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, settings)
}

func (s Server) updateSettings(c echo.Context) error {
	var payloadSettings entity.Settings

	err := s.bindAndValidate(c, &payloadSettings)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = s.settingsApp.Update(&payloadSettings)
	if err != nil {
		s.logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
