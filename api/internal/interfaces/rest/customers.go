package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"net/http"
	"strconv"
)

func (s Server) createCustomer(c echo.Context) error {
	payloadCustomer := new(entity.Customer)

	if err := s.bindAndValidate(c, payloadCustomer); err != nil {
		return err
	}

	createdCustomer, err := s.customerApp.Create(payloadCustomer)
	if err != nil {
		s.logger.Error("fail to create customer: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusCreated, createdCustomer)
}

func (s Server) getCustomer(c echo.Context) error {
	customerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "customer id should be integer")
	}

	customer, err := s.customerApp.Get(int64(customerID))
	if err != nil {
		s.logger.Error("fail to get customer: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, customer)
}

func (s Server) searchCustomer(c echo.Context) error {
	searchText := c.QueryParam("searchText")
	since := c.QueryParam("since")
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 0
	}

	customers, err := s.customerApp.Search(searchText, since, limit)
	if err != nil {
		s.logger.Error("fail to search customer: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	if len(customers) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, customers)
}

func (s Server) deleteCustomer(c echo.Context) error {
	customerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "customer id should be integer")
	}

	err = s.customerApp.Delete(int64(customerID))
	if err != nil {
		s.logger.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return nil
}

func (s Server) updateCustomer(c echo.Context) error {
	payloadCustomer := new(entity.Customer)

	if err := s.bindAndValidate(c, payloadCustomer); err != nil {
		return err
	}
	s.logger.Debugf("PAYLOAD: %+v", payloadCustomer)

	err := s.customerApp.Update(payloadCustomer)
	if err != nil {
		s.logger.Error("fail to update customer: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	return nil
}
