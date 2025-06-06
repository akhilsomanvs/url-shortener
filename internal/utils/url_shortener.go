package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const keyLength = 7

func GenerateUniqueKey() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}

func GenerateHashKey(longUrl string) string {
	hash := md5.Sum([]byte(longUrl))
	return hex.EncodeToString(hash[:])
}

func GetHashWithKeyLength(uniqueHash []string, startIndex int) string {
	endIndex := len(uniqueHash) - 1
	fmt.Println(endIndex, uniqueHash)
	if startIndex+keyLength-1 <= endIndex {
		fmt.Println(endIndex, strings.Join(uniqueHash[startIndex:startIndex+keyLength], ""))
		return strings.Join(uniqueHash[startIndex:startIndex+keyLength], "")
	}
	return ""
}
