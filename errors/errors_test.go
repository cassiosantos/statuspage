package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB_Error(t *testing.T) {

	errConsts := []string{
		ErrInvalidRef,
		ErrInvalidQuery,
		ErrNotFound,
		ErrAlreadyExists,
		ErrInvalidMonth,
		ErrInvalidYear,
		ErrTriggerUnavailable,
		ErrInvalidIncidentJSONDate,
	}

	for _, e := range errConsts {
		err := E(e)
		if assert.NotNil(t, err) && assert.IsType(t, &errorMsg{}, err) {
			assert.Equal(t, e, err.Error())
		}
	}

}
