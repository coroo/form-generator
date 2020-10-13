package repositories

import (
	"net/http/httptest"
	"testing"

	entity "github.com/coroo/form-generator/app/entity"
	"github.com/coroo/form-generator/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MasterQuestionRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func (suite *MasterQuestionRepositoryTestSuite) SetupTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *MasterQuestionRepositoryTestSuite) TestBuildNewMasterQuestionRepository() {
	repoTest := NewMasterQuestionRepository()
	var dummyImpl *MasterQuestionRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *MasterQuestionRepositoryTestSuite) TestMasterQuestionCreate() {
	repoTest := NewMasterQuestionRepository()
	dummyMasterAnswer := entity.MasterQuestion{
		QuestionText: "Fn 3",
	}
	userDummy := repoTest.SaveMasterQuestion(dummyMasterAnswer)
	assert.Nil(suite.T(), userDummy)
}

func (suite *MasterQuestionRepositoryTestSuite) TestMasterQuestionUpdate() {
	repoTest := NewMasterQuestionRepository()
	dummyMasterAnswer := entity.MasterQuestion{
		ID:           1,
		QuestionText: "Fn 3",
	}
	userDummy := repoTest.UpdateMasterQuestion(dummyMasterAnswer)
	assert.Nil(suite.T(), userDummy)
}

func (suite *MasterQuestionRepositoryTestSuite) TestMasterQuestionDelete() {
	repoTest := NewMasterQuestionRepository()
	dummyMasterAnswer := entity.MasterQuestion{
		ID: 1,
	}
	userDummy := repoTest.DeleteMasterQuestion(dummyMasterAnswer)
	assert.Nil(suite.T(), userDummy)
}

func (suite *MasterQuestionRepositoryTestSuite) TestGetAllMasterQuestions() {
	repoTest := NewMasterQuestionRepository()
	userDummy := repoTest.GetAllMasterQuestions()
	assert.NotNil(suite.T(), userDummy)
}

type ContextMock struct {
	JSONCalled bool
}

func (c *ContextMock) JSON(code int, obj interface{}) {
	c.JSONCalled = true
}

func (suite *MasterQuestionRepositoryTestSuite) TestGetMasterQuestion() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	repoTest := NewMasterQuestionRepository()
	userDummy := repoTest.GetMasterQuestion(c)
	assert.NotNil(suite.T(), userDummy)
}

func TestMasterQuestionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(MasterQuestionRepositoryTestSuite))
}
