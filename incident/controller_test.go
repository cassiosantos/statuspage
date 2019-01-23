package incident_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/mock"

	"github.com/involvestecnologia/statuspage/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const routerGroupName = "/test"
const failureRouterGroupName = "/failure"

var router = gin.Default()
var incidentMockDAO = mock.NewMockIncidentDAO()
var componentMockDAO = mock.NewMockComponentDAO()
var incidentFailureMockDAO = mock.NewMockFailureIncidentDAO()

func init() {
	incident.IncidentRouter(incidentMockDAO, component.NewService(componentMockDAO), router.Group(routerGroupName))
	incident.IncidentRouter(incidentFailureMockDAO, component.NewService(componentMockDAO), router.Group(failureRouterGroupName))
}

func performRequest(t *testing.T, r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {

	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Failed to perform request: %s", err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestController_Create(t *testing.T) {
	//Valid: new incident body
	body := []byte(`{"component_ref":"` + mock.ZeroTimeHex + `","status": 1,"description": "test", "occurrence_date": "` + time.Now().Format(time.RFC3339) + `"}`)
	resp := performRequest(t, router, "POST", routerGroupName+"/incident", body)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var data string
	err := json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	//Invalid: incident body missing required parameter status
	body = []byte(`{"component_ref":"` + mock.ZeroTimeHex + `","description": "test", "occurrence_date": "` + time.Now().Format(time.RFC3339) + `"}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/incident", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	//Invalid: incident missing required parameter occurrence_date
	body = []byte(`{"component_ref":"` + mock.ZeroTimeHex + `","status":0, "description": "test"}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/incident", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	//Invalid: incident missing required parameter component_ref
	body = []byte(`{"status":0, "description": "test", "occurrence_date": "` + time.Now().Format(time.RFC3339) + `"}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/incident", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	//Invalid: unknow component
	body = []byte(`{"component_ref":"Invalid Component Ref","status": 1,"description": "test", "occurrence_date": "` + time.Now().Format(time.RFC3339) + `"}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/incident", body)

	assert.Equal(t, http.StatusNotFound, resp.Code)

	//Failure
	body = []byte(`{"component_ref":"` + mock.ZeroTimeHex + `","status": 1,"description": "test", "occurrence_date": "` + time.Now().Format(time.RFC3339) + `"}`)
	resp = performRequest(t, router, "POST", failureRouterGroupName+"/incident", body)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}
func TestController_Find(t *testing.T) {
	//Valid: incident with exitent ref
	resp := performRequest(t, router, "GET", routerGroupName+"/incident/"+mock.ZeroTimeHex, nil)
	assert.Equal(t, http.StatusOK, resp.Code)

	var data []models.Incident
	err := json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	//Invalid: inexistent component name
	resp = performRequest(t, router, "GET", routerGroupName+"/incident/test2", nil)

	assert.Equal(t, http.StatusNotFound, resp.Code)

	//Failure
	resp = performRequest(t, router, "GET", failureRouterGroupName+"/incident/"+mock.ZeroTimeHex, nil)
	assert.Equal(t, http.StatusInternalServerError, resp.Code)

}

func TestController_List(t *testing.T) {
	// Valid: no parameters
	resp := performRequest(t, router, "GET", routerGroupName+"/incidents", nil)

	assert.Equal(t, http.StatusOK, resp.Code)

	var data []models.Incident
	err := json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	// Valid: query parameters month and year
	resp = performRequest(t, router, "GET", routerGroupName+"/incidents?year=2019&month=1", nil)

	assert.Equal(t, http.StatusOK, resp.Code)

	err = json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	// Valid: query parameters only year
	resp = performRequest(t, router, "GET", routerGroupName+"/incidents?year=2019", nil)

	assert.Equal(t, http.StatusOK, resp.Code)

	err = json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	// Valid: query parameters only month
	resp = performRequest(t, router, "GET", routerGroupName+"/incidents?month=1", nil)

	assert.Equal(t, http.StatusOK, resp.Code)

	err = json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	// Invalid: query parameters year
	resp = performRequest(t, router, "GET", routerGroupName+"/incidents?year=0", nil)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Invalid: query parameters month
	resp = performRequest(t, router, "GET", routerGroupName+"/incidents?month=0", nil)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Failure
	resp = performRequest(t, router, "GET", failureRouterGroupName+"/incidents", nil)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}
