package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main() {
	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
	fmt.Println("Finish")

	var ice string

	ice = GetMD5Hash("HEKKKI")

	fmt.Println(ice)
}
