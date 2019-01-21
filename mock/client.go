package mock

import (
	"github.com/involvestecnologia/statuspage/client"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
	"gopkg.in/mgo.v2/bson"
)

type MockClientDAO struct {
	clients []models.Client
}

func NewMockClientDAO() client.Repository {
	return &MockClientDAO{
		clients: []models.Client{
			models.Client{
				Ref:  ZeroTimeHex,
				Name: "First Client",
				Resources: []string{
					bson.NewObjectIdWithTime(bson.Now()).Hex(),
					bson.NewObjectIdWithTime(bson.Now()).Hex(),
					bson.NewObjectIdWithTime(bson.Now()).Hex(),
				},
			},
			models.Client{
				Ref:  OneSecTimeHex,
				Name: "Last Client",
				Resources: []string{
					bson.NewObjectIdWithTime(bson.Now()).Hex(),
				},
			},
		},
	}
}

func (m *MockClientDAO) Delete(clientRef string) error {
	for i, c := range m.clients {
		if c.Ref == clientRef {
			m.clients = append(m.clients[:i], m.clients[i+1:]...)
			return nil
		}
	}
	return errors.E(errors.ErrNotFound)
}

func (m *MockClientDAO) Find(q map[string]interface{}) (models.Client, error) {
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

	return models.Client{}, errors.E(errors.ErrNotFound)
}

func (m *MockClientDAO) Insert(client models.Client) (string, error) {
	if client.Ref == "" {
		client.Ref = bson.NewObjectId().Hex()
	}
	m.clients = append(m.clients, client)
	return client.Ref, nil
}

func (m *MockClientDAO) List() ([]models.Client, error) {
	return m.clients, nil
}

func (m *MockClientDAO) Update(clientRef string, client models.Client) error {
	for i, c := range m.clients {
		if c.Ref == clientRef {
			m.clients[i].Name = client.Name
			m.clients[i].Resources = client.Resources
			return nil
		}
	}
	return errors.E(errors.ErrNotFound)
}
