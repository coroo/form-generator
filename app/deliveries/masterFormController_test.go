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

type serviceMockMasterForm struct {
	mock.Mock
}

func (r *serviceMockMasterForm) SaveMasterForm(masterForm entity.MasterForm) error {
	args := r.Called(masterForm)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *serviceMockMasterForm) UpdateMasterForm(masterForm entity.MasterForm) error {
	args := r.Called(masterForm)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *serviceMockMasterForm) DeleteMasterForm(masterForm entity.MasterForm) error {
	args := r.Called(masterForm)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *serviceMockMasterForm) GetAllMasterForms() []entity.MasterForm {
	return dummyMasterForm
}

func (r *serviceMockMasterForm) GetMasterForm(ctx *gin.Context) []entity.MasterForm {
	return dummyMasterForm
}

type MasterFormDeliveryTestSuite struct {
	suite.Suite
	serviceTest usecases.MasterFormService
}

func (suite *MasterFormDeliveryTestSuite) SetupTest() {
	suite.serviceTest = new(serviceMockMasterForm)
}

func (suite *MasterFormDeliveryTestSuite) TestBuildMasterFormController() {
	resultTest := NewMasterForm(suite.serviceTest)
	var dummyImpl *MasterFormController
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
}

func (suite *MasterFormDeliveryTestSuite) TestSaveMasterFormDelivery() {
	suite.serviceTest.(*serviceMockMasterForm).On("SaveMasterForm", dummyMasterForm[0]).Return(nil)
	deliveryTest := NewMasterForm(suite.serviceTest)
	err := deliveryTest.Save(dummyMasterForm[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterFormDeliveryTestSuite) TestUpdateMasterFormDelivery() {
	suite.serviceTest.(*serviceMockMasterForm).On("UpdateMasterForm", dummyMasterForm[0]).Return(nil)
	deliveryTest := NewMasterForm(suite.serviceTest)
	err := deliveryTest.Update(dummyMasterForm[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterFormDeliveryTestSuite) TestDeleteMasterFormDelivery() {
	suite.serviceTest.(*serviceMockMasterForm).On("DeleteMasterForm", dummyMasterForm[0]).Return(nil)
	deliveryTest := NewMasterForm(suite.serviceTest)
	err := deliveryTest.Delete(dummyMasterForm[0])
	assert.Nil(suite.T(), err)
}

func (suite *MasterFormDeliveryTestSuite) TestGetAllMasterForms() {
	suite.serviceTest.(*serviceMockMasterForm).On("GetAllMasterForms", dummyMasterForm).Return(dummyMasterForm)
	deliveryTest := NewMasterForm(suite.serviceTest)
	dummyUser := deliveryTest.GetAllMasterForms()
	assert.Equal(suite.T(), dummyMasterForm, dummyUser)
}

func (suite *MasterFormDeliveryTestSuite) TestGetMasterForm() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	suite.serviceTest.(*serviceMockMasterForm).On("GetMasterForm", dummyMasterForm[0].ID).Return(dummyMasterForm[0], nil)
	deliveryTest := NewMasterForm(suite.serviceTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyUser := deliveryTest.GetMasterForm(c)
	assert.NotNil(suite.T(), dummyUser[0])
	assert.Equal(suite.T(), dummyMasterForm[0], dummyUser[0])
}

func TestMasterFormDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(MasterFormDeliveryTestSuite))
}
