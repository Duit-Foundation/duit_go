package duit_action

type eventType string

const (
	update       eventType = "update"
	navigation   eventType = "navigation"
	openUrl      eventType = "openUrl"
	custom       eventType = "custom"
	sequenced    eventType = "sequenced"
	grouped      eventType = "grouped"
	timer        eventType = "timer"
	command      eventType = "command"
)

type Event struct {
	Type eventType `json:"type"`
}

type updateEvent struct {
	Event
	Updates map[string]any `json:"updates"`
}

type navigationEvent struct {
	Event
	Path  string         `json:"path"`
	Extra map[string]any `json:"extra,omitempty"`
}

type openUrlEvent struct {
	Event
	Url string `json:"url"`
}

type customEvent struct {
	Event
	Key   string         `json:"key"`
	Extra map[string]any `json:"extra,omitempty"`
}

type SequencedEvent struct {
	Delay int `json:"delay"`
	Event any `json:"event"`
}

type sequencedEventGroup struct {
	Event
	Events []*SequencedEvent `json:"events"`
}

type commonEventGroup struct {
	Event
	Events []interface{} `json:"events"`
}

type timerEvent struct {
	Event
	Delay  int `json:"delay"`
	TEvent any `json:"event"`
}

type commandEvent struct {
	Event
	ControllerId string `json:"controllerId"`
	CommandData  any    `json:"commandData"`
}

// NewUpdateEvent creates a new update event.
//
// It takes a map of updates as a parameter.
// It returns a pointer to an updateEvent.
func UpdateEvent(updates map[string]any) *updateEvent {
	return &updateEvent{
		Event: Event{
			Type: update,
		},
		Updates: updates,
	}
}

func NavigationEvent(path string, extra map[string]any) *navigationEvent {
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

func CustomEvent(key string, extra map[string]any) *customEvent {
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
		Event: Event{
			Type: sequenced,
		},
		Events: events,
	}
}

func CommonEventGroup(events []any) *commonEventGroup {
	return &commonEventGroup{
		Event: Event{
			Type: grouped,
		},
		Events: events,
	}
}

func TimerEvent(delay int, event any) *timerEvent {
	return &timerEvent{
		Event: Event{
			Type: timer,
		},
		Delay:  delay,
		TEvent: event,
	}
}

func CommandEvent(controllerId string, commandData any) *commandEvent {
	return &commandEvent{
		Event: Event{
			Type: command,
		},
		ControllerId: controllerId,
		CommandData:  commandData,
	}
}
