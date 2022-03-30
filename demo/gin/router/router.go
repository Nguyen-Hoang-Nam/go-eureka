package router

import (
	"github.com/gin-gonic/gin"
	"mlm.gin/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/", controllers.Hi)

	router.GET("/hello/:framework", controllers.Hello)

	router.GET("/test", controllers.Test)

	return router
}

func Init() {
	router := NewRouter()

	router.Run(":8002")
}
