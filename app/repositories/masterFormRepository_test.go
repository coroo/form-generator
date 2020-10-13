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

type MasterFormRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func (suite *MasterFormRepositoryTestSuite) SetupTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *MasterFormRepositoryTestSuite) TestBuildNewMasterFormRepository() {
	repoTest := NewMasterFormRepository()
	var dummyImpl *MasterFormRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *MasterFormRepositoryTestSuite) TestMasterFormCreate() {
	repoTest := NewMasterFormRepository()
	dummyMasterAnswer := entity.MasterForm{
		FormName: "Fn 3",
	}
	userDummy := repoTest.SaveMasterForm(dummyMasterAnswer)
	assert.Nil(suite.T(), userDummy)
}

func (suite *MasterFormRepositoryTestSuite) TestMasterFormUpdate() {
	repoTest := NewMasterFormRepository()
	dummyMasterAnswer := entity.MasterForm{
		ID:       1,
		FormName: "Fn 3",
	}
	userDummy := repoTest.UpdateMasterForm(dummyMasterAnswer)
	assert.Nil(suite.T(), userDummy)
}

func (suite *MasterFormRepositoryTestSuite) TestMasterFormDelete() {
	repoTest := NewMasterFormRepository()
	dummyMasterAnswer := entity.MasterForm{
		ID: 1,
	}
	userDummy := repoTest.DeleteMasterForm(dummyMasterAnswer)
	assert.Nil(suite.T(), userDummy)
}

func (suite *MasterFormRepositoryTestSuite) TestGetAllMasterForms() {
	repoTest := NewMasterFormRepository()
	userDummy := repoTest.GetAllMasterForms()
	assert.NotNil(suite.T(), userDummy)
}

func (suite *MasterFormRepositoryTestSuite) TestGetMasterForm() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	repoTest := NewMasterFormRepository()
	userDummy := repoTest.GetMasterForm(c)
	assert.NotNil(suite.T(), userDummy)
}

func TestMasterFormRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(MasterFormRepositoryTestSuite))
}
