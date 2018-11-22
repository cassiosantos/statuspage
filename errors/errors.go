package errors

const (
	ErrInvalidHexID            = "Invalid ID"
	ErrInvalidMonth            = "Invalid month"
	ErrTriggerUnavailable      = "Unavailable Trigger"
	ErrInvalidIncidentJSONDate = "Field occurence_date not found"
)

func E(msg string) error {
	return &errorMsg{msg}
}

type errorMsg struct {
	msg string
}

func (e *errorMsg) Error() string {
	return e.msg
}
