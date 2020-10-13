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

type MasterFormRepository interface {
	SaveMasterForm(masterForm entity.MasterForm) error
	UpdateMasterForm(masterForm entity.MasterForm) error
	DeleteMasterForm(masterForm entity.MasterForm) error
	GetAllMasterForms() []entity.MasterForm
	GetMasterForm(ctx *gin.Context) []entity.MasterForm
}

type masterFormDatabase struct {
	connection *gorm.DB
}

func NewMasterFormRepository() MasterFormRepository {
	db, _ := config.ConnectDB()
	db.AutoMigrate(&entity.MasterForm{})
	return &masterFormDatabase{
		connection: db,
	}
}

func (db *masterFormDatabase) SaveMasterForm(masterForm entity.MasterForm) error {
	data := &masterForm
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	db.connection.Create(data)
	return nil
}

func (db *masterFormDatabase) UpdateMasterForm(masterForm entity.MasterForm) error {
	data := &masterForm
	data.UpdatedAt = time.Now()
	db.connection.Save(data)
	return nil
}

func (db *masterFormDatabase) DeleteMasterForm(masterForm entity.MasterForm) error {
	db.connection.Delete(&masterForm)
	return nil
}

func (db *masterFormDatabase) GetAllMasterForms() []entity.MasterForm {
	var masterForms []entity.MasterForm
	// db.connection.Set("gorm:auto_preload", true).Select("*").Joins("left join master_questions on master_questions.id = etl_sy_payments.id").Find(&masterForms)
	db.connection.Set("gorm:auto_preload", true).Find(&masterForms)
	return masterForms
}

func (db *masterFormDatabase) GetMasterForm(ctx *gin.Context) []entity.MasterForm {
	var masterForm []entity.MasterForm
	db.connection.Set("gorm:auto_preload", true).Where("id = ?", ctx.Param("id")).First(&masterForm)
	return masterForm
}
