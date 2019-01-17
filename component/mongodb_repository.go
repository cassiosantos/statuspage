package component

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

func (r *MongoRepository) Insert(comp models.Component) (string, error) {
	if comp.Ref == "" {
		comp.Ref = bson.NewObjectId().Hex()
	}
	err := r.db.DB(databaseName).C("Component").Insert(comp)
	if err != nil {
		log.Panicf("Error inserting component: %s\n", err)
	}
	return comp.Ref, err
}

func (r *MongoRepository) List() ([]models.Component, error) {
	var list []models.Component
	err := r.db.DB(databaseName).C("Component").Find(nil).All(&list)
	if err != nil {
		log.Panicf("Error searching components: %s\n", err)
	}
	return list, err
}

func (r *MongoRepository) Find(q map[string]interface{}) (models.Component, error) {
	var c models.Component
	err := r.db.DB(databaseName).C("Component").Find(q).One(&c)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error searching component: %s\n", err)
	}
	return c, err
}

func (r *MongoRepository) Update(ref string, comp models.Component) error {
	compQ := bson.M{"ref": ref}
	change := bson.M{"$set": bson.M{
		"name":    comp.Name,
		"address": comp.Address,
	}}
	err := r.db.DB(databaseName).C("Component").Update(compQ, change)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error updating component: %s\n", err)
	}
	return err
}

func (r *MongoRepository) Delete(ref string) error {
	compQ := bson.M{"ref": ref}
	err := r.db.DB(databaseName).C("Component").Remove(compQ)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error deleting component: %s\n", err)
	}
	return err
}
