package emailvalidator

import "errors"

var (
	// ErrEmptyLocalPart is returned when a local part is empty
	ErrEmptyLocalPart = errors.New("invalid local part, cannot be empty")
	// ErrEmptyDomain is returned when a local part is empty
	ErrEmptyDomain = errors.New("invalid local part, cannot be empty")
	// ErrLocalPartLeadingDot is returned when a local part is led with a period
	ErrLocalPartLeadingDot = errors.New("invalid local part, cannot lead with \".\"")
	// ErrLocalPartTrailingDot is returned when a local part is ended with a period
	ErrLocalPartTrailingDot = errors.New("invalid local part, cannot end with \".\"")
	// ErrLocalPartMultiplePeriods is returned when a local part has periods which immediately repeat
	ErrLocalPartMultiplePeriods = errors.New("invalid local part, cannot have immediately repeating periods")
)

// Validate will validate a provided email address
func Validate(emailAddress string) (err error) {
	var e *EmailAddress
	if e, err = NewEmailAddress(emailAddress); err != nil {
		return
	}

	return e.Validate()
}
