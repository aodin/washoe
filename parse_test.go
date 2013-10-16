package washoe

import (
	"encoding/csv"
	"os"
	"testing"
)

var testFile = "./testdata/addresses.csv"

func TestParseCSV(t *testing.T) {
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatal(err)
	}
	csvReader := csv.NewReader(file)
	addresses, parseErr := ParseAddressesCSV(csvReader)
	if parseErr != nil {
		t.Fatal(parseErr)
	}
	if len(addresses) != 3 {
		t.Fatalf("Unexpected length of addresses: %d", len(addresses))
	}

	patterson := addresses[0]
	if patterson.Full != `12 S PATTERSON PL` {
		t.Errorf("Unexpected address: %s", patterson.Full)
	}
	if patterson.Number != 12 {
		t.Errorf("Unexpected number: %d", patterson.Number)
	}
	if patterson.Zip != 89436 {
		t.Errorf("Unexpected ZIP: %d", patterson.Zip)
	}
	if patterson.Latitude != 39.632746 {
		t.Errorf("Unexpected latitude: %s", patterson.Latitude)
	}
	if patterson.Longitude != -119.717986 {
		t.Errorf("Unexpected longitude: %s", patterson.Longitude)
	}
}
