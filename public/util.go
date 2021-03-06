package public

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
)

func GenSaltPassword(salt, password string) string {
	//创建一个基于SHA256算法的hash.Hash接口的对象
	s1 := sha256.New()
	//输入数据
	s1.Write([]byte(password))
	//计算哈希值
	//str1 := fmt.Sprintf("%x",s1.Sum([]byte(password)))
	str1 := fmt.Sprintf("%x", s1.Sum(nil))
	s2 := sha256.New()
	s2.Write([]byte(str1 + salt))
	//return fmt.Sprintf("%x",s2.Sum([]byte(str1 + salt)))
	return fmt.Sprintf("%x", s2.Sum(nil))
}

//MD5 hash function
func MD5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Obj2Json(s interface{}) string {
	bts, _ := json.Marshal(s)
	return string(bts)
}
func InStringSlice(slice []string, str string) bool {
	for _, item := range slice {
		if str == item {
			return true
		}
	}
	return false
}
