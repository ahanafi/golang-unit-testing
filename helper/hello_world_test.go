package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Hanafi")
	if result != "Hello Hanafi" {
		t.Fail() // <-- Fail method will continue the code execution
	}

	fmt.Println("TestHelloWorld has done executed.")
}

func TestHelloAhmad(t *testing.T) {
	result := HelloWorld("Ahmad")
	if result != "Hello Ahmad" {
		t.FailNow() // <-- FailNow will stop to next code execution
	}

	fmt.Println("TestHelloAhmad has done executed.")
}