package client

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

func TestNewClientService(t *testing.T) {
	dao := newMockClientDAO()
	s := NewService(dao)
	assert.Equal(t, dao, s.repo)

}

func TestClientService_CreateClient(t *testing.T) {
	s := NewService(newMockClientDAO())
	c := models.Client{
		Name:      "test",
		Resources: make([]string, 0),
	}

	ref, err := s.CreateClient(c)
	c.Ref = ref

	assert.Nil(t, err)
	c2, err := s.FindClient(map[string]interface{}{"ref": c.Ref})
	if assert.Nil(t, err) && assert.NotNil(t, c2) {
		if assert.IsType(t, models.Client{}, c2) {
			assert.Equal(t, c, c2)
		}
	}

	_, err = s.CreateClient(c)
	assert.NotNil(t, err)
}
func TestClientService_FindClient(t *testing.T) {
	s := NewService(newMockClientDAO())

	c, err := s.FindClient(map[string]interface{}{"ref": zeroTimeHex})
	if assert.Nil(t, err) && assert.NotNil(t, c) {
		if assert.IsType(t, models.Client{}, c) {
			assert.Equal(t, zeroTimeHex, c.Ref)
		}
	}

	_, err = s.FindClient(map[string]interface{}{"ref": bson.NewObjectId().Hex()})
	assert.NotNil(t, err)

	c, err = s.FindClient(map[string]interface{}{"name": "First Client"})
	if assert.Nil(t, err) && assert.NotNil(t, c) {
		if assert.IsType(t, models.Client{}, c) {
			assert.Equal(t, "First Client", c.Name)
		}
	}

	_, err = s.FindClient(map[string]interface{}{"name": "test"})
	assert.NotNil(t, err)

}
func TestClientService_ListClients(t *testing.T) {
	s := NewService(newMockClientDAO())

	c, err := s.ListClients()
	if assert.Nil(t, err) && assert.NotNil(t, c) {
		if assert.IsType(t, []models.Client{}, c) {
			assert.Equal(t, zeroTimeHex, c[0].Ref)
			assert.Equal(t, oneSecTimeHex, c[len(c)-1].Ref)
		}
	}
}
func TestClientService_RemoveClient(t *testing.T) {
	s := NewService(newMockClientDAO())
	err := s.RemoveClient(oneSecTimeHex)
	assert.Nil(t, err)

	_, err = s.FindClient(map[string]interface{}{"ref": oneSecTimeHex})
	assert.NotNil(t, err)
}
func TestClientService_UpdateClient(t *testing.T) {
	s := NewService(newMockClientDAO())

	c, err := s.FindClient(map[string]interface{}{"ref": zeroTimeHex})
	assert.Nil(t, err)
	assert.NotNil(t, c)

	c.Name = "Modified First Client"

	err = s.UpdateClient(zeroTimeHex, c)
	assert.Nil(t, err)

	c2, err := s.FindClient(map[string]interface{}{"ref": zeroTimeHex})
	assert.Nil(t, err)
	assert.NotNil(t, c2)

	assert.Equal(t, c, c2)
}

type mockClientDAO struct {
	clients []models.Client
}

func newMockClientDAO() Repository {
	return &mockClientDAO{
		clients: []models.Client{
			models.Client{
				Ref:  zeroTimeHex,
				Name: "First Client",
				Resources: []string{
					bson.NewObjectIdWithTime(bson.Now()).Hex(),
					bson.NewObjectIdWithTime(bson.Now()).Hex(),
					bson.NewObjectIdWithTime(bson.Now()).Hex(),
				},
			},
			models.Client{
				Ref:  oneSecTimeHex,
				Name: "Last Client",
				Resources: []string{
					bson.NewObjectIdWithTime(bson.Now()).Hex(),
				},
			},
		},
	}
}

func (m *mockClientDAO) Delete(clientRef string) error {
	for i, c := range m.clients {
		if c.Ref == clientRef {
			m.clients = append(m.clients[:i], m.clients[i+1:]...)
			return nil
		}
	}
	return errors.E(fmt.Sprintf("Client with id %s was not found", clientRef))
}

func (m *mockClientDAO) Find(q map[string]interface{}) (models.Client, error) {
	if keyValue, hasKey := q["ref"]; hasKey {
		for _, c := range m.clients {
			if c.Ref == keyValue {
				return c, nil
			}
		}
	} else {
		if keyValue, hasKey := q["name"]; hasKey {
			for _, c := range m.clients {
				if c.Name == keyValue {
					return c, nil
				}
			}
		} else {
			return models.Client{}, errors.E("No queryable parameters passed")
		}
	}

	return models.Client{}, errors.E("Client not found")
}

func (m *mockClientDAO) Insert(client models.Client) (string, error) {
	if client.Ref == "" {
		client.Ref = bson.NewObjectId().Hex()
	}
	m.clients = append(m.clients, client)
	return client.Ref, nil
}

func (m *mockClientDAO) List() ([]models.Client, error) {
	return m.clients, nil
}

func (m *mockClientDAO) Update(clientRef string, client models.Client) error {
	for i, c := range m.clients {
		if c.Ref == clientRef {
			m.clients[i] = client
			return nil
		}
	}
	return errors.E(fmt.Sprintf("Client with id %s was not found", clientRef))
}
