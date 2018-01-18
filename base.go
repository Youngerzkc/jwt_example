// 这是一个基本的JWT生成过程
//用到的知识点有map转json,base64编码，sha256加密
package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func main() {
	head := map[string]string{}
	head["alg"] = "HS256"
	head["typ"] = "JWT"

	head_json, _ := json.Marshal(head)
	base := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+-")
	res := base.EncodeToString(head_json)
	fmt.Println("res is ", res)
	fmt.Println("head_json is ", string(head_json))
	playload := map[string]string{}
	playload["iss"] = "John Wu JWT"
	playload["iat"] = "1441593502"
	playload["exp"] = "1441594722"
	playload["aud"] = "www.example.com"
	playload["sub"] = "jrocket@example.com"
	playload["from_user"] = "B"
	playload["target_user"] = "A"
	playload_json, err := json.Marshal(playload)
	fmt.Println("playload_json is", string(playload_json))
	if err != nil {
		panic(err)
	}
	playload64 := base.EncodeToString(playload_json)
	fmt.Println("playload64 is ", playload64)

	signstring := res + "." + playload64
	sha := sha256.New()
	sha.Write([]byte(signstring))              //Write写入需要哈希的内容
	signres := sha.Sum([]byte("secrethelloo")) //salt,增加安全性的
	fmt.Println("signres is ", string(signres))
	signbase := base.EncodeToString(signres)
	fmt.Println("signbase is ", signbase)

	jwtres := signstring + "." + signbase
	fmt.Println("jwt res is ", jwtres)

	return
}
