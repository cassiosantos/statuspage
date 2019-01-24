package incident_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/involvestecnologia/statuspage/component"

	"github.com/globalsign/mgo/bson"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/mock"
	"github.com/involvestecnologia/statuspage/models"

	"github.com/stretchr/testify/assert"
)

func TestNewIncidentService(t *testing.T) {
	incidentDAO := mock.NewMockIncidentDAO()
	componentService := component.NewService(mock.NewMockComponentDAO())
	s := incident.NewService(incidentDAO, componentService)
	assert.Implements(t, (*incident.Service)(nil), s)
}

func TestIncidentService_ListIncidents(t *testing.T) {
	s := incident.NewService(
		mock.NewMockIncidentDAO(),
		component.NewService(mock.NewMockComponentDAO()),
	)

	i, err := s.ListIncidents("", "")
	if assert.Nil(t, err) && assert.NotNil(t, i) {
		if assert.IsType(t, []models.Incident{}, i) {
			assert.Equal(t, mock.ZeroTimeHex, i[0].ComponentRef)
			assert.Equal(t, mock.OneSecTimeHex, i[len(i)-1].ComponentRef)
		}
	}

	// Valid Month
	monthOnly, err := s.ListIncidents("", "1")
	if assert.Nil(t, err) && assert.NotNil(t, monthOnly) {
		assert.Equal(t, 1, int(monthOnly[0].Date.Month()))
	}
	// Valid Year
	yearOnly, err := s.ListIncidents("1", "")
	if assert.Nil(t, err) && assert.NotNil(t, yearOnly) {
		assert.Equal(t, 1, yearOnly[0].Date.Year())
	}

	// Invalid Year
	_, err = s.ListIncidents("0", "1")
	assert.NotNil(t, err)

	// Invalid Month
	_, err = s.ListIncidents("1", "0")
	assert.NotNil(t, err)

	// Invalid Month
	_, err = s.ListIncidents("", "13")
	assert.NotNil(t, err)

	// Invalid Month
	_, err = s.ListIncidents("", "test")
	assert.NotNil(t, err)

	// Invalid Year
	_, err = s.ListIncidents("test", "")
	assert.NotNil(t, err)

	// Invalid Year
	_, err = s.ListIncidents("-1", "")
	assert.NotNil(t, err)

}
func TestIncidentService_CreateIncident(t *testing.T) {
	s := incident.NewService(
		mock.NewMockIncidentDAO(),
		component.NewService(mock.NewMockComponentDAO()),
	)

	i := models.Incident{
		ComponentRef: mock.ZeroTimeHex,
		Status:       models.IncidentStatusOutage,
		Description:  "",
		Date:         time.Time{},
	}

	err := s.CreateIncidents(i)
	assert.Nil(t, err)

	inc, err := s.FindIncidents(map[string]interface{}{"component_ref": mock.ZeroTimeHex})
	if assert.Nil(t, err) && assert.NotNil(t, i) {
		fmt.Printf("%v", inc)
		assert.Equal(t, i, inc[len(inc)-1])
	}

	i.ComponentRef = "Invalid Component Ref"
	err = s.CreateIncidents(i)
	assert.NotNil(t, err)
}
func TestIncidentService_FindIncidents(t *testing.T) {
	s := incident.NewService(
		mock.NewMockIncidentDAO(),
		component.NewService(mock.NewMockComponentDAO()),
	)

	i, err := s.FindIncidents(map[string]interface{}{"component_ref": mock.ZeroTimeHex})
	if assert.Nil(t, err) && assert.NotNil(t, i) {
		if assert.IsType(t, []models.Incident{}, i) {
			assert.Equal(t, i[0].Status, models.IncidentStatusOK)
			assert.Equal(t, i[len(i)-1].Status, models.IncidentStatusOK)
		}
	}

	_, err = s.FindIncidents(map[string]interface{}{"component_ref": "Invalid Component Ref"})
	assert.NotNil(t, err)

	_, err = s.FindIncidents(map[string]interface{}{"component_ref": bson.NewObjectId().Hex()})
	assert.NotNil(t, err)

}
func TestIncidentService_ValidateMonth(t *testing.T) {
	s := incident.NewService(
		mock.NewMockIncidentDAO(),
		component.NewService(mock.NewMockComponentDAO()),
	)

	_, err := s.ValidateMonth("1")
	assert.Nil(t, err)

	_, err = s.ValidateMonth("0")
	assert.NotNil(t, err)

	_, err = s.ValidateMonth("13")
	assert.NotNil(t, err)

	_, err = s.ValidateMonth("test")
	assert.NotNil(t, err)
}