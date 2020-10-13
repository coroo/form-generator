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

type repoMockMasterQuestion struct {
	mock.Mock
}

func (r *repoMockMasterQuestion) SaveMasterQuestion(masterQuestion entity.MasterQuestion) error {
	args := r.Called(masterQuestion)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockMasterQuestion) UpdateMasterQuestion(masterQuestion entity.MasterQuestion) error {
	args := r.Called(masterQuestion)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockMasterQuestion) DeleteMasterQuestion(masterQuestion entity.MasterQuestion) error {
	args := r.Called(masterQuestion)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockMasterQuestion) GetAllMasterQuestions() []entity.MasterQuestion {
	return dummyMasterQuestion
}

func (r *repoMockMasterQuestion) GetMasterQuestion(ctx *gin.Context) []entity.MasterQuestion {
	return dummyMasterQuestion
}

func (r *repoMockMasterQuestion) CloseDB() {
}

type MasterQuestionDeliveryTestSuite struct {
	suite.Suite
	repositoryTest repositories.MasterQuestionRepository
}

func (suite *MasterQuestionDeliveryTestSuite) SetupTest() {
	suite.repositoryTest = new(repoMockMasterQuestion)
}

func (suite *MasterQuestionDeliveryTestSuite) TestBuildMasterQuestionService() {
	resultTest := NewMasterQuestion(suite.repositoryTest)
	var dummyImpl *MasterQuestionService
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
	// assert.NotNil(suite.T(), resultTest.(*MasterQuestionService).repositories)
}

func (suite *MasterQuestionDeliveryTestSuite) TestSaveMasterQuestionDelivery() {
	suite.repositoryTest.(*repoMockMasterQuestion).On("SaveMasterQuestion", dummyMasterQuestion[0]).Return(nil)
	useCaseTest := NewMasterQuestion(suite.repositoryTest)
	err := useCaseTest.SaveMasterQuestion(dummyMasterQuestion[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterQuestionDeliveryTestSuite) TestUpdateMasterQuestionDelivery() {
	suite.repositoryTest.(*repoMockMasterQuestion).On("UpdateMasterQuestion", dummyMasterQuestion[0]).Return(nil)
	useCaseTest := NewMasterQuestion(suite.repositoryTest)
	err := useCaseTest.UpdateMasterQuestion(dummyMasterQuestion[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterQuestionDeliveryTestSuite) TestDeleteMasterQuestionDelivery() {
	suite.repositoryTest.(*repoMockMasterQuestion).On("DeleteMasterQuestion", dummyMasterQuestion[0]).Return(nil)
	useCaseTest := NewMasterQuestion(suite.repositoryTest)
	err := useCaseTest.DeleteMasterQuestion(dummyMasterQuestion[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterQuestionDeliveryTestSuite) TestGetAllMasterQuestions() {
	suite.repositoryTest.(*repoMockMasterQuestion).On("GetAllMasterQuestions", dummyMasterQuestion).Return(dummyMasterQuestion)
	useCaseTest := NewMasterQuestion(suite.repositoryTest)
	dummyUser := useCaseTest.GetAllMasterQuestions()
	assert.Equal(suite.T(), dummyMasterQuestion, dummyUser)
}

func (suite *MasterQuestionDeliveryTestSuite) TestGetMasterQuestion() {
	suite.repositoryTest.(*repoMockMasterQuestion).On("GetMasterQuestion", dummyMasterQuestion[0].ID).Return(dummyMasterQuestion[0], nil)
	useCaseTest := NewMasterQuestion(suite.repositoryTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyUser := useCaseTest.GetMasterQuestion(c)
	assert.NotNil(suite.T(), dummyUser[0])
	assert.Equal(suite.T(), dummyMasterQuestion[0], dummyUser[0])
}

func TestMasterQuestionDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(MasterQuestionDeliveryTestSuite))
}
