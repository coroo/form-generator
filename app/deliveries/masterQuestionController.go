package deliveries

import (
	entity "github.com/coroo/form-generator/app/entity"
	usecases "github.com/coroo/form-generator/app/usecases"

	"github.com/gin-gonic/gin"
)

type MasterQuestionController interface {
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
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

func (c *masterQuestionDeliveries) Save(ctx *gin.Context) error {
	var masterQuestion entity.MasterQuestion
	err := ctx.ShouldBindJSON(&masterQuestion)
	if err != nil {
		return err
	}
	c.usecases.SaveMasterQuestion(masterQuestion)
	return nil
}

func (c *masterQuestionDeliveries) Update(ctx *gin.Context) error {
	var masterQuestion entity.MasterQuestion
	err := ctx.ShouldBindJSON(&masterQuestion)
	if err != nil {
		return err
	}
	c.usecases.UpdateMasterQuestion(masterQuestion)
	return nil
}

func (c *masterQuestionDeliveries) Delete(ctx *gin.Context) error {
	var masterQuestion entity.MasterQuestion
	err := ctx.ShouldBindJSON(&masterQuestion)
	if err != nil {
		return err
	}
	c.usecases.DeleteMasterQuestion(masterQuestion)
	return nil
}
