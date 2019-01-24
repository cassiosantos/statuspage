package errors

const (
	ErrAlreadyExists              = "Already exists"
	ErrInvalidRef                 = "Invalid reference %s"
	ErrComponentNameIsEmpty       = "Component name shouldn't be empty"
	ErrInvalidQuery               = "Invalid query"
	ErrNotFound                   = "Not found"
	ErrComponentRefAlreadyExists  = "Component ref already exists"
	ErrComponentNameAlreadyExists = "Component name already exists"
	ErrClientRefAlreadyExists     = "Client ref already exists"
	ErrClientNameAlreadyExists    = "Client name already exists"
	ErrInvalidMonth               = "Invalid month"
	ErrInvalidYear                = "Invalid year"
	ErrTriggerUnavailable         = "Unavailable Trigger"
	ErrInvalidIncidentJSONDate    = "Field occurence_date not found"
	ErrMongoFailuere              = "Failed to perform operation on MongoDB"
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
