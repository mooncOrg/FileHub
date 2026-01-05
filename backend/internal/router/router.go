package router

import (
	"backend/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.MaxMultipartMemory = 100 << 20

	r.POST("/upload", controller.UploadFile)

	return r
}
