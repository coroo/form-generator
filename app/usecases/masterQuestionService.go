package usecases

import (
	entity "github.com/coroo/form-generator/app/entity"
	repositories "github.com/coroo/form-generator/app/repositories"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type MasterQuestionService interface {
	Save(entity.MasterQuestion) error
	Update(entity.MasterQuestion) error
	Delete(entity.MasterQuestion) error
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

func (usecases *masterQuestionService) Save(masterQuestion entity.MasterQuestion) error {
	usecases.repositories.Save(masterQuestion)
	return nil
}

func (usecases *masterQuestionService) Update(masterQuestion entity.MasterQuestion) error {
	usecases.repositories.Update(masterQuestion)
	return nil
}

func (usecases *masterQuestionService) Delete(masterQuestion entity.MasterQuestion) error {
	usecases.repositories.Delete(masterQuestion)
	return nil
}
