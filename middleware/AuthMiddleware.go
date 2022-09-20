package middleware

import (
	"bmsgo/common"
	"bmsgo/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		//验证token的格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			//抛弃该次请求
			ctx.Abort()
			return
		}

		//提取token的有效部分
		tokenString = tokenString[7:]

		//解析token失败或解析到的token无效，则返回权限不足
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//否则证明token通过了验证，其中claims是解析出token的有效部分,获取claims中的userid
		userId := claims.UserID
		db := common.GetDB()
		var user model.User
		db.First(&user, userId)

		//用户不存在
		if userId == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//用户存在，将user信息写入上下文中
		ctx.Set("user", user)

		//执行路由 HandlerFunc
		ctx.Next()

	}
}
