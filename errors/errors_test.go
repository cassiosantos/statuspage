package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB_Error(t *testing.T) {

	errConsts := []string{
		ErrAlreadyExists,
		ErrInvalidRef,
		ErrComponentNameIsEmpty,
		ErrInvalidQuery,
		ErrNotFound,
		ErrComponentRefAlreadyExists,
		ErrComponentNameAlreadyExists,
		ErrClientRefAlreadyExists,
		ErrClientNameAlreadyExists,
		ErrInvalidMonth,
		ErrInvalidYear,
		ErrTriggerUnavailable,
		ErrInvalidIncidentJSONDate,
		ErrMongoFailuere,
	}

	for _, e := range errConsts {
		err := E(e)
		if assert.NotNil(t, err) && assert.IsType(t, &errorMsg{}, err) {
			assert.Equal(t, e, err.Error())
		}
	}

}
