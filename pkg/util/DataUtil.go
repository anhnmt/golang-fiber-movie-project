package util

import "strconv"

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
