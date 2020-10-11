package deliveries

import (
	entity "github.com/coroo/form-generator/app/entity"
	usecases "github.com/coroo/form-generator/app/usecases"

	"github.com/gin-gonic/gin"
)

type FormGeneratorController interface {
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
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

func (c *formGeneratorDeliveries) Save(ctx *gin.Context) error {
	var formGenerator entity.FormGenerator
	err := ctx.ShouldBindJSON(&formGenerator)
	if err != nil {
		return err
	}
	c.usecases.Save(formGenerator)
	return nil
}

func (c *formGeneratorDeliveries) Update(ctx *gin.Context) error {
	var formGenerator entity.FormGenerator
	err := ctx.ShouldBindJSON(&formGenerator)
	if err != nil {
		return err
	}
	c.usecases.Update(formGenerator)
	return nil
}

func (c *formGeneratorDeliveries) Delete(ctx *gin.Context) error {
	var formGenerator entity.FormGenerator
	err := ctx.ShouldBindJSON(&formGenerator)
	if err != nil {
		return err
	}
	c.usecases.Delete(formGenerator)
	return nil
}
