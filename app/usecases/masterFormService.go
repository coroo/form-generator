package usecases

import (
	entity "github.com/coroo/form-generator/app/entity"
	repositories "github.com/coroo/form-generator/app/repositories"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type MasterFormService interface {
	Save(entity.MasterForm) error
	Update(entity.MasterForm) error
	Delete(entity.MasterForm) error
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

func (usecases *masterFormService) Save(masterForm entity.MasterForm) error {
	usecases.repositories.Save(masterForm)
	return nil
}

func (usecases *masterFormService) Update(masterForm entity.MasterForm) error {
	usecases.repositories.Update(masterForm)
	return nil
}

func (usecases *masterFormService) Delete(masterForm entity.MasterForm) error {
	usecases.repositories.Delete(masterForm)
	return nil
}
