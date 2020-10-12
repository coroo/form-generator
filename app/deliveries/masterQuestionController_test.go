package deliveries

import (
	"testing"
	"time"

	entity "github.com/coroo/form-generator/app/entity"
	usecases "github.com/coroo/form-generator/app/usecases"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// dummy data
var dummyMasterAnswer = []*entity.MasterQuestion{
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

type serviceMock struct {
	mock.Mock
}

func (s *serviceMock) SaveFormGenerator(masterQuestion entity.MasterQuestion) error {
	return nil
}

func (s *serviceMock) UpdateFormGenerator(masterQuestion entity.MasterQuestion) error {
	return nil
}

func (s *serviceMock) DeleteFormGenerator(masterQuestion entity.MasterQuestion) error {
	return nil
}

func (s *serviceMock) GetAllMasterQuestions() []entity.MasterQuestion {
	return nil
}

func (s *serviceMock) GetMasterQuestion(ctx *gin.Context) []entity.MasterQuestion {
	return nil
}

type MasterQuestionDeliveryTestSuite struct {
	suite.Suite
	serviceTest usecases.MasterQuestionService
}

// func (suite *MasterQuestionDeliveryTestSuite) SetupTest() {
// 	suite.serviceTest = new(serviceMock)
// }

// func (suite *MasterQuestionDeliveryTestSuite) TestBuildMasterQuestionController() {
// 	resultTest := NewMasterQuestion(suite.serviceTest)
// 	var dummyImpl *MasterQuestionController
// 	assert.NotNil(suite.T(), resultTest)
// 	assert.Implements(suite.T(), dummyImpl, resultTest)
// 	// assert.NotNil(suite.T(), resultTest.(*MasterQuestionController).usecases)
// }

// func (suite *MasterQuestionDeliveryTestSuite) TestSaveDelivery() {
// 	// Switch to test mode so you don't get such noisy output
// 	gin.SetMode(gin.TestMode)

// 	r := gin.Default()
// 	w := httptest.NewRecorder()

// 	resultTest := NewMasterQuestion(suite.serviceTest)
// 	handler := resultTest.(MasterQuestionController).Save
// 	r.POST("masterAnswer/create", handler)
// 	// r.POST("masterAnswer/create", NewMasterQuestion(suite.serviceTest).Save)

// 	jsonValue, _ := json.Marshal(dummyMasterAnswer[0])
// 	req, _ := http.NewRequest(http.MethodPost, "/masterAnswer/create", bytes.NewBuffer(jsonValue))
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	r.ServeHTTP(w, req)
// 	assert.Equal(suite.T(), w.Code, 200)

// 	// r.POST("/syUserInvoice/map-etl-payment", routes.SyMapEtlLatestPayment)

// 	// resultTest := NewMasterQuestion(suite.serviceTest)
// 	// rr := httptest.NewRecorder()
// 	// reqBody, _ := json.Marshal(dummyMasterAnswer[0])

// 	// req, _ := http.NewRequest("POST", "api/v1/masterQuestion/create", nil)
// 	// suite.routerTest.ServeHTTP(rr, req)
// 	// assert.Equal(suite.T(), rr.Code, 200)

// 	// handler := resultTest.(*MasterQuestionController).save
// 	// suite.routerTest.HandleFunc("/dummy-user", handler)
// 	// req, _ := http.NewRequest("POST", "/dummy-user", bytes.NewBuffer(reqBody))
// 	// req.Header.Set("Content-Type", "application/json")

// 	// suite.routerTest.ServeHTTP(rr, req)
// 	// respTest := rr.Result()
// 	// respBody := new(appHttpResponse.Response)

// 	// if err := json.NewDecoder(respTest.Body).Decode(respBody); err != nil {
// 	// }
// 	// assert.Equal(suite.T(), rr.Code, 200)
// 	// assert.Equal(suite.T(), respBody.Data.(map[string]interface{})["firstName"], "Dummy First Name 1")
// }

func TestMasterQuestionDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(MasterQuestionDeliveryTestSuite))
}
