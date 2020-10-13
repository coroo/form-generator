package deliveries

import (
	entity "github.com/coroo/form-generator/app/entity"
	usecases "github.com/coroo/form-generator/app/usecases"

	"github.com/gin-gonic/gin"
)

type MasterQuestionController interface {
	Save(masterQuestion entity.MasterQuestion) error
	Update(masterQuestion entity.MasterQuestion) error
	Delete(masterQuestion entity.MasterQuestion) error
	GetAllMasterQuestions() []entity.MasterQuestion
	GetMasterQuestion(ctx *gin.Context) []entity.MasterQuestion
}

type masterQuestionDeliveries struct {
	usecases usecases.MasterQuestionService
}

func NewMasterQuestion(usecases usecases.MasterQuestionService) MasterQuestionController {
	return &masterQuestionDeliveries{
		usecases: usecases,
	}
}

func (c *masterQuestionDeliveries) GetAllMasterQuestions() []entity.MasterQuestion {
	return c.usecases.GetAllMasterQuestions()
}

func (c *masterQuestionDeliveries) GetMasterQuestion(ctx *gin.Context) []entity.MasterQuestion {
	return c.usecases.GetMasterQuestion(ctx)
}

func (c *masterQuestionDeliveries) Save(masterQuestion entity.MasterQuestion) error {
	c.usecases.SaveMasterQuestion(masterQuestion)
	return nil
}

func (c *masterQuestionDeliveries) Update(masterQuestion entity.MasterQuestion) error {
	c.usecases.UpdateMasterQuestion(masterQuestion)
	return nil
}

func (c *masterQuestionDeliveries) Delete(masterQuestion entity.MasterQuestion) error {
	c.usecases.DeleteMasterQuestion(masterQuestion)
	return nil
}
