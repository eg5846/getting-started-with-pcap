package utils

import (
	"crypto/md5"
	"fmt"
)

func CalculateMd5Sum(bytes []byte) string {
	return fmt.Sprintf("%x", md5.Sum(bytes))
}
