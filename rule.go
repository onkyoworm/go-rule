package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GetSHA256HashCode(message []byte)string{
	hash := sha256.New()
	hash.Write(message)
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}

func main() {
	message := []byte("test")
	hashCode := GetSHA256HashCode(message)
	fmt.Println(hashCode)

}
