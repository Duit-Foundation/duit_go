package duit_attributes

type Validatable interface {
	Validate() error
}
