package zipcodes

//go:generate go run gen.go
type Zipcodes struct {}

func New() Zipcodes {
	return Zipcodes{}
}

func (z Zipcodes) Lookup(zip string) ZipcodeData {
	if data, ok := zipcodes[zip]; ok {
	    return data
	}
	return ZipcodeData{}
}
