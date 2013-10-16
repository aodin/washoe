package washoe

import ()

type Address struct {
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
