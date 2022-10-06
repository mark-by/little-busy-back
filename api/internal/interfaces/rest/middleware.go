package rest

import "github.com/labstack/echo/v4"

func (s Server) authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := s.getUserBySessionForMiddleware(c)
		if err != nil {
			return err
		}

		return next(c)
	}
}
