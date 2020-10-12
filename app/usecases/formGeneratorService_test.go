package usecases

import (
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

type repoMockFormGenerator struct {
	mock.Mock
}

func (r *repoMockFormGenerator) SaveFormGenerator(masterQuestion entity.FormGenerator) error {
	args := r.Called(masterQuestion)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockFormGenerator) UpdateFormGenerator(masterQuestion entity.FormGenerator) error {
	args := r.Called(masterQuestion)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockFormGenerator) DeleteFormGenerator(masterQuestion entity.FormGenerator) error {
	args := r.Called(masterQuestion)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockFormGenerator) GetAllFormGenerators() []entity.FormGenerator {
	return dummyFormGenerator
}

func (r *repoMockFormGenerator) CloseDB() {
}

func (r *repoMockFormGenerator) GetFormGenerator(ctx *gin.Context) []entity.FormGenerator {
	return nil
}

type FormGeneratorDeliveryTestSuite struct {
	suite.Suite
	repositoryTest repositories.FormGeneratorRepository
}

func (suite *FormGeneratorDeliveryTestSuite) SetupTest() {
	suite.repositoryTest = new(repoMockFormGenerator)
}

func (suite *FormGeneratorDeliveryTestSuite) TestBuildFormGeneratorService() {
	resultTest := NewFormGenerator(suite.repositoryTest)
	var dummyImpl *FormGeneratorService
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
	// assert.NotNil(suite.T(), resultTest.(*FormGeneratorService).repositories)
}

func (suite *FormGeneratorDeliveryTestSuite) TestSaveFormGeneratorDelivery() {
	suite.repositoryTest.(*repoMockFormGenerator).On("SaveFormGenerator", dummyFormGenerator[0]).Return(nil)
	useCaseTest := NewFormGenerator(suite.repositoryTest)
	err := useCaseTest.SaveFormGenerator(dummyFormGenerator[0])
	assert.Nil(suite.T(), err)
}

func (suite *FormGeneratorDeliveryTestSuite) TestUpdateFormGeneratorDelivery() {
	suite.repositoryTest.(*repoMockFormGenerator).On("UpdateFormGenerator", dummyFormGenerator[0]).Return(nil)
	useCaseTest := NewFormGenerator(suite.repositoryTest)
	err := useCaseTest.UpdateFormGenerator(dummyFormGenerator[0])
	assert.Nil(suite.T(), err)
}

func (suite *FormGeneratorDeliveryTestSuite) TestDeleteFormGeneratorDelivery() {
	suite.repositoryTest.(*repoMockFormGenerator).On("DeleteFormGenerator", dummyFormGenerator[0]).Return(nil)
	useCaseTest := NewFormGenerator(suite.repositoryTest)
	err := useCaseTest.DeleteFormGenerator(dummyFormGenerator[0])
	assert.Nil(suite.T(), err)
}

func (suite *FormGeneratorDeliveryTestSuite) TestGetAllFormGenerators() {
	suite.repositoryTest.(*repoMockFormGenerator).On("GetAllFormGenerators", dummyFormGenerator).Return(dummyFormGenerator)
	useCaseTest := NewFormGenerator(suite.repositoryTest)
	dummyUser := useCaseTest.GetAllFormGenerators()
	assert.Equal(suite.T(), dummyFormGenerator, dummyUser)
}

// func (suite *FormGeneratorDeliveryTestSuite) TestGetFormGenerator() {
// 	suite.repositoryTest.(*repoMockFormGenerator).On("FindOneById", dummyFormGenerator[0].ID).Return(dummyFormGenerator[0], nil)
// 	useCaseTest := NewFormGenerator(suite.repositoryTest)
// 	dummyUser := useCaseTest.GetFormGenerator(dummyFormGenerator[0].ID)
// 	assert.Equal(suite.T(), dummyFormGenerator[0].ID, dummyUser.ID)
// }

func TestFormGeneratorDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(FormGeneratorDeliveryTestSuite))
}
