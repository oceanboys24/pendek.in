package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GeneratorURL(url string)  string {
	hash := md5.Sum([]byte(url))
	result := hex.EncodeToString(hash[:])[:6]
	
	return result
}
