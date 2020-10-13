package deliveries

import (
	entity "github.com/coroo/form-generator/app/entity"
	usecases "github.com/coroo/form-generator/app/usecases"

	"github.com/gin-gonic/gin"
)

type MasterAnswerController interface {
	Save(masterAnswer entity.MasterAnswer) error
	Update(masterAnswer entity.MasterAnswer) error
	Delete(masterAnswer entity.MasterAnswer) error
	GetAllMasterAnswers() []entity.MasterAnswer
	GetMasterAnswer(ctx *gin.Context) []entity.MasterAnswer
	GetMasterAnswerByQuestion(ctx *gin.Context) []entity.MasterAnswer
}

type masterAnswerDeliveries struct {
	usecases usecases.MasterAnswerService
}

func NewMasterAnswer(usecases usecases.MasterAnswerService) MasterAnswerController {
	return &masterAnswerDeliveries{
		usecases: usecases,
	}
}

func (c *masterAnswerDeliveries) GetAllMasterAnswers() []entity.MasterAnswer {
	return c.usecases.GetAllMasterAnswers()
}

func (c *masterAnswerDeliveries) GetMasterAnswer(ctx *gin.Context) []entity.MasterAnswer {
	return c.usecases.GetMasterAnswer(ctx)
}

func (c *masterAnswerDeliveries) GetMasterAnswerByQuestion(ctx *gin.Context) []entity.MasterAnswer {
	return c.usecases.GetMasterAnswerByQuestion(ctx)
}

func (c *masterAnswerDeliveries) Save(masterAnswer entity.MasterAnswer) error {
	c.usecases.SaveMasterAnswer(masterAnswer)
	return nil
}

func (c *masterAnswerDeliveries) Update(masterAnswer entity.MasterAnswer) error {
	c.usecases.UpdateMasterAnswer(masterAnswer)
	return nil
}

func (c *masterAnswerDeliveries) Delete(masterAnswer entity.MasterAnswer) error {
	c.usecases.DeleteMasterAnswer(masterAnswer)
	return nil
}
