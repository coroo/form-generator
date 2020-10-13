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
var dummyFormGenerator = []entity.FormGenerator{
	entity.FormGenerator{
		ID:               1,
		MasterQuestionId: 1,
		MasterFormId:     1,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}, entity.FormGenerator{
		ID:               2,
		MasterQuestionId: 1,
		MasterFormId:     1,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	},
}

type serviceMockFormGenerator struct {
	mock.Mock
}

func (r *serviceMockFormGenerator) SaveFormGenerator(formGenerator entity.FormGenerator) error {
	args := r.Called(formGenerator)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *serviceMockFormGenerator) UpdateFormGenerator(formGenerator entity.FormGenerator) error {
	args := r.Called(formGenerator)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *serviceMockFormGenerator) DeleteFormGenerator(formGenerator entity.FormGenerator) error {
	args := r.Called(formGenerator)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *serviceMockFormGenerator) GetAllFormGenerators() []entity.FormGenerator {
	return dummyFormGenerator
}

func (r *serviceMockFormGenerator) GetFormGenerator(ctx *gin.Context) []entity.FormGenerator {
	return dummyFormGenerator
}

type FormGeneratorDeliveryTestSuite struct {
	suite.Suite
	serviceTest usecases.FormGeneratorService
}

func (suite *FormGeneratorDeliveryTestSuite) SetupTest() {
	suite.serviceTest = new(serviceMockFormGenerator)
}

func (suite *FormGeneratorDeliveryTestSuite) TestBuildFormGeneratorController() {
	resultTest := NewFormGenerator(suite.serviceTest)
	var dummyImpl *FormGeneratorController
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
}

func (suite *FormGeneratorDeliveryTestSuite) TestSaveFormGeneratorDelivery() {
	suite.serviceTest.(*serviceMockFormGenerator).On("SaveFormGenerator", dummyFormGenerator[0]).Return(nil)
	deliveryTest := NewFormGenerator(suite.serviceTest)
	err := deliveryTest.Save(dummyFormGenerator[0])
	assert.Nil(suite.T(), err)
}

func (suite *FormGeneratorDeliveryTestSuite) TestUpdateFormGeneratorDelivery() {
	suite.serviceTest.(*serviceMockFormGenerator).On("UpdateFormGenerator", dummyFormGenerator[0]).Return(nil)
	deliveryTest := NewFormGenerator(suite.serviceTest)
	err := deliveryTest.Update(dummyFormGenerator[0])
	assert.Nil(suite.T(), err)
}

func (suite *FormGeneratorDeliveryTestSuite) TestDeleteFormGeneratorDelivery() {
	suite.serviceTest.(*serviceMockFormGenerator).On("DeleteFormGenerator", dummyFormGenerator[0]).Return(nil)
	deliveryTest := NewFormGenerator(suite.serviceTest)
	err := deliveryTest.Delete(dummyFormGenerator[0])
	assert.Nil(suite.T(), err)
}

func (suite *FormGeneratorDeliveryTestSuite) TestGetAllFormGenerators() {
	suite.serviceTest.(*serviceMockFormGenerator).On("GetAllFormGenerators", dummyFormGenerator).Return(dummyFormGenerator)
	deliveryTest := NewFormGenerator(suite.serviceTest)
	dummyUser := deliveryTest.GetAllFormGenerators()
	assert.Equal(suite.T(), dummyFormGenerator, dummyUser)
}

func (suite *FormGeneratorDeliveryTestSuite) TestGetFormGenerator() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	suite.serviceTest.(*serviceMockFormGenerator).On("GetFormGenerator", dummyFormGenerator[0].ID).Return(dummyFormGenerator[0], nil)
	deliveryTest := NewFormGenerator(suite.serviceTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyUser := deliveryTest.GetFormGenerator(c)
	assert.NotNil(suite.T(), dummyUser[0])
	assert.Equal(suite.T(), dummyFormGenerator[0], dummyUser[0])
}

func TestFormGeneratorDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(FormGeneratorDeliveryTestSuite))
}
