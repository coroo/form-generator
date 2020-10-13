package routes

import (
	"fmt"
	"net/http"

	deliveries "github.com/coroo/form-generator/app/deliveries"
	"github.com/coroo/form-generator/app/entity"
	repositories "github.com/coroo/form-generator/app/repositories"
	usecases "github.com/coroo/form-generator/app/usecases"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	masterQuestionRepository repositories.MasterQuestionRepository = repositories.NewMasterQuestionRepository()
	masterQuestionService    usecases.MasterQuestionService        = usecases.NewMasterQuestion(masterQuestionRepository)
	masterQuestionController deliveries.MasterQuestionController   = deliveries.NewMasterQuestion(masterQuestionService)
)

// GetMasterQuestionsIndex godoc
// @Security basicAuth
// @Summary Show all existing Master_Questions
// @Description Get all existing Master_Questions
// @Tags Master_Questions
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.MasterQuestion
// @Failure 401 {object} dto.Response
// @Router /masterQuestion/index [get]
func MasterQuestionsIndex(c *gin.Context) {
	masterQuestions := masterQuestionController.GetAllMasterQuestions()
	c.JSON(http.StatusOK, gin.H{"data": masterQuestions})
}

// GetMasterQuestionsDetail godoc
// @Security basicAuth
// @Summary Show an existing Master_Questions
// @Description Get detail the existing Master_Questions
// @Tags Master_Questions
// @Accept  json
// @Produce  json
// @Param  id path int true "Master Question ID"
// @Success 200 {array} entity.MasterQuestion
// @Failure 401 {object} dto.Response
// @Router /masterQuestion/detail/{id} [get]
func MasterQuestionsDetail(c *gin.Context) {
	masterQuestion := masterQuestionController.GetMasterQuestion(c)
	// buf := make([]byte, 1024)
	// num := c.Param("id")
	// reqBody := string(buf[0:num])
	// log.Print(num)
	c.JSON(http.StatusOK, gin.H{"data": masterQuestion})
}

// CreateMasterQuestions godoc
// @Security basicAuth
// @Summary Create new Master_Questions
// @Description Create a new Master_Questions
// @Tags Master_Questions
// @Accept  json
// @Produce  json
// @Param masterQuestion body entity.MasterQuestionCreate true "Create masterQuestion"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /masterQuestion/create [post]
func MasterQuestionCreate(c *gin.Context) {
	var masterQuestionEntity entity.MasterQuestion
	fmt.Println(c.Request.Body)
	c.ShouldBindJSON(&masterQuestionEntity)
	masterQuestion := masterQuestionController.Save(masterQuestionEntity)
	if masterQuestion != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": masterQuestion.Error()})
	} else {
		c.JSON(http.StatusOK, masterQuestion)
	}
}

// UpdateMasterQuestions godoc
// @Security basicAuth
// @Summary Update Master_Questions
// @Description Update a Master_Questions
// @Tags Master_Questions
// @Accept  json
// @Produce  json
// @Param masterQuestion body entity.MasterQuestion true "Update masterQuestion"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /masterQuestion/update [put]
func MasterQuestionUpdate(c *gin.Context) {
	var masterQuestionEntity entity.MasterQuestion
	fmt.Println(c.Request.Body)
	c.ShouldBindJSON(&masterQuestionEntity)
	masterQuestion := masterQuestionController.Update(masterQuestionEntity)
	if masterQuestion != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": masterQuestion.Error()})
	} else {
		c.JSON(http.StatusOK, masterQuestion)
	}
}

// DeleteMasterQuestions godoc
// @Security basicAuth
// @Summary Delete Master_Questions
// @Description Delete a Master_Questions
// @Tags Master_Questions
// @Accept  json
// @Produce  json
// @Param masterQuestion body entity.MasterQuestionDelete true "Delete masterQuestion"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /masterQuestion/delete [delete]
func MasterQuestionDelete(c *gin.Context) {
	var masterQuestionEntity entity.MasterQuestion
	fmt.Println(c.Request.Body)
	c.ShouldBindJSON(&masterQuestionEntity)
	masterQuestion := masterQuestionController.Delete(masterQuestionEntity)
	if masterQuestion != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": masterQuestion.Error()})
	} else {
		c.JSON(http.StatusOK, masterQuestion)
	}
}
