package rest

import (
	"github.com/go-playground/form"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mark-by/little-busy-back/api/internal/application"
	"github.com/mark-by/little-busy-back/api/internal/config"
	"go.uber.org/zap"
	"net/http"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

type Server struct {
	logger      *zap.SugaredLogger
	eventsApp   application.EventsI
	userApp     application.UserI
	authApp     application.AuthorizationI
	customerApp application.CustomersI
	recordsApp  application.RecordsI
	settingsApp application.SettingsI
	config      *config.Config
	formDecoder *form.Decoder
}

type ServerOptions struct {
	UserApp     application.UserI
	AuthApp     application.AuthorizationI
	EventsApp   application.EventsI
	CustomerApp application.CustomersI
	RecordsApp  application.RecordsI
	SettingsApp application.SettingsI
	Logger      *zap.SugaredLogger
	Config      *config.Config
}

func NewServer(options *ServerOptions) *Server {
	return &Server{
		logger:      options.Logger,
		userApp:     options.UserApp,
		authApp:     options.AuthApp,
		eventsApp:   options.EventsApp,
		customerApp: options.CustomerApp,
		recordsApp:  options.RecordsApp,
		settingsApp: options.SettingsApp,
		config:      options.Config,
		formDecoder: form.NewDecoder(),
	}
}

func (s Server) Start() error {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost", "http://localhost:3000", "https://irina-massage.ru", "https://localhost.ru"},
		AllowMethods:     middleware.DefaultCORSConfig.AllowMethods,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/api/user", s.getUserBySession)
	e.POST("/api/user", s.signUp)

	e.POST("/api/session", s.logIn)

	e.GET("/api/customers", s.searchCustomer, s.authMiddleware)
	e.GET("/api/customers/:id", s.getCustomer, s.authMiddleware)
	e.POST("/api/customers", s.createCustomer, s.authMiddleware)
	e.PUT("/api/customers", s.updateCustomer, s.authMiddleware)
	e.DELETE("/api/customers/:id", s.deleteCustomer, s.authMiddleware)

	e.GET("/api/events", s.getEvents, s.authMiddleware)
	e.GET("/api/events/:id", s.getEvent, s.authMiddleware)
	e.POST("/api/events", s.createEvent, s.authMiddleware)
	e.PUT("/api/events", s.updateEvent, s.authMiddleware)
	e.DELETE("/api/events/:id", s.deleteEvent, s.authMiddleware)
	e.GET("/api/events/notPaid", s.getNotPaidEvents, s.authMiddleware)

	e.GET("/api/records", s.getRecords, s.authMiddleware)
	e.POST("/api/records", s.createRecord, s.authMiddleware)
	e.DELETE("/api/records/:id", s.deleteRecord, s.authMiddleware)
	e.GET("/api/records/stat", s.getStat, s.authMiddleware)
	e.POST("/api/records/saveEvents", s.saveEventsToRecords, s.authMiddleware)

	e.GET("/api/settings", s.getSettings, s.authMiddleware)
	e.PUT("/api/settings", s.updateSettings, s.authMiddleware)

	return e.Start(s.config.Address)
}
