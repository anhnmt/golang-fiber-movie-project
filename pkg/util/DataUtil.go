package util

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

//func parseStringToInt(str string) int64 {
//	n, err := strconv.ParseUint(str, 10, 64)
//	if err != nil {
//		return 0
//	}
//
//	return n
//}

func ParseStringToUInt64(str string) uint64 {
	n, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}

	return n
}

func ParseStringToUInt(str string) uint {
	return uint(ParseStringToUInt64(str))
}

func MD5(text string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
