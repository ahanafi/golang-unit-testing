package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Hanafi")
	if result != "Hello Hanafi" {
		// t.Fail() // <-- Fail method will continue the code execution
		t.Error("Result must be 'Hello Hanafi'") // <-- Will call Fail but with additional message/information
	}

	fmt.Println("TestHelloWorld has done executed.")
}

func TestHelloAhmad(t *testing.T) {
	result := HelloWorld("Ahmad")
	if result != "Hello Ahmad" {
		// t.FailNow() // <-- FailNow will stop to next code execution
		t.Fatal("Result must be 'Hello Ahmad'")
	}

	fmt.Println("TestHelloAhmad has done executed.")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Ahmad")
	expectedResult := "Hello Ahmad"
	assert.Equal(t, expectedResult, result, "Result mus be 'Hello Ahmad'")

	fmt.Println("TestHelloWorldAssertion has been executed")
}