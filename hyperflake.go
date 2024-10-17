package main

import (
	"fmt"
	"github.com/chirag3003/hyperflake-go/lib"
	"time"
)

type HyperflakeConfig struct {
	epoch              int64
	datacenterIDBits   int
	machineIDBits      int
	sequenceNumberBits int
	SignBit            int
	datacenterIDBinary string
	machineIDBinary    string
}

type HyperFlakeID struct {
	ID             int64
	Signbit        int
	DatacenterID   int
	MachineID      int
	SequenceNumber int
	Timestamp      int64
}

var DefaultEpoch = time.Date(2010, time.November, 4, 1, 42, 54, 0, time.UTC)
var defaultEpochMilli = DefaultEpoch.UnixMilli()

func NewHyperflakeConfig(
	datacenterIDBits int,
	machineIDBits int,
	signBit ...int) *HyperflakeConfig {
	sBit := 0
	if len(signBit) > 0 {
		sBit = signBit[0]
	}
	return &HyperflakeConfig{
		epoch:              defaultEpochMilli,
		datacenterIDBits:   datacenterIDBits,
		machineIDBits:      machineIDBits,
		sequenceNumberBits: 0,
		SignBit:            sBit,
		datacenterIDBinary: lib.IntToBinaryString(datacenterIDBits, 5),
		machineIDBinary:    lib.IntToBinaryString(machineIDBits, 5),
	}
}

func NewHyperflakeConfigWithEpoch(
	datacenterIDBits int,
	machineIDBits int,
	epoch int64,
	signBit ...int,
) *HyperflakeConfig {
	config := NewHyperflakeConfig(datacenterIDBits, machineIDBits, signBit...)
	config.epoch = epoch
	return config
}

func (config *HyperflakeConfig) GenerateSnowflakeID() (int64, error) {
	timestamp := lib.GetCurrentTimestampSinceEpoch(config.epoch)
	signBitBinary := lib.IntToBinaryString(config.SignBit, 1)
	timestampBinary := lib.IntToBinaryString(int(timestamp), 41)
	sequenceNumberBinary := lib.IntToBinaryString(config.sequenceNumberBits, 12)

	hyperflakeBinary := lib.BuildString(64,
		signBitBinary,
		timestampBinary,
		config.datacenterIDBinary,
		config.machineIDBinary,
		sequenceNumberBinary,
	)
	hyperflakeID, err := lib.BinaryStringToInt(hyperflakeBinary)
	return hyperflakeID, err

}

func DecodeID(id int64) (*HyperFlakeID, error) {
	// Convert the ID to a 64-bit binary string
	hyperflakeBinary := lib.IntToBinaryString(int(id), 64)

	// Extract and convert the sign bit
	signBitBinary := hyperflakeBinary[:1]
	signBit, err := lib.BinaryStringToInt(signBitBinary)
	if err != nil {
		return nil, err
	}

	// Extract and convert the timestamp
	timestampBinary := hyperflakeBinary[1:42]
	timestamp, err := lib.BinaryStringToInt(timestampBinary)
	if err != nil {
		return nil, err
	}

	// Extract and convert the datacenter ID
	datacenterIDBinary := hyperflakeBinary[42:47]
	datacenterID, err := lib.BinaryStringToInt(datacenterIDBinary)
	if err != nil {
		return nil, err
	}

	// Extract and convert the machine ID
	machineIDBinary := hyperflakeBinary[47:52]
	machineID, err := lib.BinaryStringToInt(machineIDBinary)
	if err != nil {
		return nil, err
	}

	// Extract and convert the sequence number
	sequenceNumberBinary := hyperflakeBinary[52:64]
	sequenceNumber, err := lib.BinaryStringToInt(sequenceNumberBinary)
	if err != nil {
		return nil, err
	}

	// Create and return the HyperFlakeID struct
	hyperflakeID := &HyperFlakeID{
		ID:             id,
		Signbit:        int(signBit),
		Timestamp:      timestamp,
		DatacenterID:   int(datacenterID),
		MachineID:      int(machineID),
		SequenceNumber: int(sequenceNumber),
	}
	return hyperflakeID, nil
}

func (config *HyperflakeConfig) SetMachineID(machineID int) {
	config.machineIDBits = machineID
	config.machineIDBinary = lib.IntToBinaryString(machineID, 5)
}

func (config *HyperflakeConfig) SetDatacenterID(datacenterID int) {
	config.datacenterIDBits = datacenterID
	config.datacenterIDBinary = lib.IntToBinaryString(datacenterID, 5)
}

func (config *HyperflakeConfig) GetMachineID() int {
	return config.machineIDBits
}

func (config *HyperflakeConfig) GetDatacenterID() int {
	return config.datacenterIDBits
}
func main() {
	config := NewHyperflakeConfig(1, 1)
	id, _ := config.GenerateSnowflakeID()
	println(id)
	decodedID, _ := DecodeID(id)
	fmt.Printf("%+v\n", decodedID)
}
