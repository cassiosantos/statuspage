package incident

import (
	"time"

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

func (r *MongoRepository) AddIncidentToComponent(componentID string, incident models.Incident) error {
	if !bson.IsObjectIdHex(componentID) {
		return errors.E(errors.ErrInvalidHexID)
	}
	findCompQ := bson.M{"_id": bson.ObjectIdHex(componentID)}
	insertIncidentQ := bson.M{"$push": bson.M{"incidents": incident}}

	return r.db.DB("status").C("Component").Update(findCompQ, insertIncidentQ)
}

func (r *MongoRepository) GetIncidentsByComponentID(componentID string) ([]models.Incident, error) {
	var component models.Component
	if !bson.IsObjectIdHex(componentID) {
		return component.Incidents, errors.E(errors.ErrInvalidHexID)
	}

	findCompQ := bson.M{"_id": bson.ObjectIdHex(componentID)}
	incidentQ := bson.M{"incidents": bson.M{"$not": bson.M{"$size": 0}}}
	query := bson.M{"$and": []bson.M{findCompQ, incidentQ}}
	incidentFilter := bson.M{"incidents": 1}

	err := r.db.DB("status").C("Component").Find(query).Select(incidentFilter).One(&component)
	return component.Incidents, err
}

func (r *MongoRepository) GetAllIncidents() ([]models.IncidentWithComponentID, error) {
	var incidents []models.IncidentWithComponentID

	unwind := bson.M{"$unwind": "$incident"}

	project := bson.M{"$project": bson.M{
		"component": "$name",
		"incident":  "$incidents",
	}}

	pipeline := []bson.M{project, unwind}

	err := r.db.DB("status").C("Component").Pipe(pipeline).All(&incidents)
	return incidents, err
}

func (r *MongoRepository) GetIncidentsByMonth(month int) ([]models.IncidentWithComponentID, error) {
	var incidents []models.IncidentWithComponentID

	year := time.Now().Year()
	fromDate := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(year, time.Month(month+2), 0, 23, 59, 59, 0, time.UTC)
	unwind := bson.M{"$unwind": "$incident"}

	project := bson.M{"$project": bson.M{
		"component": "$name",
		"incident":  "$incidents",
	}}

	match := bson.M{"$match": bson.M{
		"incident.date": bson.M{
			"$gt": fromDate,
			"$lt": toDate,
		},
	}}

	pipeline := []bson.M{project, unwind, match}

	err := r.db.DB("status").C("Component").Pipe(pipeline).All(&incidents)
	return incidents, err
}
