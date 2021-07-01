package entity_test

import (
		"eventlog/entity"
		"github.com/stretchr/testify/assert"
		"testing"
)

func TestNewEvent(t *testing.T) {
		e, err := entity.NewEvent("1625053791", "user_registration", "user_id: 1000, user_name: thor")
		assert.Nil(t, err)
		assert.Equal(t, e.EventTimestamp, "1625053791")
		assert.Equal(t, e.EventType, "user_registration")
		assert.NotNil(t, e.EventType)
}

func TestEvent_Validate(t *testing.T) {
		type test struct {
				EventTimestamp string
				EventType      string
				EventContent   string
				Want           error
		}

		multipleTest := []test{
				{EventTimestamp: "1625053791", EventType: "user_registration", EventContent: "user_id: 1000, user_name: thor", Want: nil},
		}

		for _, mt := range multipleTest {

				_, err := entity.NewEvent(mt.EventTimestamp, mt.EventType, mt.EventContent)
				assert.Equal(t, err, mt.Want)
		}
}
