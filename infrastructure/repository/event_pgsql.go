package repository

import (
		"database/sql"
		"eventlog/cmd/query"
		"eventlog/entity"
)

type repositoryPGSQL struct {
		DB *sql.DB
}

func NewRepository(DB *sql.DB) *repositoryPGSQL {
		return &repositoryPGSQL{DB: DB}
}

func (r *repositoryPGSQL) SaveEvent(event *entity.Event) (*entity.Event, error) {
		stmt, err := r.DB.Prepare(`INSERT INTO event_log (ts, type, content) VALUES($1::bigint, $2, $3)`)
		if err != nil {
				return nil, err
		}
		_, err = stmt.Exec(
				event.EventTimestamp,
				event.EventType,
				event.EventContent,
		)
		if err != nil {
				return nil, err
		}
		err = stmt.Close()
		if err != nil {
				return nil, err
		}
		return event, nil
}

func (r *repositoryPGSQL) GetEventsByType(query *query.TypeQuery) ([]entity.Event, error) {

		rows, err := r.DB.Query(`SELECT * FROM event_log el WHERE el.type LIKE $1`, query.Type)
		if err != nil {
				return nil, err
		}

		events, err := prepareEntitySlice(rows)
		_ = rows.Close()

		return events, nil
}

func (r *repositoryPGSQL) GetEventsByTimeRange(query *query.TimeRangeQuery) ([]entity.Event, error) {

		rows, err := r.DB.Query(`SELECT * FROM event_log el WHERE el.ts >= $1 AND el.ts <= $2`, query.From, query.To)
		if err != nil {
				return nil, err
		}

		events, err := prepareEntitySlice(rows)
		_ = rows.Close()

		return events, nil
}

func (r *repositoryPGSQL) GetEventsByTypeAndTimeRange(query *query.TypeTimeRangeQuery) ([]entity.Event, error) {

		rows, err := r.DB.Query(`SELECT * FROM event_log el WHERE (el.ts >= $1 AND el.ts <= $2) AND el.type LIKE $3`, query.From, query.To, query.Type)
		if err != nil {
				return nil, err
		}

		events, err := prepareEntitySlice(rows)
		_ = rows.Close()

		return events, nil

}

func prepareEntitySlice(rows *sql.Rows) ([]entity.Event, error) {

		var events []entity.Event

		for rows.Next() {
				var eventTs string
				var eventType string
				var eventContent string
				err := rows.Scan(&eventTs, &eventType, &eventContent)

				event := entity.Event{
						EventTimestamp: eventTs,
						EventType:      eventType,
						EventContent:   eventContent,
				}
				err = event.Validate()
				if err != nil {
						return nil, err
				}

				events = append(events, event)

		}

		return events, nil
}
