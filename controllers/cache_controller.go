package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/umitdogan0/CustomCacheApp/entities"
	"github.com/umitdogan0/CustomCacheApp/services"
	"net/http"
	"time"
)

type DynamicData struct {
	Data       map[string]interface{} `json:"data"`
	Priority   int                    `json:"priority"`
	Expiration int                    `json:"expiration,omitempty"`
}

func RouteCacheController(server *gin.Engine) {
	server.GET("cache/", func(c *gin.Context) {
		data := services.GetCaches()
		c.JSON(http.StatusOK, data)
	})

	server.GET("cache/ram", func(c *gin.Context) {
		data := services.GetCollectionAlloc()
		c.JSON(http.StatusOK, data)
	})

	server.GET("cache/get", func(c *gin.Context) {
		key := c.Query("key")
		data, err := services.GetCache(key)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, data)
	})

	server.POST("cache/add", func(c *gin.Context) {
		fmt.Println(services.GetCollectionAlloc())
		var dynamicData DynamicData

		// JSON body'den verileri almak için BindJSON kullanılır
		if err := c.BindJSON(&dynamicData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(dynamicData)
		for key, value := range dynamicData.Data {
			err := services.MakeCache(&entities.CacheEntity{
				Data:       map[string]interface{}{key: value},
				Expiration: dynamicData.Expiration,
				Priority:   dynamicData.Priority,
				CreateDate: time.Now(),
			}, key)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		c.JSON(http.StatusOK, "Success")
	})
}
