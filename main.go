package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.devgroup.tech/shkolkovo/romanych/api"
	"gitlab.devgroup.tech/shkolkovo/romanych/db"
	"gitlab.devgroup.tech/shkolkovo/romanych/helpers"
	"gitlab.devgroup.tech/shkolkovo/romanych/log"
	"os"
)


func main() {
	_ = godotenv.Load()
	var err error

	if err = db.ConnectDatabase(); err != nil {
		log.GetLogger().Errorf("Can't connect to database: %s", err.Error())
		os.Exit(1)
	}

	r := gin.Default()

	api.AllApi(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := helpers.GetEnvDefault("PORT", "3000")

	if err = r.Run(":" + port); err != nil {
		log.GetLogger().Errorf("Can't start server: %s", err.Error())
		os.Exit(1)
	}
}
