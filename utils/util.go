package utils

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

// RandomString 返回随机字符串
func RandomString(n int) string {
	var letters = []byte("abcdhenHDINEGHKALENG")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

//10进制转化为16进制
func DecConvertToX(n, num int) (string,error) {
	if n < 0 {
		return strconv.Itoa(n),errors.New("只支持正整数")
	}
	if num != 2 && num != 8 && num != 16 {
		return strconv.Itoa(n),errors.New("只支持二、八、十六进制的转换")
	}
	result := ""
	h:=map[int]string{
		0:"0",
		1:"1",
		2:"2",
		3:"3",
		4:"4",
		5:"5",
		6:"6",
		7:"7",
		8:"8",
		9:"9",
		10:"A",
		11:"B",
		12:"C",
		13:"D",
		14:"E",
		15:"F",
	}
	for ; n > 0; n /= num {
		lsb := h[n % num]
		result = lsb + result
	}
	return result,nil
}
