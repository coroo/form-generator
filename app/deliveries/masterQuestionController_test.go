package deliveries

import (
	"net/http/httptest"
	"testing"
	"time"

	entity "github.com/coroo/form-generator/app/entity"
	usecases "github.com/coroo/form-generator/app/usecases"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// dummy data
var dummyMasterQuestion = []entity.MasterQuestion{
	entity.MasterQuestion{
		ID:           1,
		QuestionText: "Dummy First Name 1",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, entity.MasterQuestion{
		ID:           2,
		QuestionText: "Dummy Last Name 2",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	},
}

type serviceMockMasterQuestion struct {
	mock.Mock
}

func (r *serviceMockMasterQuestion) SaveMasterQuestion(masterQuestion entity.MasterQuestion) error {
	args := r.Called(masterQuestion)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *serviceMockMasterQuestion) UpdateMasterQuestion(masterQuestion entity.MasterQuestion) error {
	args := r.Called(masterQuestion)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *serviceMockMasterQuestion) DeleteMasterQuestion(masterQuestion entity.MasterQuestion) error {
	args := r.Called(masterQuestion)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *serviceMockMasterQuestion) GetAllMasterQuestions() []entity.MasterQuestion {
	return dummyMasterQuestion
}

func (r *serviceMockMasterQuestion) GetMasterQuestion(ctx *gin.Context) []entity.MasterQuestion {
	return dummyMasterQuestion
}

type MasterQuestionDeliveryTestSuite struct {
	suite.Suite
	serviceTest usecases.MasterQuestionService
}

func (suite *MasterQuestionDeliveryTestSuite) SetupTest() {
	suite.serviceTest = new(serviceMockMasterQuestion)
}

func (suite *MasterQuestionDeliveryTestSuite) TestBuildMasterQuestionController() {
	resultTest := NewMasterQuestion(suite.serviceTest)
	var dummyImpl *MasterQuestionController
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
}

func (suite *MasterQuestionDeliveryTestSuite) TestSaveMasterQuestionDelivery() {
	suite.serviceTest.(*serviceMockMasterQuestion).On("SaveMasterQuestion", dummyMasterQuestion[0]).Return(nil)
	deliveryTest := NewMasterQuestion(suite.serviceTest)
	err := deliveryTest.Save(dummyMasterQuestion[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterQuestionDeliveryTestSuite) TestUpdateMasterQuestionDelivery() {
	suite.serviceTest.(*serviceMockMasterQuestion).On("UpdateMasterQuestion", dummyMasterQuestion[0]).Return(nil)
	deliveryTest := NewMasterQuestion(suite.serviceTest)
	err := deliveryTest.Update(dummyMasterQuestion[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterQuestionDeliveryTestSuite) TestDeleteMasterQuestionDelivery() {
	suite.serviceTest.(*serviceMockMasterQuestion).On("DeleteMasterQuestion", dummyMasterQuestion[0]).Return(nil)
	deliveryTest := NewMasterQuestion(suite.serviceTest)
	err := deliveryTest.Delete(dummyMasterQuestion[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterQuestionDeliveryTestSuite) TestGetAllMasterQuestions() {
	suite.serviceTest.(*serviceMockMasterQuestion).On("GetAllMasterQuestions", dummyMasterQuestion).Return(dummyMasterQuestion)
	deliveryTest := NewMasterQuestion(suite.serviceTest)
	dummyUser := deliveryTest.GetAllMasterQuestions()
	assert.Equal(suite.T(), dummyMasterQuestion, dummyUser)
}

func (suite *MasterQuestionDeliveryTestSuite) TestGetMasterQuestion() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	suite.serviceTest.(*serviceMockMasterQuestion).On("GetMasterQuestion", dummyMasterQuestion[0].ID).Return(dummyMasterQuestion[0], nil)
	deliveryTest := NewMasterQuestion(suite.serviceTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyUser := deliveryTest.GetMasterQuestion(c)
	assert.NotNil(suite.T(), dummyUser[0])
	assert.Equal(suite.T(), dummyMasterQuestion[0], dummyUser[0])
}

func TestMasterQuestionDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(MasterQuestionDeliveryTestSuite))
}
