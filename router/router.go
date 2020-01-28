package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-crud/controller"
)

// SetupRouter - setando rotas da api
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("user", controller.FindAll)
		v1.POST("user", controller.Create)
		v1.PUT("user/:id", controller.Update)
		v1.GET("user/:id", controller.FindOne)
	}
	return r
}
