package client

import (
	"log"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/involvestecnologia/statuspage/models"
)

const databaseName = "status"

type MongoRepository struct {
	db *mgo.Session
}

func NewMongoRepository(session *mgo.Session) *MongoRepository {
	return &MongoRepository{db: session}
}

func (r *MongoRepository) Insert(client models.Client) (string, error) {
	if client.Ref == "" {
		client.Ref = bson.NewObjectId().Hex()
	}
	err := r.db.DB(databaseName).C("Client").Insert(client)
	if err != nil {
		log.Panicf("Error inserting client: %s\n", err)
	}
	return client.Ref, err
}

func (r *MongoRepository) Update(clientID string, client models.Client) error {
	clientQ := bson.M{"ref": clientID}
	change := bson.M{"$set": bson.M{
		"name":      client.Name,
		"resources": client.Resources,
	}}
	err := r.db.DB(databaseName).C("Client").Update(clientQ, change)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error updating client: %s\n", err)
	}
	return err
}

func (r *MongoRepository) Find(q map[string]interface{}) (models.Client, error) {
	var client models.Client
	err := r.db.DB(databaseName).C("Client").Find(q).One(&client)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error searching client: %s\n", err)
	}
	return client, err
}

func (r *MongoRepository) Delete(clientID string) error {
	clientQ := bson.M{"ref": clientID}
	err := r.db.DB(databaseName).C("Client").Remove(clientQ)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error deleting client: %s\n", err)
	}
	return err
}

func (r *MongoRepository) List() ([]models.Client, error) {
	var clients []models.Client
	err := r.db.DB(databaseName).C("Client").Find(nil).All(&clients)
	if err != nil {
		log.Panicf("Error searching clients: %s\n", err)
	}
	return clients, err
}
