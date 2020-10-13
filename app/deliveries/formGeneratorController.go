package deliveries

import (
	entity "github.com/coroo/form-generator/app/entity"
	usecases "github.com/coroo/form-generator/app/usecases"

	"github.com/gin-gonic/gin"
)

type FormGeneratorController interface {
	Save(formGenerator entity.FormGenerator) error
	Update(formGenerator entity.FormGenerator) error
	Delete(formGenerator entity.FormGenerator) error
	GetAllFormGenerators() []entity.FormGenerator
	GetFormGenerator(ctx *gin.Context) []entity.FormGenerator
}

type formGeneratorDeliveries struct {
	usecases usecases.FormGeneratorService
}

func NewFormGenerator(usecases usecases.FormGeneratorService) FormGeneratorController {
	return &formGeneratorDeliveries{
		usecases: usecases,
	}
}

func (c *formGeneratorDeliveries) GetAllFormGenerators() []entity.FormGenerator {
	return c.usecases.GetAllFormGenerators()
}

func (c *formGeneratorDeliveries) GetFormGenerator(ctx *gin.Context) []entity.FormGenerator {
	return c.usecases.GetFormGenerator(ctx)
}

func (c *formGeneratorDeliveries) Save(formGenerator entity.FormGenerator) error {
	c.usecases.SaveFormGenerator(formGenerator)
	return nil
}

func (c *formGeneratorDeliveries) Update(formGenerator entity.FormGenerator) error {
	c.usecases.UpdateFormGenerator(formGenerator)
	return nil
}

func (c *formGeneratorDeliveries) Delete(formGenerator entity.FormGenerator) error {
	c.usecases.DeleteFormGenerator(formGenerator)
	return nil
}
