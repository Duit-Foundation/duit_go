package duit_core

type eventType string

const (
	update       eventType = "update"
	updateLayout eventType = "updateLayout"
	navigation   eventType = "navigation"
	openUrl      eventType = "openUrl"
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

type navigationEvent struct {
	event
	Path  string                 `json:"path"`
	Extra map[string]interface{} `json:"extra,omitempty"`
}

type openUrlEvent struct {
	event
	Url string `json:"url"`
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

func NewNavigationEvent(path string, extra map[string]interface{}) *navigationEvent {
	return &navigationEvent{
		event: event{
			Type: navigation,
		},
		Path:  path,
		Extra: extra,
	}
}

func NewOpenUrlEvent(url string) *openUrlEvent {
	return &openUrlEvent{
		event: event{
			Type: openUrl,
		},
		Url: url,
	}
}
