package routes

import (
	"net/http"

	deliveries "github.com/coroo/form-generator/app/deliveries"
	"github.com/coroo/form-generator/app/entity"
	repositories "github.com/coroo/form-generator/app/repositories"
	usecases "github.com/coroo/form-generator/app/usecases"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	formGeneratorRepository repositories.FormGeneratorRepository = repositories.NewFormGeneratorRepository()
	formGeneratorService    usecases.FormGeneratorService        = usecases.NewFormGenerator(formGeneratorRepository)
	formGeneratorController deliveries.FormGeneratorController   = deliveries.NewFormGenerator(formGeneratorService)
)

// GetFormGeneratorsIndex godoc
// @Security basicAuth
// @Summary Show all existing Form_Generators
// @Description Get all existing Form_Generators
// @Tags Form_Generators
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.FormGenerator
// @Failure 401 {object} dto.Response
// @Router /formGenerator/index [get]
func FormGeneratorsIndex(c *gin.Context) {
	formGenerators := formGeneratorController.GetAllFormGenerators()
	c.JSON(http.StatusOK, gin.H{"data": formGenerators})
}

// GetFormGeneratorsDetail godoc
// @Security basicAuth
// @Summary Show an existing Form_Generators
// @Description Get detail the existing Form_Generators
// @Tags Form_Generators
// @Accept  json
// @Produce  json
// @Param  id path int true "Master Question ID"
// @Success 200 {array} entity.FormGenerator
// @Failure 401 {object} dto.Response
// @Router /formGenerator/detail/{id} [get]
func FormGeneratorsDetail(c *gin.Context) {
	formGenerator := formGeneratorController.GetFormGenerator(c)
	c.JSON(http.StatusOK, gin.H{"data": formGenerator})
}

// CreateFormGenerators godoc
// @Security basicAuth
// @Summary Create new Form_Generators
// @Description Create a new Form_Generators
// @Tags Form_Generators
// @Accept  json
// @Produce  json
// @Param formGenerator body entity.FormGeneratorCreate true "Create formGenerator"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /formGenerator/create [post]
func FormGeneratorCreate(c *gin.Context) {
	var formGeneratorEntity entity.FormGenerator
	c.ShouldBindJSON(&formGeneratorEntity)
	formGenerator := formGeneratorController.Save(formGeneratorEntity)
	c.JSON(http.StatusOK, formGenerator)
}

// UpdateFormGenerators godoc
// @Security basicAuth
// @Summary Update Form_Generators
// @Description Update a Form_Generators
// @Tags Form_Generators
// @Accept  json
// @Produce  json
// @Param formGenerator body entity.FormGenerator true "Update formGenerator"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /formGenerator/update [put]
func FormGeneratorUpdate(c *gin.Context) {
	var formGeneratorEntity entity.FormGenerator
	c.ShouldBindJSON(&formGeneratorEntity)
	formGenerator := formGeneratorController.Update(formGeneratorEntity)
	c.JSON(http.StatusOK, formGenerator)
}

// DeleteFormGenerators godoc
// @Security basicAuth
// @Summary Delete Form_Generators
// @Description Delete a Form_Generators
// @Tags Form_Generators
// @Accept  json
// @Produce  json
// @Param formGenerator body entity.FormGeneratorDelete true "Delete formGenerator"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /formGenerator/delete [delete]
func FormGeneratorDelete(c *gin.Context) {
	var formGeneratorEntity entity.FormGenerator
	c.ShouldBindJSON(&formGeneratorEntity)
	formGenerator := formGeneratorController.Delete(formGeneratorEntity)
	c.JSON(http.StatusOK, formGenerator)
}
