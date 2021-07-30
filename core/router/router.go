package router

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/devops-backend-gin/routers"
)

func Routers() *gin.Engine{
	var Router = gin.Default()


	PublicGroup := Router.Group("")
	{
		routers.InitBaseRouter(PublicGroup)  // 注册基础功能路由 不做鉴权
		routers.InitInitRouter(PublicGroup)
	}

    PrivateGroup := Router.Group("")
    PrivateGroup.Use()
    {
    	routers.InitUserRouter(PrivateGroup)        //注册用户路由


	}


	return Router
}