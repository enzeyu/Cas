package controller

import (
	"github.com/gin-gonic/gin"
	"Cas/agent"
)

func DefaultRoute(router *gin.Engine) {
	router.GET("/ceshi", func(context *gin.Context) {
		context.JSON(200, gin.H{"status": "OK"})
	})
	router.POST("/geturls", agent.Get_Result)
	router.POST("/darkchecks",agent.Get_Check)

}

