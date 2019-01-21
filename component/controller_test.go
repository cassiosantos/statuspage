package component_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/mock"
	"github.com/involvestecnologia/statuspage/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const routerGroupName = "/test"

var router = gin.Default()
var componentMockDAO = mock.NewMockComponentDAO()

func init() {
	component.ComponentRouter(componentMockDAO, router.Group(routerGroupName))
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
	//Valid: new component body
	body := []byte(`{"name": "test component","address": "t.e.s.t", "incidents_history": []}`)
	resp := performRequest(t, router, "POST", routerGroupName+"/component", body)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var data string
	err := json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	//Invalid: component body with exitent name
	body = []byte(`{"name": "first","incidents_history": []}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/component", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	//Invalid: component missing required parameter name
	body = []byte(`{"incidents_history": []}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/component", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	//Invalid: component body with exitent ref
	body = []byte(`{"ref":"886e09000000000000000000","name": "test2","incidents_history": []}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/component", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestController_Update(t *testing.T) {
	//Valid: component with exitent ref
	body := []byte(`{"name": "test1","resources": []}`)
	resp := performRequest(t, router, "PATCH", routerGroupName+"/component/886e09000000000000000000", body)

	assert.Equal(t, http.StatusOK, resp.Code)

	//Invalid: inexistent component ref
	body = []byte(`{"name": "test2","resources": []}`)
	resp = performRequest(t, router, "PATCH", routerGroupName+"/component/test2", body)

	assert.Equal(t, http.StatusNotFound, resp.Code)

	//Invalid: missing name parameter
	body = []byte(`{"resources": []}`)
	resp = performRequest(t, router, "PATCH", routerGroupName+"/component/886e09000000000000000000", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestController_Find(t *testing.T) {
	//Valid: component with exitent ref
	resp := performRequest(t, router, "GET", routerGroupName+"/component/886e09000000000000000000", nil)
	assert.Equal(t, http.StatusOK, resp.Code)

	var data models.Component
	err := json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	//Invalid: inexistent component ref
	resp = performRequest(t, router, "GET", routerGroupName+"/component/test2", nil)

	assert.Equal(t, http.StatusNotFound, resp.Code)

	//Valid: finding component by name
	resp = performRequest(t, router, "GET", routerGroupName+"/component/test?search=name", nil)
	assert.Equal(t, http.StatusOK, resp.Code)

	err = json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)
}

func TestController_Delete(t *testing.T) {
	resp := performRequest(t, router, "DELETE", routerGroupName+"/component/886e09000000000000000000", nil)

	assert.Equal(t, http.StatusNoContent, resp.Code)

	resp = performRequest(t, router, "DELETE", routerGroupName+"/component/invalidRef", nil)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestController_List(t *testing.T) {
	resp := performRequest(t, router, "GET", routerGroupName+"/components", nil)

	assert.Equal(t, http.StatusOK, resp.Code)

	var data []models.Component
	err := json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)
}
