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
	SaveFormGenerator(formGenerator entity.FormGenerator) error
	UpdateFormGenerator(formGenerator entity.FormGenerator) error
	DeleteFormGenerator(formGenerator entity.FormGenerator) error
	GetAllFormGenerators() []entity.FormGenerator
	GetFormGenerator(ctx *gin.Context) []entity.FormGenerator
}

type formGeneratorDatabase struct {
	connection *gorm.DB
}

func NewFormGeneratorRepository() FormGeneratorRepository {
	db, _ := config.ConnectDB()
	db.AutoMigrate(&entity.FormGenerator{})
	return &formGeneratorDatabase{
		connection: db,
	}
}

func (db *formGeneratorDatabase) SaveFormGenerator(formGenerator entity.FormGenerator) error {
	data := &formGenerator
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	db.connection.Create(data)
	return nil
}

func (db *formGeneratorDatabase) UpdateFormGenerator(formGenerator entity.FormGenerator) error {
	data := &formGenerator
	data.UpdatedAt = time.Now()
	db.connection.Save(data)
	return nil
}

func (db *formGeneratorDatabase) DeleteFormGenerator(formGenerator entity.FormGenerator) error {
	db.connection.Delete(&formGenerator)
	return nil
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
