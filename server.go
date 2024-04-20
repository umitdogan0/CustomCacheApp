package main

import (
	"custom_cache_app/configuration"
	"custom_cache_app/controllers"
	"custom_cache_app/cron_jobs"
	"github.com/gin-gonic/gin"
)

func main() {
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
