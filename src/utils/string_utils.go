package utils

import "strconv"

func StringOrNil(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func StringToInt64(s string) int64 {
	return int64(StringToInt(s))
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
