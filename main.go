package main

import (
	"log"
	"os"

	"github.com/coroo/form-generator/app/middlewares"
	"github.com/coroo/form-generator/app/routes"
	"github.com/coroo/form-generator/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = os.Getenv("MAIN_TITLE")
	docs.SwaggerInfo.Description = os.Getenv("MAIN_DESCRIPTION")
	docs.SwaggerInfo.Version = os.Getenv("MAIN_VERSION")
	docs.SwaggerInfo.Host = os.Getenv("MAIN_URL")
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	_, err := os.Stat("storage/logs")

	if os.IsNotExist(err) {
		err_0 := os.Mkdir("storage/logs", 0755)
		if err_0 != nil {
			log.Fatal(err_0)
		}
		err_1 := os.Mkdir("storage/logs/errors", 0755)
		if err_1 != nil {
			log.Fatal(err_1)
		}
		err_2 := os.Mkdir("storage/logs/information", 0755)
		if err_2 != nil {
			log.Fatal(err_2)
		}
	}

	router := gin.Default()
	router.Use(middlewares.BasicAuth(), middlewares.Logger())

	API_PREFIX := os.Getenv("API_PREFIX")

	router.GET("/", func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": os.Getenv("MAIN_DESCRIPTION"),
		})
	})

	masterQuestionsGroup := router.Group(API_PREFIX + "masterQuestion")
	{
		masterQuestionsGroup.GET("index", routes.MasterQuestionsIndex)
		masterQuestionsGroup.GET("detail/:id", routes.MasterQuestionsDetail)
		masterQuestionsGroup.POST("create", routes.MasterQuestionCreate)
		masterQuestionsGroup.PUT("update", routes.MasterQuestionUpdate)
		masterQuestionsGroup.DELETE("delete", routes.MasterQuestionDelete)
	}

	masterAnswersGroup := router.Group(API_PREFIX + "masterAnswer")
	{
		masterAnswersGroup.GET("index", routes.MasterAnswersIndex)
		masterAnswersGroup.GET("byQuestion/:id", routes.MasterAnswersByQuestion)
		masterAnswersGroup.GET("detail/:id", routes.MasterAnswersDetail)
		masterAnswersGroup.POST("create", routes.MasterAnswerCreate)
		masterAnswersGroup.PUT("update", routes.MasterAnswerUpdate)
		masterAnswersGroup.DELETE("delete", routes.MasterAnswerDelete)
	}

	masterFormsGroup := router.Group(API_PREFIX + "masterForm")
	{
		masterFormsGroup.GET("index", routes.MasterFormsIndex)
		masterFormsGroup.GET("detail/:id", routes.MasterFormsDetail)
		masterFormsGroup.POST("create", routes.MasterFormCreate)
		masterFormsGroup.PUT("update", routes.MasterFormUpdate)
		masterFormsGroup.DELETE("delete", routes.MasterFormDelete)
	}

	formGeneratorsGroup := router.Group(API_PREFIX + "formGenerator")
	{
		formGeneratorsGroup.GET("index", routes.FormGeneratorsIndex)
		formGeneratorsGroup.GET("detail/:id", routes.FormGeneratorsDetail)
		formGeneratorsGroup.POST("create", routes.FormGeneratorCreate)
		formGeneratorsGroup.PUT("update", routes.FormGeneratorUpdate)
		formGeneratorsGroup.DELETE("delete", routes.FormGeneratorDelete)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":3000")

}
