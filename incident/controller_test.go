package incident

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/involvestecnologia/statuspage/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const routerGroupName = "/test"

var router = gin.Default()
var incidentMockDAO = newMockIncidentDAO()

func init() {
	IncidentRouter(incidentMockDAO, router.Group(routerGroupName))
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
	body := []byte(`{"status": 1,"description": "test", "occurrence_date": "` + time.Now().Format(time.RFC3339) + `"}`)
	resp := performRequest(t, router, "POST", routerGroupName+"/incident/last", body)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var data string
	err := json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	//Invalid: incident body missing required parameter status
	body = []byte(`{"description": "test", "occurrence_date": "` + time.Now().Format(time.RFC3339) + `"}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/incident/last", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	//Invalid: incident missing required parameter occurrence_date
	body = []byte(`{"status":0, "description": "test"}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/incident/last", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	//Invalid: unknow component
	body = []byte(`{"status": 1,"description": "test", "occurrence_date": "` + time.Now().Format(time.RFC3339) + `"}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/incident/invalid_component_ref", body)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestController_Find(t *testing.T) {
	//Valid: incident with exitent ref
	resp := performRequest(t, router, "GET", routerGroupName+"/incident/last", nil)
	assert.Equal(t, http.StatusOK, resp.Code)

	var data []models.Incident
	err := json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	//Invalid: inexistent component name
	resp = performRequest(t, router, "GET", routerGroupName+"/incident/test2", nil)

	assert.Equal(t, http.StatusNotFound, resp.Code)

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
}
