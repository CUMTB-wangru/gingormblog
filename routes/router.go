package routes

// 这里的路由是后端路由：根据前端发起的请求路径url，找到对应请求路径下相应的api接口，
// 		通过调用api接口中的函数操作后端相应的数据模型(model)
import (
	v1 "ginblog-master/api/v1"
	"ginblog-master/middleware"
	"ginblog-master/utils"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func InitRouter() {
	// 读取setting 中的 AppMode 属性
	gin.SetMode(utils.AppMode)

	// 返回一个不带任何中间件的空白引擎
	r := gin.New()
	// 设置信任网络 []string
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)

	r.HTMLRender = createMyRender()

	// 使用自定义中间件--全局中间件：前后端交互必要经过中间件
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	// 将前端代码打包成静态资源，托管到后端静态资源管理
	// 这样的好处：部署到服务器，只需要部署后端
	r.LoadHTMLGlob("static/admin/index.html")
	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	// 托管之后的路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	/*
		后台管理路由接口
	*/
	// r.Group("api/v1") 路由组， 只有前缀是api/v1 的url 才能访问
	auth := r.Group("api/v1")
	// 该组路由全部使用jwt中间件做权限验证
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		// auth.XXX("url",callbacks)
		auth.GET("admin/users", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//修改密码
		// auth.XXX("url",callbacks)
		auth.PUT("admin/changepw/:id", v1.ChangeUserPassword)
		// 分类模块的路由接口
		// auth.XXX("url",callbacks)
		auth.GET("admin/category", v1.GetCate)
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		// 文章模块的路由接口
		// auth.XXX("url",callbacks)
		auth.GET("admin/article/info/:id", v1.GetArtInfo)
		auth.GET("admin/article", v1.GetArt)
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)
		// 上传文件
		// auth.XXX("url",callbacks)
		auth.POST("upload", v1.UpLoad)
		// 更新个人设置
		// auth.XXX("url",callbacks)
		auth.GET("admin/profile/:id", v1.GetProfile)
		auth.PUT("profile/:id", v1.UpdateProfile)
		// 评论模块
		// auth.XXX("url",callbacks)
		auth.GET("comment/list", v1.GetCommentList)
		auth.DELETE("delcomment/:id", v1.DeleteComment)
		auth.PUT("checkcomment/:id", v1.CheckComment)
		auth.PUT("uncheckcomment/:id", v1.UncheckComment)
	}

	/*
		前端展示页面接口
	*/
	router := r.Group("api/v1")
	{
		// 用户信息模块
		// router.XXX("url",callbacks)
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("users", v1.GetUsers)

		// 文章分类信息模块
		// router.XXX("url",callbacks)
		router.GET("category", v1.GetCate)
		router.GET("category/:id", v1.GetCateInfo)

		// 文章模块
		// router.XXX("url",callbacks)
		router.GET("article", v1.GetArt)
		router.GET("article/list/:id", v1.GetCateArt)
		router.GET("article/info/:id", v1.GetArtInfo)

		// 登录控制模块
		// router.XXX("url",callbacks)
		router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront)

		// 获取个人设置信息
		// router.XXX("url",callbacks)
		router.GET("profile/:id", v1.GetProfile)

		// 评论模块
		// router.XXX("url",callbacks)
		router.POST("addcomment", v1.AddComment)
		router.GET("comment/info/:id", v1.GetComment)
		router.GET("commentfront/:id", v1.GetCommentListFront)
		router.GET("commentcount/:id", v1.GetCommentCount)
	}

	_ = r.Run(utils.HttpPort)

}
