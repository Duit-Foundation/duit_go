package duit_core

import "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"

type eventType string

const (
	update           eventType = "update"
	updateLayout     eventType = "updateLayout"
	navigation       eventType = "navigation"
	openUrl          eventType = "openUrl"
	custom           eventType = "custom"
	sequenced        eventType = "sequenced"
	grouped          eventType = "grouped"
	animationTrigger eventType = "animationTrigger"
	timer            eventType = "timer"
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

type animationTriggerEvent struct {
	Event
	ControllerId    string                          `json:"controllerId"`
	AnimatedPropKey string                          `json:"animatedPropKey"`
	Method          duit_animations.AnimationMethod `json:"method"`
}

type timerEvent struct {
	Event
	Delay  int         `json:"delay"`
	TEvent interface{} `json:"event"`
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

func AnimationTriggerEvent(animatedPropKey, controllerId string, method duit_animations.AnimationMethod) *animationTriggerEvent {
	return &animationTriggerEvent{
		Event: Event{
			Type: animationTrigger,
		},
		AnimatedPropKey: animatedPropKey,
		ControllerId:    controllerId,
		Method:          method,
	}
}

func TimerEvent(delay int, event interface{}) *timerEvent {
	return &timerEvent{
		Event: Event{
			Type: timer,
		},
		Delay:  delay,
		TEvent: event,
	}
}
