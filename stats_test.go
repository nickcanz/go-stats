package main

import "testing"

func TestParseLine(t *testing.T) {
	delimiter = ","
	field = 2

	actual, _ := parseLine("foo,1,bar")

	if actual != 1 {
		t.Errorf("Expected %v got %v", 1, actual)
	}
}

func ExamplePrintStats() {
	exampleData := []float64{1, 99}

	printStats(exampleData)

	// Output:
	// Count: 2
	// Min: 1
	// Max: 99
	// Average: 50
}
