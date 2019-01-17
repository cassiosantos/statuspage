package incident

import (
	"log"
	"time"

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

func (r *MongoRepository) Insert(componentRef string, incident models.Incident) error {
	findCompQ := bson.M{"ref": componentRef}
	insertIncidentQ := bson.M{"$push": bson.M{"incidents": incident}}
	err := r.db.DB(databaseName).C("Component").Update(findCompQ, insertIncidentQ)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error updating component: %s\n", err)
	}
	return err
}

func (r *MongoRepository) Find(componentRef string) ([]models.Incident, error) {
	var component models.Component

	findCompQ := bson.M{"ref": componentRef}
	incidentQ := bson.M{"incidents": bson.M{"$not": bson.M{"$size": 0}}}
	query := bson.M{"$and": []bson.M{findCompQ, incidentQ}}
	incidentFilter := bson.M{"incidents": 1}

	err := r.db.DB(databaseName).C("Component").Find(query).Select(incidentFilter).One(&component)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error finding component: %s\n", err)
	}
	return component.Incidents, err
}

func (r *MongoRepository) List(startDt time.Time, endDt time.Time) ([]models.IncidentWithComponentName, error) {
	var incidents []models.IncidentWithComponentName

	unwind := bson.M{"$unwind": "$incident"}

	project := bson.M{"$project": bson.M{
		"component": "$name",
		"incident":  "$incidents",
	}}

	pipeline := []bson.M{project, unwind}
	defaultTime := time.Time{}
	if defaultTime != startDt {
		match := bson.M{"$match": bson.M{
			"incident.date": bson.M{
				"$gt": startDt.Add(-(24 * time.Hour)),
				"$lt": endDt.Add(24 * time.Hour),
			},
		}}
		pipeline = append(pipeline, match)
	}

	err := r.db.DB(databaseName).C("Component").Pipe(pipeline).All(&incidents)
	if err != nil {
		log.Panicf("Error searching components: %s\n", err)
	}
	return incidents, err
}
