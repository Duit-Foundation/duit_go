package duit_core

type eventType string

const (
	update       eventType = "update"
	updateLayout eventType = "updateLayout"
	navigation   eventType = "navigation"
	openUrl      eventType = "openUrl"
)

type Event struct {
	Type eventType `json:"type"`
}

type updateEvent struct {
	Event
	Updates map[string]interface{} `json:"updates"`
}

type layoutUpdateEvent struct {
	Event
	Layout *DuitElementModel `json:"layout"`
}

type navigationEvent struct {
	Event
	Path  string                 `json:"path"`
	Extra map[string]interface{} `json:"extra,omitempty"`
}

type openUrlEvent struct {
	Event
	Url string `json:"url"`
}

// NewUpdateEvent creates a new update event.
//
// It takes a map of updates as a parameter.
// It returns a pointer to an updateEvent.
func NewUpdateEvent(updates map[string]interface{}) *updateEvent {
	return &updateEvent{
		Event: Event{
			Type: update,
		},
		Updates: updates,
	}
}

func NewLayoutUpdateEvent(payload *DuitElementModel) *layoutUpdateEvent {
	return &layoutUpdateEvent{
		Event: Event{
			Type: updateLayout,
		},
		Layout: payload,
	}
}

func NewNavigationEvent(path string, extra map[string]interface{}) *navigationEvent {
	return &navigationEvent{
		Event: Event{
			Type: navigation,
		},
		Path:  path,
		Extra: extra,
	}
}

func NewOpenUrlEvent(url string) *openUrlEvent {
	return &openUrlEvent{
		Event: Event{
			Type: openUrl,
		},
		Url: url,
	}
}
