package errors

const (
	ErrInvalidRef              = "Invalid Ref, the reference %s is already in use"
	ErrInvalidQuery            = "Invalid query"
	ErrNotFound                = "Not found"
	ErrAlreadyExists           = "%s already exists"
	ErrInvalidMonth            = "Invalid month"
	ErrInvalidYear             = "Invalid year"
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
