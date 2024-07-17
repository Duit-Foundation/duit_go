package duit_core

type eventType string

const (
	update       eventType = "update"
	updateLayout eventType = "updateLayout"
	navigation   eventType = "navigation"
	openUrl      eventType = "openUrl"
	custom       eventType = "custom"
	sequenced    eventType = "sequenced"
	grouped      eventType = "grouped"
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

type customEvent struct {
	Event
	Key   string                 `json:"key"`
	Extra map[string]interface{} `json:"extra,omitempty"`
}

type SequencedEvent struct {
	Delay int         `json:"delay"`
	Event interface{} `json:"event"`
}

type sequencedEventGroup struct {
	Event
	Events []*SequencedEvent `json:"events"`
}

type commonEventGroup struct {
	Event
	Events []interface{} `json:"events"`
}

// NewUpdateEvent creates a new update event.
//
// It takes a map of updates as a parameter.
// It returns a pointer to an updateEvent.
func UpdateEvent(updates map[string]interface{}) *updateEvent {
	return &updateEvent{
		Event: Event{
			Type: update,
		},
		Updates: updates,
	}
}

func LayoutUpdateEvent(payload *DuitElementModel) *layoutUpdateEvent {
	return &layoutUpdateEvent{
		Event: Event{
			Type: updateLayout,
		},
		Layout: payload,
	}
}

func NavigationEvent(path string, extra map[string]interface{}) *navigationEvent {
	return &navigationEvent{
		Event: Event{
			Type: navigation,
		},
		Path:  path,
		Extra: extra,
	}
}

func OpenUrlEvent(url string) *openUrlEvent {
	return &openUrlEvent{
		Event: Event{
			Type: openUrl,
		},
		Url: url,
	}
}

func CustomEvent(key string, extra map[string]interface{}) *customEvent {
	return &customEvent{
		Event: Event{
			Type: custom,
		},
		Key:   key,
		Extra: extra,
	}
}

func SequencedEventGroup(events []*SequencedEvent) *sequencedEventGroup {
	return &sequencedEventGroup{
		Events: events,
	}
}

func CommonEventGroup(events []interface{}) *commonEventGroup {
	return &commonEventGroup{
		Events: events,
	}
}
