package washoe

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestAutoAddressId(t *testing.T) {
	// The variable testFile is set in parse_test.go
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatal(err)
	}
	addresses, parseErr := ParseAddressesCSV(csv.NewReader(file))
	if parseErr != nil {
		t.Fatal(parseErr)
	}

	AutoAddressId(addresses)
	// Remember, id are int64!
	for index, address := range addresses {
		id := int64(index + 1)
		if id != address.Id {
			t.Errorf("Bad address id: %d != %d", address.Id, id)
		}
	}
}
