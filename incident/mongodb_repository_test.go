package incident

import (
	"fmt"
	"log"
	"testing"
	"time"

	mgo "github.com/globalsign/mgo"
	"github.com/involvestecnologia/statuspage/db"
	"github.com/involvestecnologia/statuspage/models"
	"github.com/stretchr/testify/assert"
)

const validMongoArgs = "localhost"

var testSession *mgo.Session
var i = models.Incident{
	Status:      models.IncidentStatusOutage,
	Description: "",
	Date:        time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
}
var c = models.Component{
	Ref:     zeroTimeHex,
	Name:    "Component",
	Address: "",
	Incidents: []models.Incident{
		models.Incident{
			Status:      models.IncidentStatusOK,
			Description: "",
			Date:        time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	},
}

func init() {
	testSession = db.InitMongo(validMongoArgs)
	err := testSession.DB("status").DropDatabase()
	if err != nil {
		log.Panicf("%s\n", err)
	}
	testSession.DB("status").C("Component").Insert(c)
	fmt.Println("First insert done.")
}

func TestIncidentMongoDB_Repository_NewMongoRepository(t *testing.T) {
	var mongoRepo *MongoRepository
	repo := NewMongoRepository(testSession)
	assert.IsType(t, mongoRepo, repo)
	assert.Equal(t, testSession, repo.db)
}

func TestIncidentMongoDB_Repository_Insert(t *testing.T) {

	repo := NewMongoRepository(testSession)

	err := repo.Insert(c.Name, i)
	assert.Nil(t, err)

	c.Incidents = append(c.Incidents, i)

	incidents, err := repo.Find(c.Name)
	if assert.Nil(t, err) && assert.NotNil(t, incidents) {
		assert.Equal(t, c.Incidents, incidents)
	}

	err = repo.Insert("Invalid Ref", i)
	assert.NotNil(t, err)
}

func TestIncidentMongoDB_Repository_Find(t *testing.T) {
	repo := NewMongoRepository(testSession)

	incidents, err := repo.Find(c.Name)
	if assert.Nil(t, err) && assert.NotNil(t, incidents) {
		assert.Equal(t, c.Incidents, incidents)
	}

	_, err = repo.Find("Invalid Name")
	assert.NotNil(t, err)
}

func TestIncidentMongoDB_Repository_List(t *testing.T) {
	repo := NewMongoRepository(testSession)

	startDt := time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDt := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)

	incCompName := make([]models.IncidentWithComponentName, 0)
	for _, v := range c.Incidents {
		incCompName = append(incCompName,
			models.IncidentWithComponentName{
				Component: c.Name,
				Incident:  v,
			})
	}

	incidents, err := repo.List(startDt, endDt)
	if assert.Nil(t, err) && assert.NotNil(t, incidents) {
		assert.Equal(t, incCompName, incidents)
	}

	firstIncCompName := []models.IncidentWithComponentName{
		models.IncidentWithComponentName{
			Component: c.Name,
			Incident:  c.Incidents[0],
		},
	}

	endDt = time.Date(2018, time.January, 2, 0, 0, 0, 0, time.UTC)
	incidents, err = repo.List(startDt, endDt)
	if assert.Nil(t, err) && assert.NotNil(t, incidents) {
		assert.Equal(t, firstIncCompName, incidents)
	}
}
