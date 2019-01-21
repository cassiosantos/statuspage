package incident

import (
	"log"
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

func (r *MongoRepository) Insert(incident models.Incident) error {
	err := r.db.DB(databaseName).C("Incidents").Insert(incident)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error inserting incident: %s\n", err)
	}
	return err
}

func (r *MongoRepository) Find(queryParam map[string]interface{}) ([]models.Incident, error) {
	var incidents []models.Incident

	err := r.db.DB(databaseName).C("Incidents").Find(queryParam).All(&incidents)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error finding component: %s\n", err)
	}
	if incidents == nil {
		return incidents, errors.E(errors.ErrNotFound)
	}
	return incidents, err
}

func (r *MongoRepository) List(startDt time.Time, endDt time.Time) ([]models.Incident, error) {
	var incidents []models.Incident

	findQ := bson.M{
		"date": bson.M{
			"$gt": startDt.Add(-(24 * time.Hour)),
			"$lt": endDt.Add(24 * time.Hour),
		},
	}

	err := r.db.DB(databaseName).C("Incidents").Find(findQ).All(&incidents)
	if err != nil {
		log.Panicf("Error searching components: %s\n", err)
	}
	return incidents, err
}
