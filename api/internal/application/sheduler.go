package application

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"time"
)

type SchedulerI interface {
	Start()
	SaveEventsToRecords()
}

type Scheduler struct {
	recordsApp RecordsI
	eventsApp  EventsI
	logger     *zap.SugaredLogger
}

func NewScheduler(recordsApp RecordsI, eventsApp EventsI, logger *zap.SugaredLogger) *Scheduler {
	return &Scheduler{
		recordsApp: recordsApp,
		eventsApp:  eventsApp,
		logger:     logger,
	}
}

func (s Scheduler) Start() {
	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		s.logger.Error("fail to create location: ", err)
		return
	}

	scheduler := cron.New(cron.WithLocation(location))
	_, err = scheduler.AddFunc("@midnight", s.SaveEventsToRecords)
	if err != nil {
		s.logger.Error("fail to create job: ", err)
		return
	}

	go scheduler.Start()
}

func (s Scheduler) SaveEventsToRecords() {
	s.logger.Info("start save records")
	searchDate := time.Now().AddDate(0, 0, -1)
	s.logger.Info("search date: ", searchDate.Format(time.RFC3339))
	events, err := s.eventsApp.GetForDay(searchDate.Year(), int(searchDate.Month()), searchDate.Day())
	if err != nil {
		s.logger.Error("fail to save events to records: ", err)
		return
	}

	if len(events) == 0 {
		return
	}

	err = s.recordsApp.SaveFromEvents(events)
	if err != nil {
		s.logger.Error("fail to save events to records: ", err)
		return
	}
}
