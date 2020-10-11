package usecases

import (
	entity "github.com/coroo/form-generator/app/entity"
	repositories "github.com/coroo/form-generator/app/repositories"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type MasterAnswerService interface {
	Save(entity.MasterAnswer) error
	Update(entity.MasterAnswer) error
	Delete(entity.MasterAnswer) error
	GetAllMasterAnswers() []entity.MasterAnswer
	GetMasterAnswer(ctx *gin.Context) []entity.MasterAnswer
	GetMasterAnswerByQuestion(ctx *gin.Context) []entity.MasterAnswer
}

type masterAnswerService struct {
	repositories repositories.MasterAnswerRepository
}

func NewMasterAnswer(masterAnswerRepository repositories.MasterAnswerRepository) MasterAnswerService {
	return &masterAnswerService{
		repositories: masterAnswerRepository,
	}
}

func (usecases *masterAnswerService) GetAllMasterAnswers() []entity.MasterAnswer {
	return usecases.repositories.GetAllMasterAnswers()
}

func (usecases *masterAnswerService) GetMasterAnswer(ctx *gin.Context) []entity.MasterAnswer {
	return usecases.repositories.GetMasterAnswer(ctx)
}

func (usecases *masterAnswerService) GetMasterAnswerByQuestion(ctx *gin.Context) []entity.MasterAnswer {
	return usecases.repositories.GetMasterAnswerByQuestion(ctx)
}

func (usecases *masterAnswerService) Save(masterAnswer entity.MasterAnswer) error {
	usecases.repositories.Save(masterAnswer)
	return nil
}

func (usecases *masterAnswerService) Update(masterAnswer entity.MasterAnswer) error {
	usecases.repositories.Update(masterAnswer)
	return nil
}

func (usecases *masterAnswerService) Delete(masterAnswer entity.MasterAnswer) error {
	usecases.repositories.Delete(masterAnswer)
	return nil
}
