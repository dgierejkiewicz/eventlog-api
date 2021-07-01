package entity

import "gopkg.in/go-playground/validator.v9"

type Event struct {
		EventTimestamp string `json:"event_timestamp" validate:"required"`
		EventType      string `json:"event_type" validate:"required"`
		EventContent   string `json:"event_content" validate:"required"`
}

func NewEvent(ts string, etype string, content string) (*Event, error) {

		e := &Event{
				EventTimestamp: ts,
				EventType:      etype,
				EventContent:   content,
		}
		err := e.Validate()

		if err != nil {
				return nil, err
		}

		return e, nil

}

func (e *Event) Validate() error {

		v := validator.New()
		if err := v.Struct(e);err != nil {
				errs := err.(validator.ValidationErrors)
				return errs
		}
		return nil
}