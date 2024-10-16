package main

import (
	"fmt"
	"log"
	"strconv"
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
		generateSnowflakeID(intToBinaryString(signBit),
			intToBinaryString(int(timestamp)),
			intToBinaryString(datacenterID),
			intToBinaryString(machineID),
			intToBinaryString(sequenceNumber)),
	)
}

func getCurrentTimestamp() int64 {
	return time.Now().UnixMilli() - epochMilli
}

func intToBinaryString(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func generateSnowflakeBinary(signBit string, timestamp string, datacenterID string, machineID string, sequenceNumber string) string {
	return signBit + padString(timestamp, 41) + padString(datacenterID, 5) + padString(machineID, 5) + padString(sequenceNumber, 12)
}

func generateSnowflakeID(signBit string, timestamp string, datacenterID string, machineID string, sequenceNumber string) int64 {
	snowflakeBinary := generateSnowflakeBinary(signBit, timestamp, datacenterID, machineID, sequenceNumber)
	snowflakeID, err := strconv.ParseInt(snowflakeBinary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return snowflakeID
}

func padString(s string, length int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(length)+"s", s)
}
