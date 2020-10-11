package deliveries

import (
	entity "github.com/coroo/form-generator/app/entity"
	usecases "github.com/coroo/form-generator/app/usecases"

	"github.com/gin-gonic/gin"
)

type MasterFormController interface {
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
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

func (c *masterFormDeliveries) Save(ctx *gin.Context) error {
	var masterForm entity.MasterForm
	err := ctx.ShouldBindJSON(&masterForm)
	if err != nil {
		return err
	}
	c.usecases.Save(masterForm)
	return nil
}

func (c *masterFormDeliveries) Update(ctx *gin.Context) error {
	var masterForm entity.MasterForm
	err := ctx.ShouldBindJSON(&masterForm)
	if err != nil {
		return err
	}
	c.usecases.Update(masterForm)
	return nil
}

func (c *masterFormDeliveries) Delete(ctx *gin.Context) error {
	var masterForm entity.MasterForm
	err := ctx.ShouldBindJSON(&masterForm)
	if err != nil {
		return err
	}
	c.usecases.Delete(masterForm)
	return nil
}
