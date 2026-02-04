package duit_props

type TimeOfDay = [2]uint8 // [hour, minute]

func NewTimeOfDay(hour, minute uint8) *TimeOfDay {
	return &TimeOfDay{hour, minute}
}