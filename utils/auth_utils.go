package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

type User struct {
	UserName string `json:"username,omitempty"`
	PassWord string `json:"password,omitempty"`
}

var users []User

func InitUser() {
	user := User{
		UserName: "xiamo",
		PassWord: "17906193.",
	}
	users = append(users, user)
}

var hashStr string = "19991107"

func Auth(code string) {
	strArr := strings.Split(code, ".")
	user := FindUser(strArr[0])
	if (user == User{}) {
		panic("请先登陆！")
	}
	md5Str := GenMd5(user.UserName, user.PassWord)
	if md5Str != strArr[1] {
		panic("请先登陆！")
	}
}

func FindUser(username string) User {
	var user User
	for _, v := range users {
		if v.UserName == username {
			user = v
		}
	}
	return user
}

func GenMd5(str1, str2 string) string {
	str := str1 + str2 + hashStr
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
