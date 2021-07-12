package main

import (
	"flag"
	"fmt"
	"randomStrings"
)
var (
	iteration int
	length    int
	randType  string
)

func main() {
	flag.IntVar(&iteration, "n", 1, "Number of random strings to generate")
	flag.IntVar(&length, "l", 32, "Length of the random string to generate")
	flag.StringVar(&randType, "t", "all",
		"all|alpha|number|alphanum Length of the random string to generate")
	flag.Parse()

	n := 0
	for n < iteration {
		randomString := randomStrings.RandStr(length, randType)
		fmt.Println(string(randomString))
		n += 1
	}
}
