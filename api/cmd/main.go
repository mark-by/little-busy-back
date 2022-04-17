package main

import (
	"github.com/mark-by/little-busy-back/api/internal/application"
	"github.com/mark-by/little-busy-back/api/internal/infrastructure/microservices"
	"github.com/mark-by/little-busy-back/api/internal/interfaces/rest"
	"go.uber.org/zap"
	"log"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("fail to create logger: ", err)
	}

	authClient := microservices.NewAuthorizationClient("127.0.0.1:8000")
	crmClient := microservices.NewCRMClient("127.0.0.1:8001")
	schedulerClient := microservices.NewSchedulerClient("127.0.0.1:8002")

	authorizationRepository := microservices.NewAuthorization(authClient)
	userRepository := microservices.NewUser(crmClient)
	customerRepository := microservices.NewCustomers(crmClient)
	eventsRepository := microservices.NewEvents(schedulerClient)

	userApp := application.NewUser(userRepository, authorizationRepository)
	authApp := application.NewAuthorization(authorizationRepository)
	customerApp := application.NewCustomers(customerRepository, eventsRepository)
	eventsApp := application.NewEvents(eventsRepository, customerRepository)

	log.Print(rest.NewServer(&rest.ServerOptions{
		UserApp:     userApp,
		AuthApp:     authApp,
		CustomerApp: customerApp,
		EventsApp:   eventsApp,
		Logger:      logger.Sugar(),
		Config: &rest.ServerConfig{
			Address: ":9595",
		},
	}).Start())
}
