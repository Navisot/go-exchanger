package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/navisot/go-exchanger/controllers"
	driver "github.com/navisot/go-exchanger/database"
	"github.com/navisot/go-exchanger/services"
)

func InitConvertRoutes(db *driver.DBDriver, route *gin.Engine) {

	convertService := services.NewECBConvertService()
	convertController := controllers.ConvertController{I: convertService}

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/convert", convertController.Convert)
}
