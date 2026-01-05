package controller

import (
	"backend/internal/service"
	"backend/pkg/utils"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.Error(c, "文件解析失败")
		return
	}

	const maxFileSize = 100 << 20
	if file.Size > maxFileSize {
		utils.Error(c, "文件太大了，请不要超过 100MB")
		return
	}

	// 1. 调用 Service 处理逻辑并入库
	record, err := service.SaveFileRecord(file)
	if err != nil {
		utils.Error(c, "数据库写入失败")
		return
	}

	// 2. 保存物理文件
	dst := filepath.Join("./uploads", record.FileUUID)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		utils.Error(c, "物理文件保存失败")
		return
	}

	utils.Success(c, record)
}
