package model

import (
	"log"

	"ginblog-master/utils/errmsg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	// 导入gorm的默认参数
	gorm.Model
	// 可以给字段指定特定规则  gorm：数据库数据规则  绑定json格式是为了前后端交互卓准备  validate: 数据验证规则 label: 进行映射
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// CheckUser 查询用户是否存在--避免添加重名的用户
func CheckUser(name string) (code int) {
	var user User
	// gorm操作数据库 按传入的名字查询user表中对应name的id 按主键排序查询第一个（类比sql语句）
	db.Select("id").Where("username = ?", name).First(&user)
	// user.ID > 0 说明数据库中已经存在这个用户 ID:gorm.Model的属性
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE
}

// CheckUpUser 更新查询
func CheckUpUser(id int, name string) (code int) {
	var user User
	db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCSE
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	// 加盐的哈希密码--写入数据库之前加密
	// data.Password = ScryptPw(data.Password)

	// gorm提供的方法,写入数据库中user表
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// GetUser 查询用户
func GetUser(id int) (User, int) {
	var user User
	// 每次查询一个用户--查处的数据放在实例对象user中
	err := db.Limit(1).Where("ID = ?", id).Find(&user).Error
	if err != nil {
		// 即使没有查询到也返回user数据  用于渲染
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCSE
}

// GetUsers 查询用户列表
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64

	// 接口没有传username过来
	if username != "" {
		// 使用LIKE 模糊查询 Find 查询所有符合记录的数据存在users中
		db.Select("id,username,role,created_at").Where(
			"username LIKE ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		// 使用LIKE 模糊查询
		db.Model(&users).Where(
			"username LIKE ?", username+"%",
		).Count(&total)
		return users, total
	}
	db.Select("id,username,role,created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	db.Model(&users).Count(&total)

	if err != nil {
		return users, 0
	}
	return users, total
}

// EditUser 编辑用户信息
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	// 通过id 编辑用户信息 将编辑后的信息插入数据库user表
	err = db.Model(&user).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// ChangePassword 修改密码
func ChangePassword(id int, data *User) int {
	//var user User
	//var maps = make(map[string]interface{})
	//maps["password"] = data.Password

	err = db.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	// gorm框架 在数据库user表中按id查询用户 并删除
	err = db.Where("id = ? ", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 钩子函数名字gorm框架固定的，不能改动  自动调用
// 钩子函数 BeforeCreate 密码加密&权限控制--每次数据库创建之前会调用这个函数
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	u.Role = 2
	return nil
}

// 钩子函数 BeforeUpdate 密码加密&权限控制--每次数据库更新之前会自动调用这个函数
func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return nil
}

// ScryptPw 生成密码---golang自带包
func ScryptPw(password string) string {
	const cost = 10
	// 生成加密的密码
	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}

	return string(HashPw)
}

// CheckLogin 后台登录验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.Where("username = ?", username).First(&user)

	// 验证密码是否正确
	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return user, errmsg.ERROR_USER_NO_RIGHT
	}
	return user, errmsg.SUCCSE
}

// CheckLoginFront 前台登录
func CheckLoginFront(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.Where("username = ?", username).First(&user)

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	return user, errmsg.SUCCSE
}
