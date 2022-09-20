package controller

import (
	"bmsgo/common"
	"bmsgo/dto"
	"bmsgo/model"
	"bmsgo/response"
	"bmsgo/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Register 注册函数控制器
func Register(ctx *gin.Context) {
	DB := common.GetDB()

	var requestRegister = model.User{}
	//第一种方案
	//json.NewDecoder(ctx.Request.Body).Decode(&requestRegister)
	//第二种方案
	_ = ctx.Bind(&requestRegister)

	//获取参数
	name := requestRegister.Name
	telephone := requestRegister.Telephone
	password := requestRegister.Password

	//手机号不可以小于11位
	if len(telephone) < 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号不可以少于11位")
		return
	}

	//验证数据
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不可以少于6位"})
		return
	}

	//判断是否存在该用户
	if isNameExist(DB, telephone) {
		//用户存在
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该用户已经存在"})
		return
	}

	//没有填入name随机生成
	if len(name) == 0 {
		name = utils.RandomString(10)
		log.Println("随机生成的name" + name)
	}

	//如果用户不存在，新建用户,首先对密码进行加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "密码加密失败"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	//新建用户
	DB.Create(&newUser)

	log.Println(name, password, telephone)

	//密码正确发放token给前端
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token}, "注册成功！")
}

// Login 登录的控制器
func Login(ctx *gin.Context) {

	db := common.GetDB()
	//使用map获取请求的参数
	//var requestMap=make(map[string]string)
	//_ = json.NewDecoder(ctx.Request.Body).Decode(&requestMap)
	//
	//telephone:=requestMap["name"]
	//password:=requestMap["password"]

	var requestUser = model.User{}
	_ = ctx.Bind(&requestUser)
	//获取参数
	telephone := requestUser.Telephone
	password := requestUser.Password

	//手机号不可以小于11位
	if len(telephone) < 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号不可以少于11位"})
		return
	}

	//验证数据
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不可以少于6位"})
		return
	}

	//判断手机号是否存在
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该用户不存在"})
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	//密码正确发放token给前端
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功！")

}

// Info 登录成功验证用户信息控制器
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

//根据密码来查询用户是否存在
func isNameExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		//用户存在
		return true
	}
	//用户不存在
	return false
}
