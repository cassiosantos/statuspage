package client

import (
	"log"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type MongoRepository struct {
	db *mgo.Session
}

func NewMongoRepository(session *mgo.Session) *MongoRepository {
	return &MongoRepository{db: session}
}

func (r *MongoRepository) AddClient(client models.Client) error {
	return r.db.DB("status").C("Client").Insert(client)
}

func (r *MongoRepository) UpdateClient(clientID string, client models.Client) error {
	if !bson.IsObjectIdHex(clientID) {
		return errors.E(errors.ErrInvalidHexID)
	}
	clientQ := bson.M{"_id": bson.ObjectIdHex(clientID)}
	change := bson.M{"$set": bson.M{
		"name":      client.Name,
		"resources": client.Resources,
	}}
	err := r.db.DB("status").C("Client").Update(clientQ, change)
	if err != nil {
		log.Panicf("Error updating client: %s\n", err)
	}
	return nil
}

func (r *MongoRepository) FindById(clientID string) (models.Client, error) {
	var client models.Client
	clientQ := bson.M{"_id": bson.ObjectIdHex(clientID)}
	err := r.db.DB("status").C("Client").Find(clientQ).One(&client)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error searching client: %s\n", err)
	}
	return client, err
}

func (r *MongoRepository) FindByName(clientID string) (models.Client, error) {
	var client models.Client
	clientQ := bson.M{"name": clientID}
	err := r.db.DB("status").C("Client").Find(clientQ).One(&client)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error searching client: %s\n", err)
	}
	return client, err
}

func (r *MongoRepository) DeleteClient(clientID string) error {
	if !bson.IsObjectIdHex(clientID) {
		return errors.E(errors.ErrInvalidHexID)
	}
	clientQ := bson.M{"_id": bson.ObjectIdHex(clientID)}
	err := r.db.DB("status").C("Client").Remove(clientQ)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error deleting client: %s\n", err)
	}
	return err
}

func (r *MongoRepository) ListClients() ([]models.Client, error) {
	var clients []models.Client
	err := r.db.DB("status").C("Client").Find(nil).All(&clients)
	if err != nil {
		log.Panicf("Error fetching clients: %s\n", err)
	}
	return clients, err
}
