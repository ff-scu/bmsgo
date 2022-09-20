package controller

import (
	"bmsgo/common"
	"bmsgo/model"
	"bmsgo/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

// Paging 分页接口
func Paging(ctx *gin.Context) {
	db := common.GetDB()

	//使用map获取请求的参数，json进行解析
	//var requestMap=make(map[string]string)
	//_ = json.NewDecoder(ctx.Request.Body).Decode(&requestMap)

	//pageSize,_:=strconv.Atoi(ctx.Query("page_size"))
	//pageNum,_:=strconv.Atoi(ctx.Query("page_num"))
	//绑定body中的参数
	var requestPaging model.Paging
	_ = ctx.Bind(&requestPaging)

	//获取path中的参数
	pageSize, _ := strconv.Atoi(ctx.Param("page_size"))
	pageNum, _ := strconv.Atoi(ctx.Params.ByName("page_num"))
	var bookName string = ctx.Param("book_name")

	var bookArr []model.Book
	var book model.Book

	//求出总数
	result := db.Find(&book)
	total := result.RowsAffected

	//进行分页查询
	db.Offset((pageNum-1)*pageSize).Limit(pageSize).Where("book_name like?", "%"+bookName+"%").Find(&bookArr)

	//不可以直接返回，放到切片中处理
	sliceBook := make([]model.Book, 0)
	for i := range bookArr {
		sliceBook = append(sliceBook, bookArr[i])
	}
	if sliceBook == nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "msg": "no bookArr found"})
		return
	}
	response.Success(ctx, gin.H{"total": total, "book": sliceBook}, "分页数据返回成功")
}

// AddBook 新增图书接口
func AddBook(ctx *gin.Context) {
	db := common.GetDB()
	var requestAddBook model.Book
	_ = ctx.Bind(&requestAddBook)

	bookName := requestAddBook.BookName
	bookAuth := requestAddBook.BookAuthor
	bookPublish := requestAddBook.BookPublish

	//判断是否存在该书名
	if isBooKNameExist(db, bookName) {
		//该书名存在
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该图书已经存在"})
		return
	}

	newBook := model.Book{
		BookName:    bookName,
		BookAuthor:  bookAuth,
		BookPublish: bookPublish,
	}
	db.Create(&newBook)
	//返回结果
	response.Success(ctx, nil, "图书添加成功！")
}

//根据密码来查询用户是否存在
func isBooKNameExist(db *gorm.DB, bookName string) bool {
	var book model.Book
	db.Where("book_name=?", bookName).First(&book)
	if book.ID != 0 {
		//用户存在
		return true
	}
	//用户不存在
	return false
}

// Paginate 分页封装
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

//编辑书名
func EditBook(ctx *gin.Context) {
	db:=common.GetDB()
	var requestEdit model.Book
	_=ctx.ShouldBind(&requestEdit)
	newBookName:=requestEdit.BookName
	db.Model(&model.Book{}).Where("id = ?", requestEdit.ID).Update("book_name", newBookName)
	response.Success(ctx,nil,"修改成功")
}

//删除图书
func DeleteBook(ctx *gin.Context) {

	db := common.GetDB()

	var requestDelete model.Book

	_ = ctx.Bind(&requestDelete)

	bookName:=requestDelete.BookName

	db.Where("book_name = ?", bookName).Delete(&requestDelete)

	response.Success(ctx,gin.H{},"删除成功")
}
