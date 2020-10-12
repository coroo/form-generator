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

// func (suite *MasterQuestionRepositoryTestSuite) TestGetMasterQuestion() {
// 	repoTest := NewMasterQuestionRepository()
// 	suite.ctx.Param("id") = 1
// 	// ctx := suite.context.WithValue(c.Request.Context(), "GinContextKey", c)
// 	userDummy := repoTest.GetMasterQuestion(suite.ctx)
// 	assert.NotNil(suite.T(), userDummy)
// }

// func (suite *MasterQuestionRepositoryTestSuite) TestFindOneById() {
// 	repoTest := NewMasterQuestionRepository(suite.db)
// 	userDummy, err := repoTest.FindOneById("1")
// 	assert.Nil(suite.T(), err)
// 	assert.Equal(suite.T(), "1", userDummy.Id)
// }

// func (suite *MasterQuestionRepositoryTestSuite) TestCreate() {
// 	repoTest := NewMasterQuestionRepository(suite.db)
// 	dummyUser := &models.User{
// 		FirstName: "Fn 3",
// 		LastName:  "Ln 3",
// 		IsActive:  "A",
// 	}
// 	err := repoTest.Create(dummyUser)
// 	assert.Nil(suite.T(), err)
// 	userDummy, err := repoTest.FindOneById(dummyUser.Id)
// 	assert.Nil(suite.T(), err)
// 	assert.Equal(suite.T(), "Fn 3", userDummy.FirstName)
// }
func TestMasterQuestionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(MasterQuestionRepositoryTestSuite))
}
