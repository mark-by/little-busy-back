package postgresql

import "github.com/mark-by/little-busy-back/scheduler/internal/domain/entity"

type eventPicker struct {
	weekDays map[int][]entity.Event
	days     map[int][]entity.Event
	daily    []entity.Event
}

func (c *eventPicker) AddToWeekDay(weekDay int, event entity.Event) {
	c.addToMapSlice(c.weekDays, weekDay, event)
}

func (c *eventPicker) AddToMonthDay(day int, event entity.Event) {
	c.addToMapSlice(c.days, day, event)
}

func (c *eventPicker) AddToDaily(event entity.Event) {
	c.daily = append(c.daily, event)
}

func (c eventPicker) GetForWeek(weekDay int) []entity.Event {
	return c.weekDays[weekDay]
}

func (c eventPicker) GetForMontDay(day int) []entity.Event {
	return c.days[day]
}

func (c eventPicker) GetDaily() []entity.Event {
	return c.daily
}

func (c eventPicker) addToMapSlice(m map[int][]entity.Event, day int, event entity.Event) {
	_, ok := m[day]
	if !ok {
		m[day] = []entity.Event{event}
		return
	}

	m[day] = append(m[day], event)
}

func newEventPicker() *eventPicker {
	return &eventPicker{
		weekDays: map[int][]entity.Event{},
		days:     map[int][]entity.Event{},
		daily:    []entity.Event{},
	}
}
