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
	masterAnswerRepository repositories.MasterAnswerRepository = repositories.NewMasterAnswerRepository()
	masterAnswerService    usecases.MasterAnswerService        = usecases.NewMasterAnswer(masterAnswerRepository)
	masterAnswerController deliveries.MasterAnswerController   = deliveries.NewMasterAnswer(masterAnswerService)
)

// GetMasterAnswersIndex godoc
// @Security basicAuth
// @Summary Show all existing Master_Answers
// @Description Get all existing Master_Answers
// @Tags Master_Answers
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.MasterAnswer
// @Failure 401 {object} dto.Response
// @Router /masterAnswer/index [get]
func MasterAnswersIndex(c *gin.Context) {
	masterAnswers := masterAnswerController.GetAllMasterAnswers()
	c.JSON(http.StatusOK, gin.H{"data": masterAnswers})
}

// GetMasterAnswersByQuestion godoc
// @Security basicAuth
// @Summary Show an existing Master_Answers by Question ID
// @Description Get detail the existing Master_Answers by Question ID
// @Tags Master_Answers
// @Accept  json
// @Produce  json
// @Param  id path int true "Master Question ID"
// @Success 200 {array} entity.MasterAnswer
// @Failure 401 {object} dto.Response
// @Router /masterAnswer/byQuestion/{id} [get]
func MasterAnswersByQuestion(c *gin.Context) {
	masterAnswer := masterAnswerController.GetMasterAnswerByQuestion(c)
	// buf := make([]byte, 1024)
	// num := c.Param("id")
	// reqBody := string(buf[0:num])
	// log.Print(num)
	c.JSON(http.StatusOK, gin.H{"data": masterAnswer})
}

// GetMasterAnswersDetail godoc
// @Security basicAuth
// @Summary Show an existing Master_Answers
// @Description Get detail the existing Master_Answers
// @Tags Master_Answers
// @Accept  json
// @Produce  json
// @Param  id path int true "Master Form ID"
// @Success 200 {array} entity.MasterAnswer
// @Failure 401 {object} dto.Response
// @Router /masterAnswer/detail/{id} [get]
func MasterAnswersDetail(c *gin.Context) {
	masterAnswer := masterAnswerController.GetMasterAnswer(c)
	// buf := make([]byte, 1024)
	// num := c.Param("id")
	// reqBody := string(buf[0:num])
	// log.Print(num)
	c.JSON(http.StatusOK, gin.H{"data": masterAnswer})
}

// CreateMasterAnswers godoc
// @Security basicAuth
// @Summary Create new Master_Answers
// @Description Create a new Master_Answers
// @Tags Master_Answers
// @Accept  json
// @Produce  json
// @Param masterAnswer body entity.MasterAnswerCreate true "Create masterAnswer"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /masterAnswer/create [post]
func MasterAnswerCreate(c *gin.Context) {
	masterAnswer := masterAnswerController.Save(c)
	if masterAnswer != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": masterAnswer.Error()})
	} else {
		c.JSON(http.StatusOK, masterAnswer)
	}
}

// UpdateMasterAnswers godoc
// @Security basicAuth
// @Summary Update Master_Answers
// @Description Update a Master_Answers
// @Tags Master_Answers
// @Accept  json
// @Produce  json
// @Param masterAnswer body entity.MasterAnswer true "Update masterAnswer"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /masterAnswer/update [put]
func MasterAnswerUpdate(c *gin.Context) {
	masterAnswer := masterAnswerController.Update(c)
	if masterAnswer != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": masterAnswer.Error()})
	} else {
		c.JSON(http.StatusOK, masterAnswer)
	}
}

// DeleteMasterAnswers godoc
// @Security basicAuth
// @Summary Delete Master_Answers
// @Description Delete a Master_Answers
// @Tags Master_Answers
// @Accept  json
// @Produce  json
// @Param masterAnswer body entity.MasterAnswerDelete true "Delete masterAnswer"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /masterAnswer/delete [delete]
func MasterAnswerDelete(c *gin.Context) {
	masterAnswer := masterAnswerController.Delete(c)
	if masterAnswer != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": masterAnswer.Error()})
	} else {
		c.JSON(http.StatusOK, masterAnswer)
	}
}
