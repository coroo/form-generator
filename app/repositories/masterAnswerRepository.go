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

type MasterAnswerRepository interface {
	SaveMasterAnswer(masterAnswer entity.MasterAnswer) error
	UpdateMasterAnswer(masterAnswer entity.MasterAnswer) error
	DeleteMasterAnswer(masterAnswer entity.MasterAnswer) error
	GetAllMasterAnswers() []entity.MasterAnswer
	GetMasterAnswer(ctx *gin.Context) []entity.MasterAnswer
	GetMasterAnswerByQuestion(ctx *gin.Context) []entity.MasterAnswer
	CloseDB()
}

type masterAnswerDatabase struct {
	connection *gorm.DB
}

func NewMasterAnswerRepository() MasterAnswerRepository {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect masterAnswerDatabase")
	}
	db.AutoMigrate(&entity.MasterAnswer{})
	return &masterAnswerDatabase{
		connection: db,
	}
}

func (db *masterAnswerDatabase) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close masterAnswerDatabase")
	}
}

func (db *masterAnswerDatabase) SaveMasterAnswer(masterAnswer entity.MasterAnswer) error {
	data := &masterAnswer
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	db.connection.Create(data)
	return nil
}

func (db *masterAnswerDatabase) UpdateMasterAnswer(masterAnswer entity.MasterAnswer) error {
	data := &masterAnswer
	data.UpdatedAt = time.Now()
	db.connection.Save(data)
	return nil
}

func (db *masterAnswerDatabase) DeleteMasterAnswer(masterAnswer entity.MasterAnswer) error {
	db.connection.Delete(&masterAnswer)
	return nil
}

func (db *masterAnswerDatabase) GetAllMasterAnswers() []entity.MasterAnswer {
	var masterAnswers []entity.MasterAnswer
	// db.connection.Set("gorm:auto_preload", true).Select("*").Joins("left join master_questions on master_questions.id = master_answers.id").Find(&masterAnswers)
	db.connection.Set("gorm:auto_preload", true).Find(&masterAnswers)
	return masterAnswers
}

func (db *masterAnswerDatabase) GetMasterAnswer(ctx *gin.Context) []entity.MasterAnswer {
	var masterAnswer []entity.MasterAnswer
	db.connection.Set("gorm:auto_preload", true).Where("id = ?", ctx.Param("id")).First(&masterAnswer)
	return masterAnswer
}

func (db *masterAnswerDatabase) GetMasterAnswerByQuestion(ctx *gin.Context) []entity.MasterAnswer {
	var masterAnswer []entity.MasterAnswer
	db.connection.Set("gorm:auto_preload", true).Where("master_question_id = ?", ctx.Param("id")).Find(&masterAnswer)
	return masterAnswer
}
