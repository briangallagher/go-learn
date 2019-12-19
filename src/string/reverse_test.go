package string


import "testing"

// Tests for a package live in the same directory in files that end in _test
// When you run go build or go install test files are not included
// When you run go test all files are compiled (test and non test) and create a 
// test harness that runs all tests beginning with a capital T

func Test(t *testing.T) {

	// Test will pass by default
	// uncomment below to make fail

	// t.Errorf("Test failed")
}


