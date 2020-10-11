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
var dummyMasterForm = []*entity.MasterForm{
	&entity.MasterForm{
		ID:        1,
		FormName:  "Dummy First Name 1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, &entity.MasterForm{
		ID:        2,
		FormName:  "Dummy Last Name 2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

type masterFormRouteMock struct {
	mock.Mock
}

func (s *masterFormRouteMock) Save(ctx *gin.Context) error {
	return nil
}

func (s *masterFormRouteMock) Update(ctx *gin.Context) error {
	return nil
}

func (s *masterFormRouteMock) Delete(ctx *gin.Context) error {
	return nil
}

func (s *masterFormRouteMock) GetAllMasterForms() []entity.MasterForm {
	return nil
}

func (s *masterFormRouteMock) GetMasterForm(ctx *gin.Context) []entity.MasterForm {
	return nil
}

type MasterFormRouteTestSuite struct {
	suite.Suite
	serviceTest deliveries.MasterFormController
}

func (suite *MasterFormRouteTestSuite) SetupTest() {
	suite.serviceTest = new(masterFormRouteMock)
}

func (suite *MasterFormRouteTestSuite) TestSaveDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("masterForm/create", MasterFormCreate)

	jsonValue, _ := json.Marshal(dummyMasterForm[0])
	req, _ := http.NewRequest(http.MethodPost, "/masterForm/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *MasterFormRouteTestSuite) TestUpdateDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("masterForm/update", MasterFormUpdate)

	jsonValue, _ := json.Marshal(dummyMasterForm[0])
	req, _ := http.NewRequest(http.MethodPost, "/masterForm/update", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *MasterFormRouteTestSuite) TestDeleteDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("masterForm/delete", MasterFormDelete)

	jsonValue, _ := json.Marshal(dummyMasterForm[0])
	req, _ := http.NewRequest(http.MethodPost, "/masterForm/delete", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *MasterFormRouteTestSuite) TestMasterFormsIndexRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("masterForm/index", MasterFormsIndex)
	req, _ := http.NewRequest(http.MethodGet, "/masterForm/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *MasterFormRouteTestSuite) TestMasterFormsDetailRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("masterForm/detail/1", MasterFormsDetail)
	req, _ := http.NewRequest(http.MethodGet, "/masterForm/detail/1", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func TestMasterFormRouteTestSuite(t *testing.T) {
	suite.Run(t, new(MasterFormRouteTestSuite))
}
