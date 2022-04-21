package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"net/http"
	"strconv"
	"time"
)

func (s Server) createEvent(c echo.Context) error {
	var payloadEvent entity.Event
	err := s.bindAndValidate(c, &payloadEvent)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	event, err := s.eventsApp.Create(&payloadEvent)
	if err != nil {
		s.logger.Error("fail to create: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	return c.JSON(201, event)
}

func (s Server) deleteEvent(c echo.Context) error {
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "event id should be integer")
	}

	dateStr := c.QueryParam("date")
	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "date query param should be date")
	}
	withNextStr := c.QueryParam("with_next")
	var withNext bool
	if len(withNextStr) > 0 {
		withNext = true
	}

	if withNextStr == "false" {
		withNext = false
	}

	err = s.eventsApp.Delete(int64(eventID), date, withNext)
	if err != nil {
		s.logger.Error("fail to delete: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}
	return nil
}

func (s Server) updateEvent(c echo.Context) error {
	var payloadEvent entity.Event
	err := s.bindAndValidate(c, &payloadEvent)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	dateStr := c.QueryParam("date")
	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "date query param should be date")
	}
	withNextStr := c.QueryParam("with_next")
	var withNext bool
	if len(withNextStr) > 0 {
		withNext = true
	}

	if withNextStr == "false" {
		withNext = false
	}

	err = s.eventsApp.Update(&payloadEvent, date, withNext)
	if err != nil {
		s.logger.Error("fail to update: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	return nil
}

func (s Server) getEvent(c echo.Context) error {
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "event id should be integer")
	}

	event, err := s.eventsApp.Get(int64(eventID))
	if err != nil {
		s.logger.Error("fail to get event: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, event)
}

type searchForCustomerRequest struct {
	CustomerID int64     `query:"customer_id" validate:"required"`
	Since      time.Time `query:"since" validate:"required"`
	Days       int       `query:"days" validate:"lte=50"`
}

type searchForDateRequest struct {
	Year  int `query:"year" validate:"required"`
	Month int `query:"month" validate:"required,lte=12"`
	Day   int `query:"day" validate:"lte=31"`
}

func (s Server) getEvents(c echo.Context) error {
	searchType := c.QueryParam("type")

	var events []entity.Event

	switch searchType {
	case "customer":
		var request searchForCustomerRequest
		err := s.bindAndValidate(c, &request)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		events, err = s.eventsApp.GetForCustomer(request.CustomerID, request.Since, request.Days)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	case "month":
		fallthrough
	case "not_paid":
		fallthrough
	case "day":
		var request searchForDateRequest
		err := s.bindAndValidate(c, &request)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		if searchType == "not_paid" {
			events, err = s.eventsApp.GetNotPaidForDay(request.Year, request.Month, request.Day)
		} else if request.Day == 0 {
			events, err = s.eventsApp.GetForMonth(request.Year, request.Month)
		} else {
			events, err = s.eventsApp.GetForDay(request.Year, request.Month, request.Day)
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "undefined type")
	}

	return c.JSON(http.StatusOK, events)
}
