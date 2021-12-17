package router

import (
	"path/handler"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", handler.GetAllStudent)
	router.POST("/", handler.CretateStudent)
	router.PATCH("/:id", handler.UpdateStudent)
	router.DELETE("/:id", handler.DeleteStudent)
}
