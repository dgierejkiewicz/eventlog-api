package entity

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
		if e.EventTimestamp == "" || e.EventType == "" || e.EventContent == "" {
				return ErrInvalidEntity
		}
		return nil
}
