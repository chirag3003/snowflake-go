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
	// Get the current timestamp
	fmt.Println(getPaddedString("h", 5))
	if i, err := strconv.ParseInt("1001", 2, 64); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
	timestamp := getTimestamp()
	println(timestamp)
	datacenter := 1
	machineID := 1
	sequenceNumber := 1
	signBit := 0
	log.Println("Timestamp: ", getBinaryFromInt(int(timestamp)))
	log.Println("Datacenter: ", getBinaryFromInt(datacenter))
	log.Println("Machine ID: ", getBinaryFromInt(machineID))
	log.Println("Sequence Number: ", getBinaryFromInt(sequenceNumber))
	log.Println("Sign Bit: ", getBinaryFromInt(signBit))
	log.Println("Snowflake ID: ",
		getSnowflakeID(getBinaryFromInt(signBit),
			getBinaryFromInt(int(timestamp)),
			getBinaryFromInt(datacenter),
			getBinaryFromInt(machineID),
			getBinaryFromInt(sequenceNumber)),
	)
}

func getTimestamp() int64 {
	return time.Now().UnixMilli() - epochMilli
}

func getBinaryFromInt(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func getSnowflakeBinary(signbit string, timestamp string, datacenter string, machineID string, sequenceNumber string) string {
	return signbit + getPaddedString(timestamp, 41) + getPaddedString(datacenter, 5) + getPaddedString(machineID, 5) + getPaddedString(sequenceNumber, 12)
}

func getSnowflakeID(signbit string, timestamp string, datacenter string, machineID string, sequenceNumber string) int64 {
	snowflakeBinary := getSnowflakeBinary(signbit, timestamp, datacenter, machineID, sequenceNumber)
	snowflakeID, err := strconv.ParseInt(snowflakeBinary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return snowflakeID
}

func getPaddedString(s string, len int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(len)+"s", s)
}
