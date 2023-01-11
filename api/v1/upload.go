package v1

import (
	"ginblog-master/model"
	"ginblog-master/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin框架接口 上下文：c *gin.Context  表明该函数是api接口的处理函数
// UpLoad 上传图片接口
func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")

	fileSize := fileHeader.Size

	// 接口api通过调用model中对应数据模型中的函数拿到后端的数据
	url, code := model.UpLoadFile(file, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
