package hyperflake

import (
	"testing"
)

func TestDecodeID(t *testing.T) {
	hf := NewHyperflakeConfig(0, 0)

	testCases := []struct {
		id                     int64
		expectedSignbit        int
		expectedTimestamp      int64
		expectedDatacenterID   int
		expectedMachineID      int
		expectedSequenceNumber int
	}{
		{
			id:                     3282575599297626112,
			expectedSignbit:        0,
			expectedTimestamp:      1729311810178,
			expectedDatacenterID:   0,
			expectedMachineID:      0,
			expectedSequenceNumber: 0,
		},
		{
			id:                     3273974649684016911,
			expectedSignbit:        0,
			expectedTimestamp:      1727261183992,
			expectedDatacenterID:   6,
			expectedMachineID:      11,
			expectedSequenceNumber: 3855,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			decodedID, err := hf.DecodeID(tc.id)
			if err != nil {
				t.Fatalf("DecodeID(%d) returned error: %v", tc.id, err)
			}
			if decodedID.Signbit != tc.expectedSignbit {
				t.Errorf("DecodeID(%d) Signbit = %d; want %d", tc.id, decodedID.Signbit, tc.expectedSignbit)
			}
			if decodedID.Timestamp != tc.expectedTimestamp {
				t.Errorf("DecodeID(%d) Timestamp = %d; want %d", tc.id, decodedID.Timestamp, tc.expectedTimestamp)
			}
			if decodedID.DatacenterID != tc.expectedDatacenterID {
				t.Errorf("DecodeID(%d) DatacenterID = %d; want %d", tc.id, decodedID.DatacenterID, tc.expectedDatacenterID)
			}
			if decodedID.MachineID != tc.expectedMachineID {
				t.Errorf("DecodeID(%d) MachineID = %d; want %d", tc.id, decodedID.MachineID, tc.expectedMachineID)
			}
			if decodedID.SequenceNumber != tc.expectedSequenceNumber {
				t.Errorf("DecodeID(%d) SequenceNumber = %d; want %d", tc.id, decodedID.SequenceNumber, tc.expectedSequenceNumber)
			}
		})
	}

	idd, _ := hf.GenerateHyperflakeID()

	println(idd)
}
