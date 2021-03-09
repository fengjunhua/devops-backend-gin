package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"tags":"sss"})
}
