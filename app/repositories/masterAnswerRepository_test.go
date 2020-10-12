package repositories

import (
	"testing"

	entity "github.com/coroo/form-generator/app/entity"
	"github.com/coroo/form-generator/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MasterAnswerRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func (suite *MasterAnswerRepositoryTestSuite) SetupTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *MasterAnswerRepositoryTestSuite) TestBuildNewMasterAnswerRepository() {
	repoTest := NewMasterAnswerRepository()
	var dummyImpl *MasterAnswerRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *MasterAnswerRepositoryTestSuite) TestMasterAnswerCreate() {
	repoTest := NewMasterAnswerRepository()
	dummyMasterAnswer := entity.MasterAnswer{
		MasterQuestionId: 1,
		FormAnswer:       "Fn 3",
		FormScore:        "60",
	}
	userDummy := repoTest.Save(dummyMasterAnswer)
	assert.Nil(suite.T(), userDummy)
}

func (suite *MasterAnswerRepositoryTestSuite) TestMasterAnswerUpdate() {
	repoTest := NewMasterAnswerRepository()
	dummyMasterAnswer := entity.MasterAnswer{
		ID:               1,
		MasterQuestionId: 1,
		FormAnswer:       "Fn 3",
		FormScore:        "60",
	}
	userDummy := repoTest.Update(dummyMasterAnswer)
	assert.Nil(suite.T(), userDummy)
}

func (suite *MasterAnswerRepositoryTestSuite) TestMasterAnswerDelete() {
	repoTest := NewMasterAnswerRepository()
	dummyMasterAnswer := entity.MasterAnswer{
		ID: 1,
	}
	userDummy := repoTest.Delete(dummyMasterAnswer)
	assert.Nil(suite.T(), userDummy)
}

func (suite *MasterAnswerRepositoryTestSuite) TestGetAllMasterAnswers() {
	repoTest := NewMasterAnswerRepository()
	userDummy := repoTest.GetAllMasterAnswers()
	assert.NotNil(suite.T(), userDummy)
}

func TestMasterAnswerRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(MasterAnswerRepositoryTestSuite))
}
