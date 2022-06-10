package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/navisot/go-exchanger/config"
	driver "github.com/navisot/go-exchanger/database"
	"github.com/navisot/go-exchanger/routes"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, _ := SetupDatabase()

	router := SetupRouter(db)

	log.Fatalln(router.Run(":9090"))
}

func SetupDatabase() (*driver.DBDriver, error) {

	db, err := config.ConnectToDatabase(os.Getenv("DSN"))

	if err != nil {
		log.Fatal("no database connection")
	}

	return db, nil

}

func SetupRouter(dbDriver *driver.DBDriver) *gin.Engine {

	router := gin.Default()

	/**
	@description Setup Middleware
	*/
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	routes.InitConvertRoutes(dbDriver, router)

	return router

}
