package Routers

import (
	"Demo/Controller"
	"github.com/gin-gonic/gin"
)

//注册路由
func RegisterRouter() *gin.Engine {
	r := gin.Default() //gin.New() //

	//路由
	v1:=r.Group("/orm/")
	{
		v1.GET("/Select/:id", Controller.Select)
		v1.POST("/Insert",Controller.Insert)
		v1.POST("/Update",Controller.Update)
		v1.DELETE(
			"/Delete",
			Controller.Delete,
		)
	}
	return r
}
