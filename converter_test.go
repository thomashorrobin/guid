package guid

import (
	"testing"

	"github.com/google/uuid"
)

func TestFromGUID(t *testing.T) {
	// set up
	guidString := "0671ba09-fd1c-4590-ab73-7f8d8a83d7d2"
	expectedUUID, _ := uuid.Parse("09ba7106-1cfd-9045-ab73-7f8d8a83d7d2")

	// act
	uuidByteArray := FromGUID(guidString)

	// assert
	calculatedUUID := uuid.UUID(uuidByteArray)
	if calculatedUUID != expectedUUID {
		t.Errorf("calculated uuid (%s) does not equeal expected (%s)", calculatedUUID.String(), expectedUUID.String())
	}
}

func TestToGUID(t *testing.T) {
	// set up
	startingUUID, _ := uuid.Parse("09ba7106-1cfd-9045-ab73-7f8d8a83d7d2")
	expectedGUID := "0671ba09-fd1c-4590-ab73-7f8d8a83d7d2"

	// act
	calculatedGUID := ToGUID(startingUUID)

	// assert
	if calculatedGUID != expectedGUID {
		t.Errorf("calculated uuid (%s) does not equeal expected (%s)", calculatedGUID, expectedGUID)
	}
}
