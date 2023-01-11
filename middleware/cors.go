package middleware

// 配置跨域--cors
import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			//AllowAllOrigins:  true, * 表示允许所有
			AllowOrigins:     []string{"*"}, // 等同于允许所有域名 #AllowAllOrigins:  true
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"*", "Authorization"},
			ExposeHeaders:    []string{"Content-Length", "text/plain", "Authorization", "Content-Type"},
			AllowCredentials: true,
			// 跨域保持12小时
			MaxAge: 12 * time.Hour,
		},
	)
}
