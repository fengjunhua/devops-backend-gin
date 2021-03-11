package controllers

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/k8smanager/utils/cmd"
	_ "github/fengjunhua/k8smanager/utils/cmd"
)

func ExecLinuxCommand(c *gin.Context) {
	result := cmd.LinuxExec()
	//c.String(http.StatusOK,result)
	c.HTML(200, "table.html", gin.H{
		"title": result,
	})
}

func ExecLinuxWebSocket(c *gin.Context) {
	result := cmd.LinuxExec()
	//c.String(http.StatusOK,result)
	c.HTML(200, "websocket.html", gin.H{
		"title": result,
	})
}