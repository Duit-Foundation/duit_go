package duit_core

type eventType string

const (
	update       eventType = "update"
	updateLayout eventType = "updateLayout"
)

type event struct {
	Type eventType `json:"type"`
}

type updateEvent struct {
	event
	Updates map[string]interface{} `json:"updates"`
}

type layoutUpdateEvent struct {
	event
	Layout *DuitElementModel `json:"layout"`
}

// NewUpdateEvent creates a new update event.
//
// It takes a map of updates as a parameter.
// It returns a pointer to an updateEvent.
func NewUpdateEvent(updates map[string]interface{}) *updateEvent {
	return &updateEvent{
		event: event{
			Type: update,
		},
		Updates: updates,
	}
}

func NewLayoutUpdateEvent(payload *DuitElementModel) *layoutUpdateEvent {
	return &layoutUpdateEvent{
		event: event{
			Type: updateLayout,
		},
		Layout: payload,
	}
}
