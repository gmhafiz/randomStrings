package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
)

var iteration int
var length int
var randType string

func init() {
	flag.IntVar(&iteration, "n", 1, "Number of random strings to generate")
	flag.IntVar(&length, "l", 31, "Length of the random string to generate")
	flag.StringVar(&randType, "t", "all",
		"all|alpha|number|alphanum Length of the random string to generate")
	flag.Parse()
}

// MODIFIED from https://socketloop.com/tutorials/golang-how-to-generate-random-string
func randStr(strSize int, randType string) []byte {

	var dictionary string
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	numbers := "0123456789"
	specialCharacters := "~@#$%^&*()_+|[]{}';\"?/.,><"

	if randType == "alpha" {
		dictionary = alphabet
	}

	if randType == "number" {
		dictionary = numbers
	}

	if randType == "alphanum" {
		dictionary = numbers + alphabet
	}

	if randType == "all" {
		dictionary = numbers + alphabet + specialCharacters
	}

	var bytes = make([]byte, strSize)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return bytes
}

func main() {
	n := 0
	for n < iteration {
		randomString := randStr(length, randType)
		fmt.Println(string(randomString))
		n += 1
	}
}
