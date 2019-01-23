package component

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

func (r *MongoRepository) Insert(comp models.Component) (compRef string, err error) {
	defer mongoFailure(&err)
	if comp.Ref == "" {
		comp.Ref = bson.NewObjectId().Hex()
	}
	compRef = comp.Ref
	err = r.db.DB(databaseName).C("Component").Insert(comp)
	return compRef, err
}

func (r *MongoRepository) List() (compList []models.Component, err error) {
	defer mongoFailure(&err)
	err = r.db.DB(databaseName).C("Component").Find(nil).All(&compList)
	return compList, err
}

func (r *MongoRepository) Find(q map[string]interface{}) (c models.Component, err error) {
	defer mongoFailure(&err)
	err = r.db.DB(databaseName).C("Component").Find(q).One(&c)
	if err == mgo.ErrNotFound {
		return c, errors.E(errors.ErrNotFound)
	}
	return c, err
}

func (r *MongoRepository) Update(ref string, comp models.Component) (err error) {
	defer mongoFailure(&err)
	compQ := bson.M{"ref": ref}
	change := bson.M{"$set": bson.M{
		"name":    comp.Name,
		"address": comp.Address,
	}}
	err = r.db.DB(databaseName).C("Component").Update(compQ, change)
	if err == mgo.ErrNotFound {
		return errors.E(errors.ErrNotFound)
	}
	return err
}

func (r *MongoRepository) Delete(ref string) (err error) {
	defer mongoFailure(&err)
	compQ := bson.M{"ref": ref}
	err = r.db.DB(databaseName).C("Component").Remove(compQ)
	if err == mgo.ErrNotFound {
		return errors.E(errors.ErrNotFound)
	}
	return err
}

func mongoFailure(e *error) {
	if r := recover(); r != nil {
		*e = errors.E(errors.ErrMongoFailuere)
	}
}
