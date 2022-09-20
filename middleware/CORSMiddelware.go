package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CORSMiddleware 跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//可以访问的域名
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//设置缓存时间
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//设置可以访问的方法
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
