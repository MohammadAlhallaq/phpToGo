package structs

var Person struct {
	name     string
	location string
	age      int
	kids     []string
	books    map[string]string
	address  struct {
		line1 string
		line2 string
	}
}
