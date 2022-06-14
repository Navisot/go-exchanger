package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/navisot/go-exchanger/database"
	"github.com/navisot/go-exchanger/http/controllers"
	"github.com/navisot/go-exchanger/services"
)

// InitConvertRoutes holds the routes for the convert controller
func InitConvertRoutes(db *database.DBDriver, route *gin.Engine) {

	convertService := services.NewECBConvertService()
	convertController := controllers.ConvertController{I: convertService}

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/convert", convertController.Convert)
}
