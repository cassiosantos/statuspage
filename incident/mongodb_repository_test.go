package incident_test

import (
	"log"
	"testing"
	"time"

	mgo "github.com/globalsign/mgo"
	"github.com/involvestecnologia/statuspage/db"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/mock"
	"github.com/involvestecnologia/statuspage/models"
	"github.com/stretchr/testify/assert"
)

const validMongoArgs = "localhost"

var testSession *mgo.Session
var failureSession *mgo.Session
var i = models.Incident{
	ComponentRef: mock.ZeroTimeHex,
	Status:       models.IncidentStatusOutage,
	Description:  "",
	Date:         time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
}
var c = models.Component{
	Ref:     mock.ZeroTimeHex,
	Name:    "Component",
	Address: "",
}

func init() {
	testSession = db.InitMongo(validMongoArgs)
	err := testSession.DB("status").DropDatabase()
	if err != nil {
		log.Panicf("%s\n", err)
	}
	testSession.DB("status").C("Component").Insert(c)
}

func TestIncidentMongoDB_Repository_NewMongoRepository(t *testing.T) {
	var mongoRepo *incident.MongoRepository
	repo := incident.NewMongoRepository(testSession)
	assert.IsType(t, mongoRepo, repo)
}

func TestIncidentMongoDB_Repository_Insert(t *testing.T) {

	repo := incident.NewMongoRepository(testSession)

	err := repo.Insert(i)
	assert.Nil(t, err)

	incidents, err := repo.Find(map[string]interface{}{"component_ref": i.ComponentRef})
	if assert.Nil(t, err) && assert.NotNil(t, incidents) {
		assert.Equal(t, []models.Incident{i}, incidents)
	}

	repo = incident.NewMongoRepository(failureSession)

	err = repo.Insert(i)
	assert.NotNil(t, err)

}

func TestIncidentMongoDB_Repository_Find(t *testing.T) {
	repo := incident.NewMongoRepository(testSession)

	incidents, err := repo.Find(map[string]interface{}{"component_ref": c.Ref})
	if assert.Nil(t, err) && assert.NotNil(t, incidents) {
		assert.Equal(t, []models.Incident{i}, incidents)
	}

	incidents, err = repo.Find(map[string]interface{}{"component_ref": "Invalid Ref"})
	assert.Nil(t, incidents)

	incidents, err = repo.Find(map[string]interface{}{"invalidQuery": "SomeValue"})
	assert.NotNil(t, err)
	assert.Nil(t, incidents)

	repo = incident.NewMongoRepository(failureSession)
	_, err = repo.Find(map[string]interface{}{"component_ref": c.Ref})
	assert.NotNil(t, err)

}

func TestIncidentMongoDB_Repository_List(t *testing.T) {
	repo := incident.NewMongoRepository(testSession)

	startDt := time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDt := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)

	incidents, err := repo.List(startDt, endDt)
	if assert.Nil(t, err) && assert.NotNil(t, incidents) {
		assert.Equal(t, []models.Incident{i}, incidents)
	}

	endDt = time.Date(2018, time.January, 2, 0, 0, 0, 0, time.UTC)
	incidents, err = repo.List(startDt, endDt)
	if assert.Nil(t, err) && assert.Nil(t, incidents) {
		assert.IsType(t, []models.Incident{}, incidents)
	}

	endDt = time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
	repo = incident.NewMongoRepository(failureSession)
	_, err = repo.List(startDt, endDt)
	assert.NotNil(t, err)
}
