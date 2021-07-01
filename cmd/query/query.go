package query

type TypeQuery struct {
		Type string
}

func NewTypeQuery(eventType string) (*TypeQuery, error) {
		q := &TypeQuery{Type: eventType}
		err := q.Validate()
		if err != nil {
				return nil, err
		}
		return q, nil
}

func (t *TypeQuery) Validate() error {
		if t.Type == "" {
				return ErrInvalidQuery
		}
		return nil
}

type TimeRangeQuery struct {
		From string
		To   string
}

func NewTimeRangeQuery(eventFrom string, eventTo string) (*TimeRangeQuery, error) {
		q := &TimeRangeQuery{
				From: eventFrom,
				To:   eventTo,
		}
		err := q.Validate()
		if err != nil {
				return nil, err
		}
		return q, nil
}

func (t *TimeRangeQuery) Validate() error {
		if t.From == "" || t.To == "" {
				return ErrInvalidQuery
		}
		return nil
}

type TypeTimeRangeQuery struct {
		Type string
		From string
		To   string
}

func NewTypeTimeRangeQuery(eventType string,eventFrom string, eventTo string) (*TypeTimeRangeQuery, error) {
		q := &TypeTimeRangeQuery{
				Type: eventType,
				From: eventFrom,
				To:   eventTo,
		}
		err := q.Validate()
		if err != nil {
				return nil, err
		}
		return q, nil
}

func (t *TypeTimeRangeQuery) Validate() error {
		if t.Type == "" || t.From == "" || t.To == "" {
				return ErrInvalidQuery
		}
		return nil
}