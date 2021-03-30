package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
	"github/fengjunhua/k8smanager/routers"
	"net/http"
)

//@title k8s管理项目
//@version 1.0
//@description k8s管理平台
func main() {
	router := gin.Default()
	router.Delims("[[", "]]")
	router.StaticFS("/static", http.Dir("./static"))
	router.LoadHTMLGlob("views/**/*")
	routers.LoadIndexRouters(router)
	routers.LoadLoginRouters(router)
	routers.LoadTagRouters(router)
	routers.LoadK8sRouters(router)
	routers.LoadMysqlRouters(router)
	routers.LoadLinuxCmdRouters(router)
	routers.LoadUsersRouters(router)
	routers.LoadViewsRouters(router)
	//router.LoadHTMLFiles("views/user/userList.tmpl")
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Run()
}