package v1

import (
	"net/http"
	"strconv"

	"ginblog-master/model"
	"ginblog-master/utils/errmsg"
	"ginblog-master/utils/validator"
	"github.com/gin-gonic/gin"
)

// gin框架接口 上下文：c *gin.Context  表明该函数是api接口的处理函数
// AddUser 添加用户
func AddUser(c *gin.Context) {
	// 创建user结构体实例
	var data model.User
	var msg string
	var validCode int
	// 要将请求体(c.ShouldBindJSON会从前端传回数据) 全部绑定到结构体的实例对象(data)
	// 什么时候使用ShouldBindJSON函数：需要将前端传回来的一些数据写入后端数据库，即需要改动数据库里的数据
	_ = c.ShouldBindJSON(&data)

	// 对传过来的数据进行验证
	msg, validCode = validator.Validate(&data)
	if validCode != errmsg.SUCCSE {
		// gin框架给前端返回JSON格式数据
		c.JSON(
			http.StatusOK, gin.H{
				"status":  validCode,
				"message": msg,
			},
		)
		//阻止调用后续的处理函数(中间件)
		c.Abort()
		return
	}

	// 接口api通过调用model中对应数据模型中的函数拿到后端的数据(数据库)  model.CheckUser(data.Username)
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		// 写入数据库表
		model.CreateUser(&data)
	}
	// gin框架给前端返回JSON格式数据
	c.JSON(
		// H is a shortcut for map[string]interface{}
		// type H map[string]interface{}  可以嵌套
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// GetUserInfo 查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data, code := model.GetUser(id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	// gin框架给前端返回JSON格式数据
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    maps,
			"total":   1,
			"message": errmsg.GetErrMsg(code),
		},
	)

}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	// c.Query 拿到请求携带的参数(string)，没有则返回空  第二个参数可以设置返回默认值
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	// 接口api通过调用model中对应数据模型中的函数拿到后端的数据(数据库)
	data, total := model.GetUsers(username, pageSize, pageNum)

	code := errmsg.SUCCSE
	// gin框架给前端返回JSON格式数据
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	// c.Param("id") 返回GET请求方法中URL的param 参数
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	// 接口api通过调用model中对应数据模型中的函数拿到后端的数据(数据库)
	code := model.CheckUpUser(id, data.Username)
	// 表明该用户名没有被占用，允许编辑
	if code == errmsg.SUCCSE {
		model.EditUser(id, &data)
	}
	// gin框架给前端返回JSON格式数据
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// ChangeUserPassword 修改密码
func ChangeUserPassword(c *gin.Context) {
	var data model.User
	// c.Param("id") 返回GET请求方法中URL的param 参数
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	// 接口api通过调用model中对应数据模型中的函数拿到后端的数据(数据库)
	code := model.ChangePassword(id, &data)
	// gin框架给前端返回JSON格式数据
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {

	// c.Param("id") 返回GET请求方法中URL的param 参数
	id, _ := strconv.Atoi(c.Param("id"))

	// 接口api通过调用model中对应数据模型中的函数拿到后端的数据(数据库)
	code := model.DeleteUser(id)

	// gin框架给前端返回JSON格式数据
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
