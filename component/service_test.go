package component

import (
	"fmt"
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/involvestecnologia/statuspage/errors"

	"github.com/involvestecnologia/statuspage/models"

	"github.com/stretchr/testify/assert"
)

const zeroTimeHex = "886e09000000000000000000"
const oneSecTimeHex = "886e09010000000000000000"

func TestNewComponentService(t *testing.T) {
	dao := newMockComponentDAO()
	s := NewService(dao)
	assert.Equal(t, dao, s.repo)

}

func TestComponentService_ListComponents(t *testing.T) {
	s := NewService(newMockComponentDAO())

	c, err := s.ListComponents()
	if assert.Nil(t, err) && assert.NotNil(t, c) {
		if assert.IsType(t, []models.Component{}, c) {
			assert.Equal(t, "first", c[0].Name)
			assert.Equal(t, "last", c[len(c)-1].Name)
		}
	}
}
func TestComponentService_FindComponent(t *testing.T) {
	s := NewService(newMockComponentDAO())

	c, err := s.FindComponent(map[string]interface{}{"ref": zeroTimeHex})
	if assert.Nil(t, err) && assert.NotNil(t, c) {
		if assert.IsType(t, models.Component{}, c) {
			assert.Equal(t, "first", c.Name)
		}
	}

	c2, err := s.FindComponent(map[string]interface{}{"name": c.Name})
	if assert.Nil(t, err) && assert.NotNil(t, c) {
		if assert.IsType(t, models.Component{}, c) {
			assert.Equal(t, c, c2)
		}
	}

	_, err = s.FindComponent(map[string]interface{}{"ref": oneSecTimeHex})
	assert.NotNil(t, err)

	_, err = s.FindComponent(map[string]interface{}{"name": oneSecTimeHex})
	assert.NotNil(t, err)

	_, err = s.FindComponent(map[string]interface{}{"ref": "test"})
	assert.NotNil(t, err)

	_, err = s.FindComponent(map[string]interface{}{"name": ""})
	assert.NotNil(t, err)

	_, err = s.FindComponent(map[string]interface{}{"ref": ""})
	assert.NotNil(t, err)

	_, err = s.FindComponent(map[string]interface{}{})
	assert.NotNil(t, err)

}
func TestComponentService_CreateComponent(t *testing.T) {
	s := NewService(newMockComponentDAO())

	c := models.Component{
		Ref:       bson.NewObjectIdWithTime(bson.Now()).Hex(),
		Name:      "New Component",
		Address:   "no-address",
		Incidents: make([]models.Incident, 0),
	}

	ref, err := s.CreateComponent(c)
	c.Ref = ref
	assert.Nil(t, err)

	comp, err := s.FindComponent(map[string]interface{}{"name": c.Name})
	if assert.Nil(t, err) && assert.NotNil(t, comp) {
		if assert.IsType(t, models.Component{}, comp) {
			assert.Equal(t, c, comp)
		}
	}

	c.Ref = zeroTimeHex
	_, err = s.CreateComponent(c)
	assert.NotNil(t, err)

}
func TestComponentService_UpdateComponent(t *testing.T) {
	s := NewService(newMockComponentDAO())

	currTime := bson.Now().String()
	c, err := s.FindComponent(map[string]interface{}{"ref": zeroTimeHex})
	if assert.Nil(t, err) && assert.NotNil(t, c) {
		c.Address = currTime
	}
	err = s.UpdateComponent(zeroTimeHex, c)
	assert.Nil(t, err)

	comp, err := s.FindComponent(map[string]interface{}{"ref": zeroTimeHex})
	if assert.Nil(t, err) && assert.NotNil(t, comp) {
		assert.Equal(t, c, comp)
	}

	err = s.UpdateComponent(oneSecTimeHex, c)
	assert.NotNil(t, err)

}
func TestComponentService_RemoveComponent(t *testing.T) {
	s := NewService(newMockComponentDAO())

	err := s.RemoveComponent(zeroTimeHex)
	assert.Nil(t, err)

	_, err = s.FindComponent(map[string]interface{}{"ref": zeroTimeHex})
	assert.NotNil(t, err)
}
func TestComponentService_componentExists(t *testing.T) {
	s := NewService(newMockComponentDAO())

	c, err := s.FindComponent(map[string]interface{}{"ref": zeroTimeHex})
	if assert.Nil(t, err) && assert.NotNil(t, c) {
		if assert.IsType(t, models.Component{}, c) {
			assert.Equal(t, "first", c.Name)
		}
	}
	_, exists := s.componentExists(c.Name)
	assert.True(t, exists)

	_, exists = s.componentExists(bson.NewObjectIdWithTime(bson.Now()).String())
	assert.False(t, exists)
}

type mockComponentDAO struct {
	components []models.Component
}

func newMockComponentDAO() Repository {
	return &mockComponentDAO{
		components: []models.Component{
			models.Component{
				Ref:       zeroTimeHex,
				Name:      "first",
				Address:   "",
				Incidents: make([]models.Incident, 0),
			},
			models.Component{
				Ref:       bson.NewObjectIdWithTime(bson.Now()).Hex(),
				Name:      "first_comp_with_group",
				Address:   "",
				Incidents: make([]models.Incident, 0),
			},
			models.Component{
				Ref:       bson.NewObjectIdWithTime(bson.Now()).Hex(),
				Name:      "test",
				Address:   "",
				Incidents: make([]models.Incident, 0),
			},
			models.Component{
				Ref:       bson.NewObjectIdWithTime(bson.Now()).Hex(),
				Name:      "last_comp_with_group",
				Address:   "",
				Incidents: make([]models.Incident, 0),
			},
			models.Component{
				Ref:       bson.NewObjectIdWithTime(bson.Now()).Hex(),
				Name:      "last",
				Address:   "",
				Incidents: make([]models.Incident, 0),
			},
		},
	}
}

func (m *mockComponentDAO) List() ([]models.Component, error) {
	return m.components, nil
}
func (m *mockComponentDAO) Find(q map[string]interface{}) (models.Component, error) {
	var c models.Component
	if keyValue, hasKey := q["ref"]; hasKey {
		for _, c := range m.components {
			if c.Ref == keyValue {
				return c, nil
			}
		}
	} else {
		if keyValue, hasKey := q["name"]; hasKey {
			for _, c := range m.components {
				if c.Name == keyValue {
					return c, nil
				}
			}
		} else {
			return c, errors.E("No query parameters passed")
		}
	}

	return c, errors.E("Component not found")
}
func (m *mockComponentDAO) Insert(component models.Component) (string, error) {
	if component.Ref == "" {
		component.Ref = bson.NewObjectId().Hex()
	}
	m.components = append(m.components, component)
	return component.Ref, nil
}
func (m *mockComponentDAO) Update(ref string, component models.Component) error {
	for k, comp := range m.components {
		if comp.Ref == ref {
			m.components[k] = component
			return nil
		}
	}
	return errors.E(fmt.Sprintf("Component with ref %s not found", ref))
}
func (m *mockComponentDAO) Delete(ref string) error {
	for k, comp := range m.components {
		if comp.Ref == ref {
			m.components = append(m.components[:k], m.components[k+1:]...)
			return nil
		}
	}
	return errors.E(fmt.Sprintf("Component with ref %s not found", ref))
}
