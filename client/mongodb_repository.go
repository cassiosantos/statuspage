package client

import (
	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

const databaseName = "status"

type MongoRepository struct {
	db *mgo.Session
}

func NewMongoRepository(session *mgo.Session) *MongoRepository {
	return &MongoRepository{db: session}
}

func (r *MongoRepository) Insert(client models.Client) (ref string, err error) {
	defer mongoFailure(&err)
	if client.Ref == "" {
		client.Ref = bson.NewObjectId().Hex()
	}
	err = r.db.DB(databaseName).C("Client").Insert(client)
	return client.Ref, err
}

func (r *MongoRepository) Update(clientID string, client models.Client) (err error) {
	defer mongoFailure(&err)

	clientQ := bson.M{"ref": clientID}
	change := bson.M{"$set": bson.M{
		"name":      client.Name,
		"resources": client.Resources,
	}}
	err = r.db.DB(databaseName).C("Client").Update(clientQ, change)
	if err == mgo.ErrNotFound {
		return errors.E(errors.ErrNotFound)
	}
	return err
}

func (r *MongoRepository) Find(q map[string]interface{}) (client models.Client, err error) {
	defer mongoFailure(&err)
	err = r.db.DB(databaseName).C("Client").Find(q).One(&client)
	if err == mgo.ErrNotFound {
		return client, errors.E(errors.ErrNotFound)
	}
	return client, err
}

func (r *MongoRepository) Delete(clientID string) (err error) {
	defer mongoFailure(&err)
	clientQ := bson.M{"ref": clientID}
	err = r.db.DB(databaseName).C("Client").Remove(clientQ)
	if err == mgo.ErrNotFound {
		return errors.E(errors.ErrNotFound)
	}
	return err
}

func (r *MongoRepository) List() (clients []models.Client, err error) {
	defer mongoFailure(&err)
	err = r.db.DB(databaseName).C("Client").Find(nil).All(&clients)
	return clients, err
}

func mongoFailure(e *error) {
	if r := recover(); r != nil {
		*e = errors.E(errors.ErrMongoFailuere)
	}
}
