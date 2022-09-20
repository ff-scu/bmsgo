package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//对返回的格式进行封装
// {
// 	code:201,
// 	data:"",
// 	msg:"hello hsm",
// }

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "records": data, "msg": msg})
}

// Success 成功和失败是Response的基础上继续进行的封装
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}

//请求直接返回字节流数据
func SendData(ctx *gin.Context,data gin.H,msg string)  {
	Response(ctx, http.StatusOK, 400, data, msg)
}
