package main

import (
	"github.com/gin-gonic/gin"
	"github.com/umitdogan0/CustomCacheApp/configuration"
	"github.com/umitdogan0/CustomCacheApp/controllers"
	"github.com/umitdogan0/CustomCacheApp/cron_jobs"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Recovery())
	controllers.RouteCacheController(r)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Custom Cache App",
		})
	})

	err := configuration.SetInitialConfiguration()

	if err != nil {
		panic(err)
	}
	go cron_jobs.CheckCacheExpr()
	r.Run(":" + configuration.GetConfiguration().Server.Port)
}
