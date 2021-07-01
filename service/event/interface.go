package event

import (
		"eventlog/cmd/query"
		"eventlog/entity"
)

//Service is an interface from which our cmd module can access our repository of all models
type Service interface {
		CreateEvent(event *entity.Event) (*entity.Event, error)
		FetchEventsByType(query *query.TypeQuery) ([]entity.Event, error)
		FetchEventsByTimeRange(query *query.TimeRangeQuery) ([]entity.Event, error)
		FetchEventsByTypeAndTimeRange(query *query.TypeTimeRangeQuery) ([]entity.Event, error)
}
