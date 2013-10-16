package washoe

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"testing"
)

var exampleAddress = &Address{
	Full:      "12 S PATTERSON PL",
	Number:    12,
	Street:    "PATTERSON",
	Zip:       89436,
	Latitude:  39.632746,
	Longitude: -119.717986,
}

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

func TestAddressJSON(t *testing.T) {
	// Test the output of an address to JSON
	jsonAddress, jsonErr := json.Marshal(exampleAddress)
	if jsonErr != nil {
		t.Fatal(jsonErr)
	}
	// Missing field will export as the initialization value
	expected := `{"id":0,"address":"12 S PATTERSON PL","number":12,"fraction":"","prefix":"","street":"PATTERSON","type":"","suffix":"","unit":"","city":"","county":"","state":"","zip":89436,"zip4":0,"latitude":39.632746,"longitude":-119.717986}`
	if string(jsonAddress) != expected {
		t.Errorf("Unexpected JSON output: %s", string(jsonAddress))
	}
}
