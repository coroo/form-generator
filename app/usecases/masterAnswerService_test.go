package usecases

import (
	"net/http/httptest"
	"testing"
	"time"

	entity "github.com/coroo/form-generator/app/entity"
	repositories "github.com/coroo/form-generator/app/repositories"
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
		FormAnswer:       "Dummy First Name 1",
		FormScore:        "1",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}, entity.MasterAnswer{
		ID:               2,
		MasterQuestionId: 1,
		FormAnswer:       "Dummy Last Name 2",
		FormScore:        "2",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	},
}

type repoMockMasterAnswer struct {
	mock.Mock
}

func (r *repoMockMasterAnswer) SaveMasterAnswer(masterAnswer entity.MasterAnswer) error {
	args := r.Called(masterAnswer)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockMasterAnswer) UpdateMasterAnswer(masterAnswer entity.MasterAnswer) error {
	args := r.Called(masterAnswer)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockMasterAnswer) DeleteMasterAnswer(masterAnswer entity.MasterAnswer) error {
	args := r.Called(masterAnswer)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockMasterAnswer) GetAllMasterAnswers() []entity.MasterAnswer {
	return dummyMasterAnswer
}

func (r *repoMockMasterAnswer) GetMasterAnswer(ctx *gin.Context) []entity.MasterAnswer {
	return dummyMasterAnswer
}

func (r *repoMockMasterAnswer) GetMasterAnswerByQuestion(ctx *gin.Context) []entity.MasterAnswer {
	return dummyMasterAnswer
}

func (r *repoMockMasterAnswer) CloseDB() {
}

type MasterAnswerDeliveryTestSuite struct {
	suite.Suite
	repositoryTest repositories.MasterAnswerRepository
}

func (suite *MasterAnswerDeliveryTestSuite) SetupTest() {
	suite.repositoryTest = new(repoMockMasterAnswer)
}

func (suite *MasterAnswerDeliveryTestSuite) TestBuildMasterAnswerService() {
	resultTest := NewMasterAnswer(suite.repositoryTest)
	var dummyImpl *MasterAnswerService
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
	// assert.NotNil(suite.T(), resultTest.(*MasterAnswerService).repositories)
}

func (suite *MasterAnswerDeliveryTestSuite) TestSaveMasterAnswerDelivery() {
	suite.repositoryTest.(*repoMockMasterAnswer).On("SaveMasterAnswer", dummyMasterAnswer[0]).Return(nil)
	useCaseTest := NewMasterAnswer(suite.repositoryTest)
	err := useCaseTest.SaveMasterAnswer(dummyMasterAnswer[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterAnswerDeliveryTestSuite) TestUpdateMasterAnswerDelivery() {
	suite.repositoryTest.(*repoMockMasterAnswer).On("UpdateMasterAnswer", dummyMasterAnswer[0]).Return(nil)
	useCaseTest := NewMasterAnswer(suite.repositoryTest)
	err := useCaseTest.UpdateMasterAnswer(dummyMasterAnswer[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterAnswerDeliveryTestSuite) TestDeleteMasterAnswerDelivery() {
	suite.repositoryTest.(*repoMockMasterAnswer).On("DeleteMasterAnswer", dummyMasterAnswer[0]).Return(nil)
	useCaseTest := NewMasterAnswer(suite.repositoryTest)
	err := useCaseTest.DeleteMasterAnswer(dummyMasterAnswer[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterAnswerDeliveryTestSuite) TestGetAllMasterAnswers() {
	suite.repositoryTest.(*repoMockMasterAnswer).On("GetAllMasterAnswers", dummyMasterAnswer).Return(dummyMasterAnswer)
	useCaseTest := NewMasterAnswer(suite.repositoryTest)
	dummyUser := useCaseTest.GetAllMasterAnswers()
	assert.Equal(suite.T(), dummyMasterAnswer, dummyUser)
}

func (suite *MasterAnswerDeliveryTestSuite) TestGetMasterAnswer() {
	suite.repositoryTest.(*repoMockMasterAnswer).On("GetMasterAnswer", dummyMasterAnswer[0].ID).Return(dummyMasterAnswer[0], nil)
	useCaseTest := NewMasterAnswer(suite.repositoryTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyUser := useCaseTest.GetMasterAnswer(c)
	assert.NotNil(suite.T(), dummyUser[0])
	assert.Equal(suite.T(), dummyMasterAnswer[0], dummyUser[0])
}

func (suite *MasterAnswerDeliveryTestSuite) TestGetMasterAnswerByQuestion() {
	suite.repositoryTest.(*repoMockMasterAnswer).On("GetMasterAnswerByQuestion", dummyMasterAnswer[0].ID).Return(dummyMasterAnswer[0], nil)
	useCaseTest := NewMasterAnswer(suite.repositoryTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyUser := useCaseTest.GetMasterAnswerByQuestion(c)
	assert.NotNil(suite.T(), dummyUser[0])
	assert.Equal(suite.T(), dummyMasterAnswer[0], dummyUser[0])
}

func TestMasterAnswerDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(MasterAnswerDeliveryTestSuite))
}
