package usecases

import (
	entity "github.com/coroo/form-generator/app/entity"
	repositories "github.com/coroo/form-generator/app/repositories"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type MasterQuestionService interface {
	SaveMasterQuestion(entity.MasterQuestion) error
	UpdateMasterQuestion(entity.MasterQuestion) error
	DeleteMasterQuestion(entity.MasterQuestion) error
	GetAllMasterQuestions() []entity.MasterQuestion
	GetMasterQuestion(ctx *gin.Context) []entity.MasterQuestion
}

type masterQuestionService struct {
	repositories repositories.MasterQuestionRepository
}

func NewMasterQuestion(masterQuestionRepository repositories.MasterQuestionRepository) MasterQuestionService {
	return &masterQuestionService{
		repositories: masterQuestionRepository,
	}
}

func (usecases *masterQuestionService) GetAllMasterQuestions() []entity.MasterQuestion {
	return usecases.repositories.GetAllMasterQuestions()
}

func (usecases *masterQuestionService) GetMasterQuestion(ctx *gin.Context) []entity.MasterQuestion {
	return usecases.repositories.GetMasterQuestion(ctx)
}

func (usecases *masterQuestionService) SaveMasterQuestion(masterQuestion entity.MasterQuestion) error {
	usecases.repositories.SaveMasterQuestion(masterQuestion)
	return nil
}

func (usecases *masterQuestionService) UpdateMasterQuestion(masterQuestion entity.MasterQuestion) error {
	usecases.repositories.UpdateMasterQuestion(masterQuestion)
	return nil
}

func (usecases *masterQuestionService) DeleteMasterQuestion(masterQuestion entity.MasterQuestion) error {
	usecases.repositories.DeleteMasterQuestion(masterQuestion)
	return nil
}
