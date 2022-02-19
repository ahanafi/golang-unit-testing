package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	// Before running test
	fmt.Println("BEFORE: running unit test")

	m.Run()

	// After running test
	fmt.Println("AFTER: All unit testing has been executed")
}

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

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ahmad")
	expectedResult := "Hello Ahmad"
	require.Equal(t, expectedResult, result, "Result mus be 'Hello Ahmad'")
	fmt.Println("TestHelloWorldAssertion has been executed") // <-- This is will be not executed when use require.Equal()
}

// Skip Testing
func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Cannot run test in Linux OS")
	}

	result := HelloWorld("Ahmad")
	expectedResult := "Hello Ahmad"
	assert.Equal(t, expectedResult, result, "Result mus be 'Hello Ahmad'")
}

// Test SubTest
func TestSubTest(t *testing.T) {
	// Will be success
	t.Run("Ahmad", func(t *testing.T) {
		result := HelloWorld("Ahmad")
		expectedResult := "Hello Ahmad"
		assert.Equal(t, expectedResult, result, "Result mus be 'Hello Ahmad'")
	})

	// Will be success
	t.Run("Hanafi", func(t *testing.T) {
		result := HelloWorld("Hanafi")
		expectedResult := "Hello Hanafi"
		assert.Equal(t, expectedResult, result, "Result mus be 'Hello Hanafi'")
	})

	// Will be fail
	t.Run("Hanif", func(t *testing.T) {
		result := HelloWorld("Hanif")
		expectedResult := "Hi Hanif"
		assert.Equal(t, expectedResult, result, "Result mus be 'Hello Hanif'")
	})
}