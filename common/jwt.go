package common

import (
	"bmsgo/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//定义加密的秘钥
var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

//发放token
func ReleaseToken(user model.User) (string, error) {

	//token的有效期
	expirationTime := time.Now().Add(1 * 24 * time.Hour)

	//创建一个claims
	Claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			//过期时间
			ExpiresAt: expirationTime.Unix(),
			//签名时间
			IssuedAt: time.Now().Unix(),
			//签名颁发者
			Issuer: "hsm",
			//签名主题
			Subject: "user token",
		},
	}

	// 使用指定的签名加密方式创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)

	//使用 secretKey 密钥进行加密处理后拿到最终 token string
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
