package component

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

func (r *MongoRepository) GetCurrentSession() *mgo.Session {
	return r.db
}

func (r *MongoRepository) AddComponent(comp models.Component) error {
	return r.db.DB("status").C("Component").Insert(comp)
}

func (r *MongoRepository) GetAllComponents() ([]models.Component, error) {
	list := []models.Component{}
	err := r.db.DB("status").C("Component").Find(nil).All(&list)
	if err != nil {
		log.Panicf("Error searching component: %s\n", err)
	}
	return list, err
}

func (r *MongoRepository) GetComponentsByGroup(groupName string) ([]models.Component, error) {
	var list []models.Component
	compQ := bson.M{"groups": bson.M{"$elemMatch": bson.M{"$eq": groupName}}}

	err := r.db.DB("status").C("Component").Find(compQ).All(&list)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error searching component: %s\n", err)
	}
	return list, err
}

func (r *MongoRepository) GetComponentById(id string) (models.Component, error) {
	var c models.Component
	if !bson.IsObjectIdHex(id) {
		return c, errors.E(errors.ErrInvalidHexID)
	}
	compQ := bson.M{"_id": bson.ObjectIdHex(id)}
	err := r.db.DB("status").C("Component").Find(compQ).One(&c)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error searching component: %s\n", err)
	}
	return c, err
}

func (r *MongoRepository) GetComponentByName(name string) (models.Component, error) {
	var c models.Component
	compQ := bson.M{"name": name}
	err := r.db.DB("status").C("Component").Find(compQ).One(&c)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error searching component: %s\n", err)
	}
	return c, err
}

func (r *MongoRepository) UpdateComponent(id string, comp models.Component) error {
	if !bson.IsObjectIdHex(id) {
		return errors.E(errors.ErrInvalidHexID)
	}
	compQ := bson.M{"_id": bson.ObjectIdHex(id)}
	change := bson.M{"$set": bson.M{
		"name":      comp.Name,
		"groups":    comp.Groups,
		"incidents": comp.Incidents,
		"address":   comp.Address,
	}}
	err := r.db.DB("status").C("Component").Update(compQ, change)
	if err != nil {
		log.Panicf("Error updating component: %s\n", err)
	}
	return nil
}

func (r *MongoRepository) DeleteComponent(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.E(errors.ErrInvalidHexID)
	}
	compQ := bson.M{"_id": bson.ObjectIdHex(id)}
	err := r.db.DB("status").C("Component").Remove(compQ)
	if err != nil {
		log.Panicf("Error deleting component: %s\n", err)
	}
	return nil
}
