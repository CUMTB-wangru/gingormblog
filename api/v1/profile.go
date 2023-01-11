package v1

import (
	"ginblog-master/model"
	"ginblog-master/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// gin框架接口 上下文：c *gin.Context  表明该函数是api接口的处理函数
func GetProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetProfile(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func UpdateProfile(c *gin.Context) {
	var data model.Profile
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.UpdateProfile(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
