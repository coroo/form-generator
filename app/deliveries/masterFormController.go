package deliveries

import (
	entity "github.com/coroo/form-generator/app/entity"
	usecases "github.com/coroo/form-generator/app/usecases"

	"github.com/gin-gonic/gin"
)

type MasterFormController interface {
	Save(masterForm entity.MasterForm) error
	Update(masterForm entity.MasterForm) error
	Delete(masterForm entity.MasterForm) error
	GetAllMasterForms() []entity.MasterForm
	GetMasterForm(ctx *gin.Context) []entity.MasterForm
}

type masterFormDeliveries struct {
	usecases usecases.MasterFormService
}

func NewMasterForm(usecases usecases.MasterFormService) MasterFormController {
	return &masterFormDeliveries{
		usecases: usecases,
	}
}

func (c *masterFormDeliveries) GetAllMasterForms() []entity.MasterForm {
	return c.usecases.GetAllMasterForms()
}

func (c *masterFormDeliveries) GetMasterForm(ctx *gin.Context) []entity.MasterForm {
	return c.usecases.GetMasterForm(ctx)
}

func (c *masterFormDeliveries) Save(masterForm entity.MasterForm) error {
	c.usecases.SaveMasterForm(masterForm)
	return nil
}

func (c *masterFormDeliveries) Update(masterForm entity.MasterForm) error {
	c.usecases.UpdateMasterForm(masterForm)
	return nil
}

func (c *masterFormDeliveries) Delete(masterForm entity.MasterForm) error {
	c.usecases.DeleteMasterForm(masterForm)
	return nil
}
