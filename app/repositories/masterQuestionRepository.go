package repositories

import (
	"time"

	entity "github.com/coroo/form-generator/app/entity"
	"github.com/coroo/form-generator/config"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type MasterQuestionRepository interface {
	Save(masterQuestion entity.MasterQuestion)
	Update(masterQuestion entity.MasterQuestion)
	Delete(masterQuestion entity.MasterQuestion)
	GetAllMasterQuestions() []entity.MasterQuestion
	GetMasterQuestion(ctx *gin.Context) []entity.MasterQuestion
	CloseDB()
}

type masterQuestionDatabase struct {
	connection *gorm.DB
}

func NewMasterQuestionRepository() MasterQuestionRepository {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect masterQuestionDatabase")
	}
	db.AutoMigrate(&entity.MasterQuestion{})
	return &masterQuestionDatabase{
		connection: db,
	}
}

func (db *masterQuestionDatabase) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close masterQuestionDatabase")
	}
}

func (db *masterQuestionDatabase) Save(masterQuestion entity.MasterQuestion) {
	data := &masterQuestion
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	db.connection.Create(data)
}

func (db *masterQuestionDatabase) Update(masterQuestion entity.MasterQuestion) {
	data := &masterQuestion
	data.UpdatedAt = time.Now()
	db.connection.Save(data)
}

func (db *masterQuestionDatabase) Delete(masterQuestion entity.MasterQuestion) {
	db.connection.Delete(&masterQuestion)
}

func (db *masterQuestionDatabase) GetAllMasterQuestions() []entity.MasterQuestion {
	var masterQuestions []entity.MasterQuestion
	db.connection.Set("gorm:auto_preload", true).Find(&masterQuestions)
	return masterQuestions
}

func (db *masterQuestionDatabase) GetMasterQuestion(ctx *gin.Context) []entity.MasterQuestion {
	var masterQuestion []entity.MasterQuestion
	db.connection.Set("gorm:auto_preload", true).Where("id = ?", ctx.Param("id")).First(&masterQuestion)
	return masterQuestion
}
