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
var dummyMasterAnswer = []*entity.MasterAnswer{
	&entity.MasterAnswer{
		ID:               1,
		MasterQuestionId: 1,
		FormAnswer:       "Pernah",
		FormScore:        "70",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}, &entity.MasterAnswer{
		ID:               2,
		MasterQuestionId: 1,
		FormAnswer:       "Tidak Pernah",
		FormScore:        "30",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	},
}

type NegativeMasterAnswer struct {
	NoColumn int `json:"no_column"`
}

var dummyNegativeMasterAnswer = []*NegativeMasterAnswer{
	&NegativeMasterAnswer{
		NoColumn: 1,
	},
}

type masterAnswerRouteMock struct {
	mock.Mock
}

func (s *masterAnswerRouteMock) Save(masterAnswer entity.MasterAnswer) error {
	return nil
}

func (s *masterAnswerRouteMock) Update(masterAnswer entity.MasterAnswer) error {
	return nil
}

func (s *masterAnswerRouteMock) Delete(masterAnswer entity.MasterAnswer) error {
	return nil
}

func (s *masterAnswerRouteMock) GetAllMasterAnswers() []entity.MasterAnswer {
	return nil
}

func (s *masterAnswerRouteMock) GetMasterAnswer(ctx *gin.Context) []entity.MasterAnswer {
	return nil
}

func (s *masterAnswerRouteMock) GetMasterAnswerByQuestion(ctx *gin.Context) []entity.MasterAnswer {
	return nil
}

type MasterAnswerRouteTestSuite struct {
	suite.Suite
	serviceTest deliveries.MasterAnswerController
}

func (suite *MasterAnswerRouteTestSuite) SetupTest() {
	suite.serviceTest = new(masterAnswerRouteMock)
}

func (suite *MasterAnswerRouteTestSuite) TestSaveDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	r := gin.Default()

	r.POST("masterAnswer/create", MasterAnswerCreate)

	jsonValue, _ := json.Marshal(dummyMasterAnswer[0])
	req, _ := http.NewRequest(http.MethodPost, "/masterAnswer/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *MasterAnswerRouteTestSuite) TestFailSaveDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	r := gin.Default()

	r.POST("masterAnswer/create", MasterAnswerCreate)

	jsonValue, _ := json.Marshal(dummyNegativeMasterAnswer[0])
	req, _ := http.NewRequest(http.MethodPost, "/masterAnswer/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *MasterAnswerRouteTestSuite) TestUpdateDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("masterAnswer/update", MasterAnswerUpdate)

	jsonValue, _ := json.Marshal(dummyMasterAnswer[0])
	req, _ := http.NewRequest(http.MethodPost, "/masterAnswer/update", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *MasterAnswerRouteTestSuite) TestDeleteDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("masterAnswer/delete", MasterAnswerDelete)

	jsonValue, _ := json.Marshal(dummyMasterAnswer[0])
	req, _ := http.NewRequest(http.MethodPost, "/masterAnswer/delete", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *MasterAnswerRouteTestSuite) TestMasterAnswersIndexRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("masterAnswer/index", MasterAnswersIndex)
	req, _ := http.NewRequest(http.MethodGet, "/masterAnswer/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *MasterAnswerRouteTestSuite) TestMasterAnswersDetailRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("masterAnswer/detail/1", MasterAnswersDetail)
	req, _ := http.NewRequest(http.MethodGet, "/masterAnswer/detail/1", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *MasterAnswerRouteTestSuite) TestMasterAnswersByQuestionRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("masterAnswer/byQuestion/1", MasterAnswersByQuestion)
	req, _ := http.NewRequest(http.MethodGet, "/masterAnswer/byQuestion/1", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func TestMasterAnswerRouteTestSuite(t *testing.T) {
	suite.Run(t, new(MasterAnswerRouteTestSuite))
}
