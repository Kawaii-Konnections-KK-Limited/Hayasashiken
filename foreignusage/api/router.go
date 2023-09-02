package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r := gin.New()

	r.Use(cors.New(config))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/test", testHandler)
	r.GET("/ping", pingSelf)

	return r
}
