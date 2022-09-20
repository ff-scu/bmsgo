package main

import (
	"bmsgo/controller"
	"bmsgo/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CollectRoute 后端路由
func CollectRoute(r *gin.Engine) *gin.Engine {

	//跨域中间件
	r.Use(middleware.CORSMiddleware())

	//静态文件  第一个参数是api 第二个静态问价的文件夹相对目录
	r.StaticFS("/static", http.Dir("E:\\Code\\go\\bmsgo\\upload"))

	r.POST("/api/auth/register", controller.Register)

	r.POST("/api/auth/login", middleware.AuthMiddleware(), controller.Login)

	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	r.GET("/api/book/paging/:page_num/:page_size/:book_name", middleware.AuthMiddleware(), controller.Paging)

	r.GET("/api/book/paging", middleware.AuthMiddleware(), controller.Paging)

	r.POST("/api/book/addBook", middleware.AuthMiddleware(), controller.AddBook)

	r.POST("/api/book/edit",middleware.AuthMiddleware(),controller.EditBook)

	r.POST("/api/book/deleteBook",middleware.AuthMiddleware(),controller.DeleteBook)

	r.POST("/api/file/upload", middleware.AuthMiddleware(), controller.Upload)

	r.POST("/api/file/submitForm", middleware.AuthMiddleware(), controller.SubmitForm)

	r.POST("/api/image/uploadImage", middleware.AuthMiddleware(), controller.UploadImage)

	r.GET("/api/image/showImage", middleware.AuthMiddleware(), controller.ShowImage)

	r.POST("/api/image/deleteImage",middleware.AuthMiddleware(),controller.DeleteImage)

	r.POST("/api/qt/test",controller.TestQtPost)

	//请求该api发送模拟的数据
	r.POST("/api/qt/navi",controller.GetNaviData)

	return r
}
