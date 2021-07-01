package repository

import (
		"eventlog/cmd/query"
		"eventlog/entity"
)

//Repository interface allows us to access the CRUD Operations here.
type Repository interface {
		SaveEvent(event *entity.Event) (*entity.Event, error)
		GetEventsByType(query *query.TypeQuery) ([]entity.Event, error)
		GetEventsByTimeRange(query *query.TimeRangeQuery) ([]entity.Event, error)
		GetEventsByTypeAndTimeRange(query *query.TypeTimeRangeQuery) ([]entity.Event, error)
}