package incident

import (
	"time"

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

func (r *MongoRepository) Insert(incident models.Incident) (err error) {
	defer mongoFailure(&err)
	err = r.db.DB(databaseName).C("Incidents").Insert(incident)
	return err
}

func (r *MongoRepository) Find(queryParam map[string]interface{}) (incidents []models.Incident, err error) {
	defer mongoFailure(&err)
	err = r.db.DB(databaseName).C("Incidents").Find(queryParam).All(&incidents)
	if incidents == nil {
		return incidents, errors.E(errors.ErrNotFound)
	}
	return incidents, err
}

func (r *MongoRepository) List(startDt time.Time, endDt time.Time) (incidents []models.Incident, err error) {
	defer mongoFailure(&err)
	findQ := bson.M{
		"date": bson.M{
			"$gt": startDt.Add(-(24 * time.Hour)),
			"$lt": endDt.Add(24 * time.Hour),
		},
	}

	err = r.db.DB(databaseName).C("Incidents").Find(findQ).All(&incidents)
	return incidents, err
}

func mongoFailure(e *error) {
	if r := recover(); r != nil {
		*e = errors.E(errors.ErrMongoFailuere)
	}
}
