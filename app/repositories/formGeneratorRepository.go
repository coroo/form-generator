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

type FormGeneratorRepository interface {
	Save(formGenerator entity.FormGenerator)
	Update(formGenerator entity.FormGenerator)
	Delete(formGenerator entity.FormGenerator)
	GetAllFormGenerators() []entity.FormGenerator
	GetFormGenerator(ctx *gin.Context) []entity.FormGenerator
	CloseDB()
}

type formGeneratorDatabase struct {
	connection *gorm.DB
}

func NewFormGeneratorRepository() FormGeneratorRepository {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect formGeneratorDatabase")
	}
	db.AutoMigrate(&entity.FormGenerator{})
	return &formGeneratorDatabase{
		connection: db,
	}
}

func (db *formGeneratorDatabase) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close formGeneratorDatabase")
	}
}

func (db *formGeneratorDatabase) Save(formGenerator entity.FormGenerator) {
	data := &formGenerator
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	db.connection.Create(data)
}

func (db *formGeneratorDatabase) Update(formGenerator entity.FormGenerator) {
	data := &formGenerator
	data.UpdatedAt = time.Now()
	db.connection.Save(data)
}

func (db *formGeneratorDatabase) Delete(formGenerator entity.FormGenerator) {
	db.connection.Delete(&formGenerator)
}

func (db *formGeneratorDatabase) GetAllFormGenerators() []entity.FormGenerator {
	var formGenerators []entity.FormGenerator
	db.connection.Set("gorm:auto_preload", true).Find(&formGenerators)
	return formGenerators
}

func (db *formGeneratorDatabase) GetFormGenerator(ctx *gin.Context) []entity.FormGenerator {
	var formGenerator []entity.FormGenerator
	db.connection.Set("gorm:auto_preload", true).Where("id = ?", ctx.Param("id")).First(&formGenerator)
	return formGenerator
}