package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Hanafi")
	if result != "Hello Hanafi" {
		panic("Results is not `Hello Hanafi`")
	}
}
