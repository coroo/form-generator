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
var dummyMasterAnswer = []entity.MasterAnswer{
	entity.MasterAnswer{
		ID:               1,
		MasterQuestionId: 1,
		FormAnswer:       "Pernah",
		FormScore:        "70",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}, entity.MasterAnswer{
		ID:               2,
		MasterQuestionId: 1,
		FormAnswer:       "Tidak Pernah",
		FormScore:        "30",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	},
}

type serviceMockMasterAnswer struct {
	mock.Mock
}

func (r *serviceMockMasterAnswer) SaveMasterAnswer(masterAnswer entity.MasterAnswer) error {
	args := r.Called(masterAnswer)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *serviceMockMasterAnswer) UpdateMasterAnswer(masterAnswer entity.MasterAnswer) error {
	args := r.Called(masterAnswer)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *serviceMockMasterAnswer) DeleteMasterAnswer(masterAnswer entity.MasterAnswer) error {
	args := r.Called(masterAnswer)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *serviceMockMasterAnswer) GetAllMasterAnswers() []entity.MasterAnswer {
	return dummyMasterAnswer
}

func (r *serviceMockMasterAnswer) GetMasterAnswer(ctx *gin.Context) []entity.MasterAnswer {
	return dummyMasterAnswer
}

func (r *serviceMockMasterAnswer) GetMasterAnswerByQuestion(ctx *gin.Context) []entity.MasterAnswer {
	return dummyMasterAnswer
}

type MasterAnswerDeliveryTestSuite struct {
	suite.Suite
	serviceTest usecases.MasterAnswerService
}

func (suite *MasterAnswerDeliveryTestSuite) SetupTest() {
	suite.serviceTest = new(serviceMockMasterAnswer)
}

func (suite *MasterAnswerDeliveryTestSuite) TestBuildMasterAnswerController() {
	resultTest := NewMasterAnswer(suite.serviceTest)
	var dummyImpl *MasterAnswerController
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
}

func (suite *MasterAnswerDeliveryTestSuite) TestSaveMasterAnswerDelivery() {
	suite.serviceTest.(*serviceMockMasterAnswer).On("SaveMasterAnswer", dummyMasterAnswer[0]).Return(nil)
	deliveryTest := NewMasterAnswer(suite.serviceTest)
	err := deliveryTest.Save(dummyMasterAnswer[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterAnswerDeliveryTestSuite) TestUpdateMasterAnswerDelivery() {
	suite.serviceTest.(*serviceMockMasterAnswer).On("UpdateMasterAnswer", dummyMasterAnswer[0]).Return(nil)
	deliveryTest := NewMasterAnswer(suite.serviceTest)
	err := deliveryTest.Update(dummyMasterAnswer[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterAnswerDeliveryTestSuite) TestDeleteMasterAnswerDelivery() {
	suite.serviceTest.(*serviceMockMasterAnswer).On("DeleteMasterAnswer", dummyMasterAnswer[0]).Return(nil)
	deliveryTest := NewMasterAnswer(suite.serviceTest)
	err := deliveryTest.Delete(dummyMasterAnswer[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterAnswerDeliveryTestSuite) TestGetAllMasterAnswers() {
	suite.serviceTest.(*serviceMockMasterAnswer).On("GetAllMasterAnswers", dummyMasterAnswer).Return(dummyMasterAnswer)
	deliveryTest := NewMasterAnswer(suite.serviceTest)
	dummyUser := deliveryTest.GetAllMasterAnswers()
	assert.Equal(suite.T(), dummyMasterAnswer, dummyUser)
}

func (suite *MasterAnswerDeliveryTestSuite) TestGetMasterAnswer() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	suite.serviceTest.(*serviceMockMasterAnswer).On("GetMasterAnswer", dummyMasterAnswer[0].ID).Return(dummyMasterAnswer[0], nil)
	deliveryTest := NewMasterAnswer(suite.serviceTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyUser := deliveryTest.GetMasterAnswer(c)
	assert.NotNil(suite.T(), dummyUser[0])
	assert.Equal(suite.T(), dummyMasterAnswer[0], dummyUser[0])
}

func (suite *MasterAnswerDeliveryTestSuite) TestGetMasterAnswerByQuestion() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	suite.serviceTest.(*serviceMockMasterAnswer).On("GetMasterAnswerByQuestion", dummyMasterAnswer[0].ID).Return(dummyMasterAnswer[0], nil)
	deliveryTest := NewMasterAnswer(suite.serviceTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyUser := deliveryTest.GetMasterAnswerByQuestion(c)
	assert.NotNil(suite.T(), dummyUser[0])
	assert.Equal(suite.T(), dummyMasterAnswer[0], dummyUser[0])
}

func TestMasterAnswerDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(MasterAnswerDeliveryTestSuite))
}
