package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"net/http"
	"strconv"
	"time"
)

type getRecordsPayload struct {
	Since int64 `query:"since"`
	Limit int   `query:"limit"`
}

func (s Server) getRecords(c echo.Context) error {
	var payload getRecordsPayload

	err := s.bindAndValidate(c, &payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	records, err := s.recordsApp.Get(payload.Since, payload.Limit)
	if err != nil {
		s.logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if len(records) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, records)
}

func (s Server) createRecord(c echo.Context) error {
	var payloadRecord entity.Record
	err := s.bindAndValidate(c, &payloadRecord)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = s.recordsApp.Create(&payloadRecord)
	if err != nil {
		s.logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, payloadRecord)
}

func (s Server) deleteRecord(c echo.Context) error {
	recordID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = s.recordsApp.Delete(int64(recordID))
	if err != nil {
		s.logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s Server) getStat(c echo.Context) error {
	statType := c.QueryParam("type")

	switch statType {
	case "day":
		return s.getStatForDay(c)
	case "month":
		return s.getStatForMonth(c)
	case "year":
		return s.getStatForYear(c)
	default:
		return echo.NewHTTPError(http.StatusBadRequest)
	}
}

func (s Server) getStatForDay(c echo.Context) error {
	dateStr := c.QueryParam("date")
	datePayload, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	records, err := s.recordsApp.GetRecordsForDay(datePayload)
	if err != nil {
		s.logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, records)
}

func (s Server) getStatForMonth(c echo.Context) error {
	var payloadDate searchForDateRequest
	err := s.bindAndValidate(c, &payloadDate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	values, err := s.recordsApp.GetStatForMonth(payloadDate.Year, payloadDate.Month)
	if err != nil {
		s.logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, values)
}

func (s Server) getStatForYear(c echo.Context) error {
	var payloadDate searchForDateRequest
	err := s.bindAndValidate(c, &payloadDate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	values, err := s.recordsApp.GetStatForYear(payloadDate.Year)
	if err != nil {
		s.logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, values)
}

func (s Server) saveEventsToRecords(c echo.Context) error {
	dateStr := c.QueryParam("date")
	searchDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	events, err := s.eventsApp.GetForDay(searchDate.Year(), int(searchDate.Month()), searchDate.Day())

	if err != nil {
		s.logger.Error("fail to save events to records: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if len(events) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	err = s.recordsApp.SaveFromEvents(events)
	if err != nil {
		s.logger.Error("fail to save events to records: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

type updatePayload struct {
	Value float32 `json:"value"`
}

func (s Server) updateRecord(c echo.Context) error {
	recordID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	var payload updatePayload
	err = s.bindAndValidate(c, &payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = s.recordsApp.Update(int64(recordID), payload.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "fail to update error")
	}

	return c.NoContent(http.StatusOK)
}
