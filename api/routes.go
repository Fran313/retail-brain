package api

import (
	"github.com/Fran313/retailBrain/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/api/upload-excel", controller.UploadExcelHandler)
}
