package usecases

import (
	entity "github.com/coroo/form-generator/app/entity"
	repositories "github.com/coroo/form-generator/app/repositories"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type MasterFormService interface {
	SaveMasterForm(entity.MasterForm) error
	UpdateMasterForm(entity.MasterForm) error
	DeleteMasterForm(entity.MasterForm) error
	GetAllMasterForms() []entity.MasterForm
	GetMasterForm(ctx *gin.Context) []entity.MasterForm
}

type masterFormService struct {
	repositories repositories.MasterFormRepository
}

func NewMasterForm(masterFormRepository repositories.MasterFormRepository) MasterFormService {
	return &masterFormService{
		repositories: masterFormRepository,
	}
}

func (usecases *masterFormService) GetAllMasterForms() []entity.MasterForm {
	return usecases.repositories.GetAllMasterForms()
}

func (usecases *masterFormService) GetMasterForm(ctx *gin.Context) []entity.MasterForm {
	return usecases.repositories.GetMasterForm(ctx)
}

func (usecases *masterFormService) SaveMasterForm(masterForm entity.MasterForm) error {
	usecases.repositories.SaveMasterForm(masterForm)
	return nil
}

func (usecases *masterFormService) UpdateMasterForm(masterForm entity.MasterForm) error {
	usecases.repositories.UpdateMasterForm(masterForm)
	return nil
}

func (usecases *masterFormService) DeleteMasterForm(masterForm entity.MasterForm) error {
	usecases.repositories.DeleteMasterForm(masterForm)
	return nil
}
