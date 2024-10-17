package lib

import "time"

func GetCurrentTimestampSinceEpoch(epoch int64) int64 {
	return time.Now().UnixMilli() - epoch
}

func GetCurrentTimestamp() int64 {
	return time.Now().UnixMilli()
}
