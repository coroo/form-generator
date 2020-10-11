package routes

import (
	"net/http"

	deliveries "github.com/coroo/form-generator/app/deliveries"
	repositories "github.com/coroo/form-generator/app/repositories"
	usecases "github.com/coroo/form-generator/app/usecases"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	masterFormRepository repositories.MasterFormRepository = repositories.NewMasterFormRepository()
	masterFormService    usecases.MasterFormService        = usecases.NewMasterForm(masterFormRepository)
	masterFormController deliveries.MasterFormController   = deliveries.NewMasterForm(masterFormService)
)

// GetMasterFormsIndex godoc
// @Security basicAuth
// @Summary Show all existing Master_Forms
// @Description Get all existing Master_Forms
// @Tags Master_Forms
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.MasterForm
// @Failure 401 {object} dto.Response
// @Router /masterForm/index [get]
func MasterFormsIndex(c *gin.Context) {
	masterForms := masterFormController.GetAllMasterForms()
	c.JSON(http.StatusOK, gin.H{"data": masterForms})
}

// GetMasterFormsDetail godoc
// @Security basicAuth
// @Summary Show an existing Master_Forms
// @Description Get detail the existing Master_Forms
// @Tags Master_Forms
// @Accept  json
// @Produce  json
// @Param  id path int true "Master Form ID"
// @Success 200 {array} entity.MasterForm
// @Failure 401 {object} dto.Response
// @Router /masterForm/detail/{id} [get]
func MasterFormsDetail(c *gin.Context) {
	masterForm := masterFormController.GetMasterForm(c)
	// buf := make([]byte, 1024)
	// num := c.Param("id")
	// reqBody := string(buf[0:num])
	// log.Print(num)
	c.JSON(http.StatusOK, gin.H{"data": masterForm})
}

// CreateMasterForms godoc
// @Security basicAuth
// @Summary Create new Master_Forms
// @Description Create a new Master_Forms
// @Tags Master_Forms
// @Accept  json
// @Produce  json
// @Param masterForm body entity.MasterFormCreate true "Create masterForm"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /masterForm/create [post]
func MasterFormCreate(c *gin.Context) {
	masterForm := masterFormController.Save(c)
	if masterForm != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": masterForm.Error()})
	} else {
		c.JSON(http.StatusOK, masterForm)
	}
}

// UpdateMasterForms godoc
// @Security basicAuth
// @Summary Update Master_Forms
// @Description Update a Master_Forms
// @Tags Master_Forms
// @Accept  json
// @Produce  json
// @Param masterForm body entity.MasterForm true "Update masterForm"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /masterForm/update [put]
func MasterFormUpdate(c *gin.Context) {
	masterForm := masterFormController.Update(c)
	if masterForm != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": masterForm.Error()})
	} else {
		c.JSON(http.StatusOK, masterForm)
	}
}

// DeleteMasterForms godoc
// @Security basicAuth
// @Summary Delete Master_Forms
// @Description Delete a Master_Forms
// @Tags Master_Forms
// @Accept  json
// @Produce  json
// @Param masterForm body entity.MasterFormDelete true "Delete masterForm"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /masterForm/delete [delete]
func MasterFormDelete(c *gin.Context) {
	masterForm := masterFormController.Delete(c)
	if masterForm != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": masterForm.Error()})
	} else {
		c.JSON(http.StatusOK, masterForm)
	}
}
