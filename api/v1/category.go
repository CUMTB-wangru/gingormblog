package v1

import (
	"ginblog-master/model"
	"ginblog-master/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// gin框架接口 上下文：c *gin.Context  表明该函数是api接口的处理函数
// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCSE {
		model.CreateCate(&data)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// GetCateInfo 查询分类信息
func GetCateInfo(c *gin.Context) {
	// c.Param("id") 返回GET请求方法中URL的param 参数
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetCateInfo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)

}

// GetCate 查询分类列表
func GetCate(c *gin.Context) {
	// c.Query 拿到请求携带的参数(string)，没有则返回空  第二个参数可以设置返回默认值
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := model.GetCate(pageSize, pageNum)
	code := errmsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 查询单个分类
//func GetCateInfo(c *gin.Context)  {
//	id, _ := strconv.Atoi(c.Param("id"))
//
//	data,code := model.GetCateInfo(id)
//
//	c.JSON(http.StatusOK, gin.H{
//		"status":  code,
//		"data":    data,
//		"message": errmsg.GetErrMsg(code),
//	})
//}

// EditCate 编辑分类名
func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCSE {
		model.EditCate(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DeleteCate 删除用户
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteCate(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
