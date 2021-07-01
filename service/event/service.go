package event

import (
		"eventlog/cmd/query"
		"eventlog/entity"
		"eventlog/infrastructure/repository"
)

type service struct {
		repository repository.Repository
}

func NewService(r repository.Repository) Service {
		return &service{repository: r}
}

func (s *service) CreateEvent(event *entity.Event) (*entity.Event, error) {
		return s.repository.SaveEvent(event)
}

func (s *service) FetchEventsByType(query *query.TypeQuery) ([]entity.Event, error) {
		return s.repository.GetEventsByType(query)
}

func (s *service) FetchEventsByTimeRange(query *query.TimeRangeQuery) ([]entity.Event, error) {
		return s.repository.GetEventsByTimeRange(query)
}

func (s *service) FetchEventsByTypeAndTimeRange(query *query.TypeTimeRangeQuery) ([]entity.Event, error) {
		return s.repository.GetEventsByTypeAndTimeRange(query)
}
