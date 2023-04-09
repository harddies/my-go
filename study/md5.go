package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	s := fmt.Sprintf("%s%d", "abcd", 5)
	hash := md5.New()
	hash.Write([]byte(s))
	fmt.Println(hex.EncodeToString(hash.Sum(nil)))
}
