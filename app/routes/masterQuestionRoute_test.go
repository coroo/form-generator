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
var dummyMasterQuestion = []*entity.MasterQuestion{
	&entity.MasterQuestion{
		ID:           1,
		QuestionText: "Dummy First Name 1",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, &entity.MasterQuestion{
		ID:           2,
		QuestionText: "Dummy Last Name 2",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	},
}

type masterQuestionRouteMock struct {
	mock.Mock
}

func (s *masterQuestionRouteMock) Save(masterQuestion entity.MasterQuestion) error {
	return nil
}

func (s *masterQuestionRouteMock) Update(masterQuestion entity.MasterQuestion) error {
	return nil
}

func (s *masterQuestionRouteMock) Delete(masterQuestion entity.MasterQuestion) error {
	return nil
}

func (s *masterQuestionRouteMock) GetAllMasterQuestions() []entity.MasterQuestion {
	return nil
}

func (s *masterQuestionRouteMock) GetMasterQuestion(ctx *gin.Context) []entity.MasterQuestion {
	return nil
}

type MasterQuestionRouteTestSuite struct {
	suite.Suite
	serviceTest deliveries.MasterQuestionController
}

func (suite *MasterQuestionRouteTestSuite) SetupTest() {
	suite.serviceTest = new(masterQuestionRouteMock)
}

func (suite *MasterQuestionRouteTestSuite) TestSaveDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("masterQuestion/create", MasterQuestionCreate)

	jsonValue, _ := json.Marshal(dummyMasterQuestion[0])
	req, _ := http.NewRequest(http.MethodPost, "/masterQuestion/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *MasterQuestionRouteTestSuite) TestUpdateDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("masterQuestion/update", MasterQuestionUpdate)

	jsonValue, _ := json.Marshal(dummyMasterQuestion[0])
	req, _ := http.NewRequest(http.MethodPost, "/masterQuestion/update", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *MasterQuestionRouteTestSuite) TestDeleteDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("masterQuestion/delete", MasterQuestionDelete)

	jsonValue, _ := json.Marshal(dummyMasterQuestion[0])
	req, _ := http.NewRequest(http.MethodPost, "/masterQuestion/delete", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *MasterQuestionRouteTestSuite) TestMasterQuestionsIndexRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("masterQuestion/index", MasterQuestionsIndex)
	req, _ := http.NewRequest(http.MethodGet, "/masterQuestion/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *MasterQuestionRouteTestSuite) TestMasterQuestionsDetailRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("masterQuestion/detail/1", MasterQuestionsDetail)
	req, _ := http.NewRequest(http.MethodGet, "/masterQuestion/detail/1", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func TestMasterQuestionRouteTestSuite(t *testing.T) {
	suite.Run(t, new(MasterQuestionRouteTestSuite))
}
