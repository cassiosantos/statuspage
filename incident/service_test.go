package incident

import (
	"fmt"
	"testing"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"

	"github.com/stretchr/testify/assert"
)

const zeroTimeHex = "886e09000000000000000000"
const oneSecTimeHex = "886e09010000000000000000"

func TestNewIncidentService(t *testing.T) {
	dao := newMockIncidentDAO()
	s := NewService(dao)
	assert.Equal(t, dao, s.repo)

}

func TestIncidentService_ListIncidents(t *testing.T) {
	s := NewService(newMockIncidentDAO())

	i, err := s.ListIncidents("", "")
	if assert.Nil(t, err) && assert.NotNil(t, i) {
		if assert.IsType(t, []models.IncidentWithComponentName{}, i) {
			assert.Equal(t, zeroTimeHex, i[0].Component)
			assert.Equal(t, oneSecTimeHex, i[len(i)-1].Component)
		}
	}

	// Valid Month
	monthOnly, err := s.ListIncidents("", "1")
	if assert.Nil(t, err) && assert.NotNil(t, monthOnly) {
		assert.Equal(t, 1, int(monthOnly[0].Incident.Date.Month()))
	}
	// Valid Year
	yearOnly, err := s.ListIncidents("1", "")
	if assert.Nil(t, err) && assert.NotNil(t, yearOnly) {
		assert.Equal(t, 1, yearOnly[0].Incident.Date.Year())
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
	s := NewService(newMockIncidentDAO())

	i := models.Incident{
		Status:      models.IncidentStatusOK,
		Description: "",
		Date:        time.Time{},
	}

	err := s.CreateIncidents(zeroTimeHex, i)
	assert.Nil(t, err)

	inc, err := s.FindIncidents(zeroTimeHex)
	if assert.Nil(t, err) && assert.NotNil(t, i) {
		fmt.Printf("%v", inc)
		assert.Equal(t, i, inc[len(inc)-1])
	}

	err = s.CreateIncidents(bson.NewObjectId().Hex(), i)
	assert.NotNil(t, err)
}
func TestIncidentService_FindIncidents(t *testing.T) {
	s := NewService(newMockIncidentDAO())

	i, err := s.FindIncidents(zeroTimeHex)
	if assert.Nil(t, err) && assert.NotNil(t, i) {
		if assert.IsType(t, []models.Incident{}, i) {
			assert.Equal(t, i[0].Status, models.IncidentStatusOK)
			assert.Equal(t, i[len(i)-1].Status, models.IncidentStatusOutage)
		}
	}

	_, err = s.FindIncidents("Invalid ID")
	assert.NotNil(t, err)

	_, err = s.FindIncidents(bson.NewObjectId().Hex())
	assert.NotNil(t, err)

}
func TestIncidentService_validateMonth(t *testing.T) {
	s := NewService(newMockIncidentDAO())

	_, err := s.validateMonth("1")
	assert.Nil(t, err)

	_, err = s.validateMonth("0")
	assert.NotNil(t, err)

	_, err = s.validateMonth("13")
	assert.NotNil(t, err)

	_, err = s.validateMonth("test")
	assert.NotNil(t, err)
}

type mockIncidentDAO struct {
	components []models.Component
}

func newMockIncidentDAO() Repository {
	return &mockIncidentDAO{
		components: []models.Component{
			models.Component{
				Ref:     zeroTimeHex,
				Name:    "first",
				Address: "",
				Incidents: []models.Incident{
					models.Incident{
						Status: models.IncidentStatusOK,
						Date:   time.Date(time.Time{}.Year(), time.Month(1), 1, 0, 0, 0, 0, time.UTC),
					},
					models.Incident{
						Status: models.IncidentStatusOutage,
						Date:   time.Date(time.Time{}.Year(), time.Month(12), 1, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			models.Component{
				Ref:     oneSecTimeHex,
				Name:    "last",
				Address: "",
				Incidents: []models.Incident{
					models.Incident{
						Status: models.IncidentStatusOK,
						Date:   time.Date(time.Time{}.Year(), time.Month(12), 1, 0, 0, 0, 0, time.UTC),
					},
					models.Incident{
						Status: models.IncidentStatusOutage,
						Date:   time.Date(time.Time{}.Year(), time.Month(1), 1, 0, 0, 0, 0, time.UTC),
					},
				},
			},
		},
	}
}

func (m *mockIncidentDAO) Find(ref string) ([]models.Incident, error) {
	var i []models.Incident

	for _, c := range m.components {
		if c.Ref == ref {
			return c.Incidents, nil
		}
	}
	return i, errors.E(fmt.Sprintf("Component with ref %s not found", ref))
}
func (m *mockIncidentDAO) List(start time.Time, end time.Time) ([]models.IncidentWithComponentName, error) {
	var inc []models.IncidentWithComponentName
	for _, c := range m.components {
		for _, i := range c.Incidents {
			inc = append(inc, models.IncidentWithComponentName{
				Component: c.Ref,
				Incident:  i,
			})
		}
	}
	return inc, nil
}
func (m *mockIncidentDAO) Insert(componentRef string, incident models.Incident) error {
	if !bson.IsObjectIdHex(componentRef) {
		return errors.E(fmt.Sprintf("%s is a invalid Ref", componentRef))
	}
	for index, c := range m.components {
		if c.Ref == componentRef {
			m.components[index].Incidents = append(c.Incidents, incident)
			return nil
		}
	}
	return errors.E(fmt.Sprintf("Component with id %s not found", componentRef))
}
