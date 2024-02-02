package routers

import (
	"projects/api"
	"projects/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "x-api-key"}
	router.Use(cors.New(config))

	router.POST("/connect", api.RedirectGoogle)

	googleRouter := router.Group("/")
	googleRouter.Use(middleware.CheckToken())
	googleRouter.GET("/getList", api.GetList)

	return router
}
