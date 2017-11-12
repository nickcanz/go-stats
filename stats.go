package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/montanaflynn/stats"
	"io"
	"os"
	"strconv"
)

func main() {
	var field int
	flag.IntVar(&field, "f", 1, "Field to calculate on (shorthand)")

	var delimiter string
	flag.StringVar(&delimiter, "d", ",", "Field delimiter (shorthand)")

	flag.Parse()

	var arg = flag.Arg(0)

	var input io.Reader
	if arg == "" {
		fmt.Println("Reading from stdin...")
		input = os.Stdin
	} else {
		fmt.Printf("First arg: %s\n", arg)
		f, err := os.Open(arg)
		if err == nil {
			input = f
		}
	}

	var numbers []float64

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if s, err := strconv.ParseFloat(scanner.Text(), 64); err == nil {
			numbers = append(numbers, s)
		}
	}

	fmt.Println(numbers)
	mean, err := stats.Mean(numbers)
	if err == nil {
		fmt.Printf("Average: %v\n", mean)
	} else {
		fmt.Println(err)
	}
}
