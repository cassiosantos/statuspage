package incident

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

const databaseName = "status"

type mongoRepository struct {
	db *mgo.Session
}

//NewMongoRepository creates a new Repository implementation using the MongoDB as database
func NewMongoRepository(session *mgo.Session) Repository {
	return &mongoRepository{db: session}
}

func (r *mongoRepository) Insert(incident models.Incident) (err error) {
	defer mongoFailure(&err)
	err = r.db.DB(databaseName).C("Incidents").Insert(incident)
	return err
}

func (r *mongoRepository) Find(queryParam map[string]interface{}) (incidents []models.Incident, err error) {
	defer mongoFailure(&err)
	err = r.db.DB(databaseName).C("Incidents").Find(queryParam).Sort("-date").All(&incidents)
	if incidents == nil {
		return incidents, &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
	}
	return incidents, err
}

func (r *mongoRepository) FindOne(queryParam map[string]interface{}) (incident models.Incident, err error) {
	defer mongoFailure(&err)
	err = r.db.DB(databaseName).C("Incidents").Find(queryParam).Sort("-date").Limit(1).One(&incident)
	if err != nil {
		return incident, &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
	}
	return incident, err
}

func (r *mongoRepository) Update(incident models.Incident) (err error) {
	defer mongoFailure(&err)
	var i models.Incident
	field := bson.M{
		"component_ref": incident.ComponentRef,
		"description":   incident.Description,
		"resolved":      false,
	}
	change := mgo.Change{
		Update:    bson.M{"$set": incident},
		ReturnNew: false,
	}
	_, err = r.db.DB(databaseName).C("Incidents").Find(field).Sort("-date").Apply(change, &i)
	if err != nil {
		return &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
	}
	return err
}

func (r *mongoRepository) List(startDt time.Time, endDt time.Time, unresolved bool) (incidents []models.Incident, err error) {
	defer mongoFailure(&err)
	findQ := bson.M{
		"date": bson.M{
			"$gt": startDt,
			"$lt": endDt,
		},
	}

	if unresolved {
		findQ = bson.M{"$and": []bson.M{findQ, {"resolved": false}}}
	}

	err = r.db.DB(databaseName).C("Incidents").Find(findQ).Sort("-date").All(&incidents)
	return incidents, err
}

func mongoFailure(e *error) {
	if r := recover(); r != nil {
		*e = &errors.ErrMongoFailure{Message: errors.ErrMongoFailureMessage}
	}
}
