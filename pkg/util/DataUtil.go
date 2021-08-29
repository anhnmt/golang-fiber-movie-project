package util

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func ParseStringToInt64(str string) int64 {
	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}

	return n
}

func MD5(text string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
