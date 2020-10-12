package usecases

import (
	entity "github.com/coroo/form-generator/app/entity"
	repositories "github.com/coroo/form-generator/app/repositories"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type FormGeneratorService interface {
	SaveFormGenerator(entity.FormGenerator) error
	UpdateFormGenerator(entity.FormGenerator) error
	DeleteFormGenerator(entity.FormGenerator) error
	GetAllFormGenerators() []entity.FormGenerator
	GetFormGenerator(ctx *gin.Context) []entity.FormGenerator
}

type formGeneratorService struct {
	// formGenerator []entity.FormGenerator
	// models models.FormGeneratorModel
	repositories repositories.FormGeneratorRepository
}

func NewFormGenerator(formGeneratorRepository repositories.FormGeneratorRepository) FormGeneratorService {
	// return &formGeneratorService{
	// 	models: models,
	// }
	return &formGeneratorService{
		repositories: formGeneratorRepository,
	}
}

func (usecases *formGeneratorService) GetAllFormGenerators() []entity.FormGenerator {
	return usecases.repositories.GetAllFormGenerators()
}

func (usecases *formGeneratorService) GetFormGenerator(ctx *gin.Context) []entity.FormGenerator {
	return usecases.repositories.GetFormGenerator(ctx)
}

func (usecases *formGeneratorService) SaveFormGenerator(formGenerator entity.FormGenerator) error {
	usecases.repositories.SaveFormGenerator(formGenerator)
	return nil
}

func (usecases *formGeneratorService) UpdateFormGenerator(formGenerator entity.FormGenerator) error {
	usecases.repositories.UpdateFormGenerator(formGenerator)
	return nil
}

func (usecases *formGeneratorService) DeleteFormGenerator(formGenerator entity.FormGenerator) error {
	usecases.repositories.DeleteFormGenerator(formGenerator)
	return nil
}
