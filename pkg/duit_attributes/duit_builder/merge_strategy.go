package duit_builder

type MergeStrategy uint8

const (
	AddToEnd MergeStrategy = iota
	AddToStart
	Replace
)
