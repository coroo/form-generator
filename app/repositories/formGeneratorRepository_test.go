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

type FormGeneratorRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func (suite *FormGeneratorRepositoryTestSuite) SetupTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *FormGeneratorRepositoryTestSuite) TestBuildNewFormGeneratorRepository() {
	repoTest := NewFormGeneratorRepository()
	var dummyImpl *FormGeneratorRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *FormGeneratorRepositoryTestSuite) TestFormGeneratorCreate() {
	repoTest := NewFormGeneratorRepository()
	dummyFormGenerator := entity.FormGenerator{
		MasterQuestionId: 1,
		MasterFormId:     1,
	}
	userDummy := repoTest.SaveFormGenerator(dummyFormGenerator)
	assert.Nil(suite.T(), userDummy)
}

func (suite *FormGeneratorRepositoryTestSuite) TestFormGeneratorUpdate() {
	repoTest := NewFormGeneratorRepository()
	dummyFormGenerator := entity.FormGenerator{
		ID:               1,
		MasterQuestionId: 1,
		MasterFormId:     1,
	}
	userDummy := repoTest.UpdateFormGenerator(dummyFormGenerator)
	assert.Nil(suite.T(), userDummy)
}

func (suite *FormGeneratorRepositoryTestSuite) TestFormGeneratorDelete() {
	repoTest := NewFormGeneratorRepository()
	dummyFormGenerator := entity.FormGenerator{
		ID: 1,
	}
	userDummy := repoTest.DeleteFormGenerator(dummyFormGenerator)
	assert.Nil(suite.T(), userDummy)
}

func (suite *FormGeneratorRepositoryTestSuite) TestGetAllFormGenerators() {
	repoTest := NewFormGeneratorRepository()
	userDummy := repoTest.GetAllFormGenerators()
	assert.NotNil(suite.T(), userDummy)
}

func (suite *FormGeneratorRepositoryTestSuite) TestGetFormGenerator() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	repoTest := NewFormGeneratorRepository()
	userDummy := repoTest.GetFormGenerator(c)
	assert.NotNil(suite.T(), userDummy)
}

func TestFormGeneratorRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(FormGeneratorRepositoryTestSuite))
}
