package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	deliveries "github.com/coroo/form-generator/app/deliveries"
	entity "github.com/coroo/form-generator/app/entity"

	"github.com/gin-gonic/gin"
)

// dummy data
var dummyFormGenerator = []*entity.FormGenerator{
	&entity.FormGenerator{
		ID:               1,
		MasterQuestionId: 1,
		MasterFormId:     1,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}, &entity.FormGenerator{
		ID:               2,
		MasterQuestionId: 1,
		MasterFormId:     1,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	},
}

type formGeneratorRouteMock struct {
	mock.Mock
}

func (s *formGeneratorRouteMock) Save(formGenerator entity.FormGenerator) error {
	return nil
}

func (s *formGeneratorRouteMock) Update(formGenerator entity.FormGenerator) error {
	return nil
}

func (s *formGeneratorRouteMock) Delete(formGenerator entity.FormGenerator) error {
	return nil
}

func (s *formGeneratorRouteMock) GetAllFormGenerators() []entity.FormGenerator {
	return nil
}

func (s *formGeneratorRouteMock) GetFormGenerator(ctx *gin.Context) []entity.FormGenerator {
	return nil
}

func (s *formGeneratorRouteMock) GetFormGeneratorByQuestion(ctx *gin.Context) []entity.FormGenerator {
	return nil
}

type FormGeneratorRouteTestSuite struct {
	suite.Suite
	serviceTest deliveries.FormGeneratorController
}

func (suite *FormGeneratorRouteTestSuite) SetupTest() {
	suite.serviceTest = new(formGeneratorRouteMock)
}

func (suite *FormGeneratorRouteTestSuite) TestSaveDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("formGenerator/create", FormGeneratorCreate)

	jsonValue, _ := json.Marshal(dummyFormGenerator[0])
	req, _ := http.NewRequest(http.MethodPost, "/formGenerator/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *FormGeneratorRouteTestSuite) TestUpdateDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("formGenerator/update", FormGeneratorUpdate)

	jsonValue, _ := json.Marshal(dummyFormGenerator[0])
	req, _ := http.NewRequest(http.MethodPost, "/formGenerator/update", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *FormGeneratorRouteTestSuite) TestDeleteDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("formGenerator/delete", FormGeneratorDelete)

	jsonValue, _ := json.Marshal(dummyFormGenerator[0])
	req, _ := http.NewRequest(http.MethodPost, "/formGenerator/delete", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *FormGeneratorRouteTestSuite) TestFormGeneratorsIndexRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("formGenerator/index", FormGeneratorsIndex)
	req, _ := http.NewRequest(http.MethodGet, "/formGenerator/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *FormGeneratorRouteTestSuite) TestFormGeneratorsDetailRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("formGenerator/detail/1", FormGeneratorsDetail)
	req, _ := http.NewRequest(http.MethodGet, "/formGenerator/detail/1", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func TestFormGeneratorRouteTestSuite(t *testing.T) {
	suite.Run(t, new(FormGeneratorRouteTestSuite))
}
