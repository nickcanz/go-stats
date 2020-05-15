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

func ExamplePrintStats_10k() {
	exampleData := []float64{}
	for i := 0; i < 10000; i++ {
		exampleData = append(exampleData, float64(i))
	}

	printStats(exampleData)

	// Output:
	// count: 10000
	// min: 0.00
	// avg: 4999.50
	// max: 9999.00
	// p50: 4999.00
	// p95: 9499.00
	// p99: 9899.00
}

func ExamplePrintStats1_10() {
	exampleData := []float64{}
	for i := 1; i <= 10; i++ {
		exampleData = append(exampleData, float64(i))
	}

	printStats(exampleData)

	// Output:
	// count: 10
	// min: 1.00
	// avg: 5.50
	// max: 10.00
	// p50: 5.00
	// p95: 9.50
	// p99: 9.50
}
