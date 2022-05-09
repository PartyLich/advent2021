package runner

// Unimpl returns the string UNIMPLEMENTED
func Unimpl(_i interface{}) interface{} { return "UNIMPLEMENTED" }

// Solution defines methods required for each AoC puzzle solution
type Solution struct {
	Parse func(string) (interface{}, error)
	One   func(interface{}) interface{}
	Two   func(interface{}) interface{}
}
