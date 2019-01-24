package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

/*
sha练习
*/
func main() {
	//testSHA()
	testBase64()
}

func testSHA() {
	s := "hello"
	hash := sha256.New()
	hash.Write([]byte(s))
	bytes := hash.Sum(nil)
	toString := hex.EncodeToString(bytes)
	fmt.Println(toString)
}

func testBase64() {
	data := "abc123!?$*&()'-=@~"
	s := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(s)
	bytes, _ := base64.StdEncoding.DecodeString(s)
	fmt.Println(string(bytes))
	fmt.Println()

	// 如果要用在url中，需要使用URLEncoding  URL中的+号不识别吧
	toString := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(toString)
	decodeString, _ := base64.URLEncoding.DecodeString(toString)
	fmt.Println(string(decodeString))
}
