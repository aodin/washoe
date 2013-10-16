package washoe

import ()

// TODO What identifies a unique address?

type Address struct {
	Id        int64
	Full      string
	Number    int
	Fraction  string
	Prefix    string
	Street    string
	Type      string
	Suffix    string
	Unit      string
	City      string // aka Muni
	County    string
	State     string // Index 10
	Zip       int
	Zip4      int
	Latitude  float64
	Longitude float64
}

func (address *Address) String() string {
	return address.Full
}

// Give each address in the array an id equal to its index + 1
func AutoAddressId(addresses []*Address) {
	for index, address := range addresses {
		address.Id = int64(index + 1)
	}
}
