package washoe

import ()

type Street struct {
	Full     string
	Number   int
	Fraction string
	Prefix   string
	Street   string
	Type     string
	Suffix   string
}

func (street *Street) String() string {
	return street.Full
}
