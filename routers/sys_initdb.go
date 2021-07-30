package routers

import "github.com/gin-gonic/gin"

func InitInitRouter(Router *gin.RouterGroup){

	ApiRouter := Router.Group("init")
	{
		ApiRouter.POST("initdb")
	}

}