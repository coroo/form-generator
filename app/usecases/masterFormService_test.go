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
var dummyMasterForm = []entity.MasterForm{
	entity.MasterForm{
		ID:        1,
		FormName:  "Dummy First Name 1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, entity.MasterForm{
		ID:        2,
		FormName:  "Dummy Last Name 2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

type repoMockMasterForm struct {
	mock.Mock
}

func (r *repoMockMasterForm) SaveMasterForm(masterForm entity.MasterForm) error {
	args := r.Called(masterForm)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockMasterForm) UpdateMasterForm(masterForm entity.MasterForm) error {
	args := r.Called(masterForm)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockMasterForm) DeleteMasterForm(masterForm entity.MasterForm) error {
	args := r.Called(masterForm)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockMasterForm) GetAllMasterForms() []entity.MasterForm {
	return dummyMasterForm
}

func (r *repoMockMasterForm) GetMasterForm(ctx *gin.Context) []entity.MasterForm {
	return dummyMasterForm
}

func (r *repoMockMasterForm) CloseDB() {
}

type MasterFormDeliveryTestSuite struct {
	suite.Suite
	repositoryTest repositories.MasterFormRepository
}

func (suite *MasterFormDeliveryTestSuite) SetupTest() {
	suite.repositoryTest = new(repoMockMasterForm)
}

func (suite *MasterFormDeliveryTestSuite) TestBuildMasterFormService() {
	resultTest := NewMasterForm(suite.repositoryTest)
	var dummyImpl *MasterFormService
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
	// assert.NotNil(suite.T(), resultTest.(*MasterFormService).repositories)
}

func (suite *MasterFormDeliveryTestSuite) TestSaveMasterFormDelivery() {
	suite.repositoryTest.(*repoMockMasterForm).On("SaveMasterForm", dummyMasterForm[0]).Return(nil)
	useCaseTest := NewMasterForm(suite.repositoryTest)
	err := useCaseTest.SaveMasterForm(dummyMasterForm[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterFormDeliveryTestSuite) TestUpdateMasterFormDelivery() {
	suite.repositoryTest.(*repoMockMasterForm).On("UpdateMasterForm", dummyMasterForm[0]).Return(nil)
	useCaseTest := NewMasterForm(suite.repositoryTest)
	err := useCaseTest.UpdateMasterForm(dummyMasterForm[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterFormDeliveryTestSuite) TestDeleteMasterFormDelivery() {
	suite.repositoryTest.(*repoMockMasterForm).On("DeleteMasterForm", dummyMasterForm[0]).Return(nil)
	useCaseTest := NewMasterForm(suite.repositoryTest)
	err := useCaseTest.DeleteMasterForm(dummyMasterForm[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterFormDeliveryTestSuite) TestGetAllMasterForms() {
	suite.repositoryTest.(*repoMockMasterForm).On("GetAllMasterForms", dummyMasterForm).Return(dummyMasterForm)
	useCaseTest := NewMasterForm(suite.repositoryTest)
	dummyUser := useCaseTest.GetAllMasterForms()
	assert.Equal(suite.T(), dummyMasterForm, dummyUser)
}

func (suite *MasterFormDeliveryTestSuite) TestGetMasterForm() {
	suite.repositoryTest.(*repoMockMasterForm).On("GetMasterForm", dummyMasterForm[0].ID).Return(dummyMasterForm[0], nil)
	useCaseTest := NewMasterForm(suite.repositoryTest)

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyUser := useCaseTest.GetMasterForm(c)
	assert.NotNil(suite.T(), dummyUser[0])
	assert.Equal(suite.T(), dummyMasterForm[0], dummyUser[0])
}

func TestMasterFormDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(MasterFormDeliveryTestSuite))
}
