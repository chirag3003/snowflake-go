package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

var epoch = time.Date(2010, time.November, 4, 1, 42, 54, 0, time.UTC)
var epochMilli = epoch.UnixMilli()

func main() {
	timestamp := getCurrentTimestamp()
	datacenterID := 1
	machineID := 1
	sequenceNumber := 1
	signBit := 0
	log.Println("Snowflake ID: ",
		generateSnowflakeID(signBit, timestamp, datacenterID, machineID, sequenceNumber),
	)
}

func getCurrentTimestamp() int64 {
	return time.Now().UnixMilli() - epochMilli
}

func intToBinaryString(n int, length int) string {
	binaryStr := strconv.FormatInt(int64(n), 2)
	return padString(binaryStr, length)
}

func generateSnowflakeBinary(signBit int, timestamp int64, datacenterID int, machineID int, sequenceNumber int) string {
	var builder strings.Builder
	builder.Grow(64) // Preallocate memory for the builder

	builder.WriteString(intToBinaryString(signBit, 1))
	builder.WriteString(intToBinaryString(int(timestamp), 41))
	builder.WriteString(intToBinaryString(datacenterID, 5))
	builder.WriteString(intToBinaryString(machineID, 5))
	builder.WriteString(intToBinaryString(sequenceNumber, 12))

	return builder.String()
}

func generateSnowflakeID(signBit int, timestamp int64, datacenterID int, machineID int, sequenceNumber int) int64 {
	snowflakeBinary := generateSnowflakeBinary(signBit, timestamp, datacenterID, machineID, sequenceNumber)
	snowflakeID, err := strconv.ParseInt(snowflakeBinary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return snowflakeID
}

func padString(s string, length int) string {
	if len(s) >= length {
		return s
	}
	return strings.Repeat("0", length-len(s)) + s
}
