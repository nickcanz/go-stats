package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"github.com/montanaflynn/stats"
)

var field int
var delimiter string

func init() {
	flag.IntVar(&field, "f", 1, "Field to calculate on (shorthand)")
	flag.StringVar(&delimiter, "d", ",", "Field delimiter (shorthand)")
}

func printStats(numbers []float64) {
  var count = len(numbers)
  fmt.Printf("Count: %v\n", count)

	min, err := stats.Min(numbers)
	if err == nil {
		fmt.Printf("Min: %v\n", min)
	} else {
		fmt.Println(err)
	}

	max, err := stats.Max(numbers)
	if err == nil {
		fmt.Printf("Max: %v\n", max)
	} else {
		fmt.Println(err)
	}

	mean, err := stats.Mean(numbers)
	if err == nil {
		fmt.Printf("Average: %v\n", mean)
	} else {
		fmt.Println(err)
	}
}

func parseLine(line string) (float64,error) {
  tokens := strings.Split(line, delimiter)
  return strconv.ParseFloat(tokens[field-1], 64)
}

func main() {

	flag.Parse()

	var arg = flag.Arg(0)

	var input io.Reader
	if arg == "" {
		input = os.Stdin
	} else {
		f, err := os.Open(arg)
		if err == nil {
			input = f
		}
	}

	var numbers []float64

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
    if number, err := parseLine(scanner.Text()); err == nil {
			numbers = append(numbers, number)
    }
	}

  printStats(numbers)

}
