package lib

import (
	"strconv"
	"strings"
)

func IntToBinaryString(n int, length int) string {
	binaryStr := strconv.FormatInt(int64(n), 2)
	return padString(binaryStr, length)
}

func BinaryStringToInt(binaryStr string) (int64, error) {
	n, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func padString(s string, length int) string {
	if len(s) >= length {
		return s
	}
	return strings.Repeat("0", length-len(s)) + s
}
